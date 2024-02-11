package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/Seiya-Tagami/pecopeco-cli/auth/util"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type authCode struct {
	code  string
	state string
	err   error
}

type codeReceiver struct {
	http.Server
	authCode chan authCode
}

func newServer() *codeReceiver {
	s := &codeReceiver{
		authCode: make(chan authCode),
	}
	router := chi.NewRouter()
	router.Get("/v1/auth/callback", s.authCallback(s.authCode))

	s.Server = http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	// TCPの接続をすぐに閉じるように
	s.Server.SetKeepAlivesEnabled(false)
	return s
}

func (s *codeReceiver) authCallback(ch chan<- authCode) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if errorMsg := r.FormValue("error"); errorMsg != "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to authorize: " + errorMsg))
			ch <- authCode{
				err: errors.New("Failed to authorize: " + errorMsg),
			}
			return
		}

		code := r.FormValue("code")
		if code == "" {
			errorMsg := "Failed to authorize. Code is empty"
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMsg))
			ch <- authCode{
				err: errors.New(errorMsg),
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized. Please back to your CLI."))

		ch <- authCode{
			code:  code,
			state: r.FormValue("state"),
		}
	}
}

func (s *codeReceiver) getAuthCode() (authCode, error) {
	ac, ok := <-s.authCode
	if !ok {
		return authCode{}, errors.New("authCode channel closed")
	}
	if ac.err != nil {
		return authCode{}, ac.err
	}
	return ac, nil
}

type OAuth struct {
	Config *oauth2.Config
	Token  *oauth2.Token
}

func NewOAuth(id, secret string, scopes []string, endpoint oauth2.Endpoint, redirectURL string) OAuth {
	return OAuth{
		Config: &oauth2.Config{
			ClientID:     id,
			ClientSecret: secret,
			Scopes:       scopes,
			Endpoint:     endpoint,
			RedirectURL:  redirectURL,
		},
	}
}

func (o *OAuth) Authorization(ctx context.Context) error {
	// 認可コードを取得のためのローカルサーバを一時的に立ち上げる
	s := newServer()
	defer s.Shutdown(context.Background())
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	const randomBytesLength = 128

	codeVerifier, err := util.GenerateRandomBytes(randomBytesLength)
	if err != nil {
		return err
	}
	b, err := util.GenerateRandomBytes(randomBytesLength)
	if err != nil {
		return err
	}
	// csrf対策に利用される文字列
	state := string(b)

	// oauthの開始
	url := o.Config.AuthCodeURL(
		string(state),
		oauth2.SetAuthURLParam("code_challenge", generateCodeChallenge(string(codeVerifier))),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)

	fmt.Printf("Logging into your Google account...\n\nClick following a URL: %v\n", url)
	if err = util.OpenBrowser(url); err != nil {
		fmt.Println(err)
	}

	authCode, err := s.getAuthCode()
	if err != nil {
		return err
	}

	if authCode.state != state {
		return errors.New("Failed to authorize. Invalid state")
	}

	token, err := o.Config.Exchange(
		ctx,
		authCode.code,
		oauth2.SetAuthURLParam("code_verifier", string(codeVerifier)),
		oauth2.SetAuthURLParam("grant_type", "authorization_code"),
	)

	if err != nil {
		return err
	}

	o.Token = token

	return nil
}

func generateCodeChallenge(codeVerifier string) string {
	hash := sha256.Sum256([]byte(codeVerifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}
