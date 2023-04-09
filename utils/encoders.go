package utils

import (
	"crypto/sha256"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

func EmptyString(str string) bool {
	str = strings.TrimSpace(str)
	return strings.EqualFold(str, "")
}

// Computes the hash value of a string
func Sha256Of(text string) ([]byte, error) {
	algorithm := sha256.New()
	text = strings.TrimSpace(text)
	if _, err := algorithm.Write([]byte(text)); err != nil {
		return nil, err
	}
	return algorithm.Sum(nil), nil
}

// Encode to base58
func Base58Encode(data []byte) string {
	return base58.Encode(data)
}
