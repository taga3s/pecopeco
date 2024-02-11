package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func Generate(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func SetHttpHeader(w http.ResponseWriter, accessToken string) {
	w.Header().Set("Authorization", accessToken)
}

func Verify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

func UserIDFromToken(tokenString string) (string, error) {
	token, err := Verify(tokenString)
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("user_id field not found or not a string")
		}
		return userID, nil
	}
	return "", errors.New("invalid token")
}
