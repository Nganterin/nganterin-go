package middleware

import (
	"net/http"
	"os"
	"strings"
	"nganterin-go/api/users/dto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")

		var secretKey = []byte(secret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Status: http.StatusForbidden,
				Error:  "Forbidden",
			})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Authorization token",
			})
			return
		}

		tokenString := authHeaderParts[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Authorization token",
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Status: http.StatusUnauthorized,
				Error:  "Invalid Authorization token",
			})
			return
		}

		user := dto.User{
			ID:              claims["id"].(string),
			Name:            claims["name"].(string),
			Email:           claims["email"].(string),
			EmailVerifiedAt: claims["email_verified_at"].(string),
			PhoneNumber:     claims["phone_number"].(string),
			Country:         claims["country"].(string),
			Province:        claims["province"].(string),
			City:            claims["city"].(string),
			ZipCode:         claims["zip_code"].(string),
			CompleteAddress: claims["complete_address"].(string),
		}

		c.Set("user", user)

		c.Next()
	}
}