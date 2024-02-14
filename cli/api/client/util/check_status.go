package util

import "errors"

func CheckStatus(status int) error {
	switch status {
	case 200:
		return nil
	case 204:
		return nil
	case 400:
		return errors.New("Bad request")
	case 401:
		return errors.New("Invalid access token")
	case 500:
		return errors.New("Server-side error occurred")
	default:
		return errors.New("Unknown status code received")
	}
}
