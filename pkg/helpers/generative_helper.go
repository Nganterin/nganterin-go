package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"nganterin-go/pkg/exceptions"
	"time"
)

func GenerateToken(size int) (string, *exceptions.Exception) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateUniqueFileName() string {
	return time.Now().Format("20060102150405") + GenerateMilliseconds()
}

func GenerateMilliseconds() string {
	return time.Now().Format(".000")[1:]
}

func GenerateSecret(byteLength int) (string, *exceptions.Exception) {
	bytes := make([]byte, byteLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}
	secret := hex.EncodeToString(bytes)

	return secret, nil
}
