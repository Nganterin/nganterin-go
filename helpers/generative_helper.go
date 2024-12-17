package helpers

import (
	"crypto/rand"
	"encoding/base64"
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