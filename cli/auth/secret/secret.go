package secret

import (
	"errors"
	"os"
)

func Load() (string, string, []string, string, error) {
	id := os.Getenv("OAUTH_CLIENT_ID")
	secret := os.Getenv("OAUTH_CLIENT_SECRET")
	scopes := []string{os.Getenv("OAUTH_SCOPE_1"), os.Getenv("OAUTH_SCOPE_2"), os.Getenv("OAUTH_SCOPE_3")}
	redirectURL := os.Getenv("OAUTH_REDIRECT_URL")

	if id == "" || secret == "" || len(scopes) != 3 || redirectURL == "" {
		return "", "", []string{}, "", errors.New("OAUTH_CLIENT_ID, OAUTH_CLIENT_SECRET, OAUTH_SCOPE and OAUTH_REDIRECT_URL are required in env")
	}
	return id, secret, scopes, redirectURL, nil
}
