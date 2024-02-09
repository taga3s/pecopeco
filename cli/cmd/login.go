package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/Seiya-Tagami/pecopeco-cli/auth"
	"github.com/Seiya-Tagami/pecopeco-cli/auth/api/userinfo"
	"github.com/Seiya-Tagami/pecopeco-cli/auth/secret"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login pecopeco-cli",
	Long:  "You can login the pecopeco CLI with sth account.",
	Run: func(cmd *cobra.Command, args []string) {
		login()
	},
}

func login() {
	id, secret, scopes, redirectURL, err := secret.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	// OAuthによる処理
	oauth := auth.NewOAuth(
		id,
		secret,
		scopes,
		google.Endpoint,
		redirectURL,
	)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	if err := oauth.Authorization(ctx); err != nil {
		fmt.Println(err)
		return
	}

	userinfo, err := userinfo.Get(ctx, oauth)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[ID] %s\n[Name] %s\n[Email] %s\n", userinfo.ID, userinfo.Name, userinfo.Email)

	// ログイン処理
	// ...
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
