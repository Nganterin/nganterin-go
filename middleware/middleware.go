package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/url"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"gorm.io/gorm"
)

func GzipResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		gzipWriter := gzip.NewWriter(c.Writer)
		defer gzipWriter.Close()

		wrappedWriter := &gzipResponseWriter{
			ResponseWriter: c.Writer,
			Writer:         gzipWriter,
		}

		c.Writer = wrappedWriter
		c.Writer.Header().Set("Content-Encoding", "gzip")
		c.Writer.Header().Set("Vary", "Accept-Encoding")

		c.Next()
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (g *gzipResponseWriter) Write(data []byte) (int, error) {
	return g.Writer.Write(data)
}

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

		data := database.Client{
			IP:      clientIP,
			Browser: name,
			Version: version,
			OS:      ua.OS(),
			Device:  ua.Platform(),
			Origin:  referer,
			API:     fullURL.String(),
		}

		go db.Create(&data)
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
