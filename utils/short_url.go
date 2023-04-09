package utils

import "fmt"

func GenerateShortLink(originalLink string) (string, error) {
	if EmptyString(originalLink) {
		return "", fmt.Errorf("empty String")
	}
	urlHash, err := Sha256Of(originalLink)
	if err != nil {
		return "", err
	}
	str := Base58Encode(urlHash)
	return str[:8], nil
}
