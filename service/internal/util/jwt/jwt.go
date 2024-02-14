package jwt

import (
	"context"
	"net/http"
	"os"

	"github.com/Seiya-Tagami/pecopeco-service/internal/config"
	"github.com/coreos/go-oidc/v3/oidc"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func SetHttpHeader(w http.ResponseWriter, accessToken string) {
	w.Header().Set("Authorization", accessToken)
}

func Verify(ctx context.Context, tokenString string) (*oidc.IDToken, error) {
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		return &oidc.IDToken{}, err
	}
	clientID := config.GetOIDC().ClientID
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	idToken, err := verifier.Verify(ctx, tokenString)
	if err != nil {
		return &oidc.IDToken{}, err
	}
	return idToken, nil
}

func GetUserIDFromToken(ctx context.Context, tokenString string) (string, error) {
	idToken, err := Verify(ctx, tokenString)
	if err != nil {
		return "", err
	}
	idTokenClaims := map[string]interface{}{}
	if err := idToken.Claims(&idTokenClaims); err != nil {
		return "", err
	}
	userID := idTokenClaims["sub"].(string)
	return userID, nil
}
