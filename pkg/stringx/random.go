package stringx

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of length n
func RandomString(n int) (string, error) {
	result := make([]byte, n)
	maxInt := big.NewInt(int64(len(charset)))

	for i := range result {
		randomInt, err := rand.Int(rand.Reader, maxInt)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomInt.Int64()]
	}

	return string(result), nil
}
