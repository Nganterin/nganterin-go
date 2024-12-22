package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateToken(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateUniqueFileName() string {
	return time.Now().Format("20060102150405") + GenerateMilliseconds()
}

func GenerateMilliseconds() string {
	return time.Now().Format(".000")[1:]
}

func GenerateSecret(byteLength int) (string, error) {
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate secret: %w", err)
	}
	secret := hex.EncodeToString(bytes)

	return secret, nil
}