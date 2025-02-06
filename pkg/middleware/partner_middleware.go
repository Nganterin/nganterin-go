package middleware

import (
	"net/http"
	"nganterin-go/partners/dto"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PartnerAuthMiddleware() gin.HandlerFunc {
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

		isPartner, ok := claims["is_partner"].(bool)
		if !ok || !isPartner {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
				Status: http.StatusForbidden,
				Error:  "Access restricted to partners",
			})
			return
		}

		partner := dto.Partner{
			ID:             claims["id"].(string),
			Name:           claims["name"].(string),
			Email:          claims["email"].(string),
			CompanyName:    claims["company_name"].(string),
			Owner:          claims["owner"].(string),
			CompanyField:   claims["company_field"].(string),
			CompanyEmail:   claims["company_email"].(string),
			CompanyAddress: claims["company_address"].(string),
		}

		c.Set("partner", partner)

		c.Next()
	}
}