package middleware

import (
	"net/http"
	"net/url"
	"nganterin-go/dto"
	"nganterin-go/models"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting secret"})
			return
		}

		var secretKey = []byte(secret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		tokenString := authHeaderParts[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		user := dto.User{
			ID:              claims["id"].(string),
			Email:           claims["email"].(string),
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

func ClientTracker(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		userAgent := c.Request.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		name, version := ua.Browser()

		referer := c.Request.Referer()

		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery

		fullURL := url.URL{
			Path:     path,
			RawQuery: rawQuery,
		}

		data := models.Client{
			IP:      clientIP,
			Browser: name,
			Version: version,
			OS:      ua.OS(),
			Device:  ua.Platform(),
			Origin:  referer,
			API:     fullURL.String(),
		}

		result := db.Create(&data)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.Next()
	}
}

func NoCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Next()
	}
}
