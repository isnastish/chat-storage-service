package utils

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(data string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
