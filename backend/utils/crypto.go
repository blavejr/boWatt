package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func HashFileContent(reader io.Reader) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
