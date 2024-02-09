package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsetLen := big.NewInt(int64(len(charset)))
	randomBytes := make([]byte, n)

	for i := range randomBytes {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return nil, err
		}
		randomBytes[i] = charset[num.Int64()]
	}

	return randomBytes, nil
}
