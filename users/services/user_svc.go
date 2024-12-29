package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	CreateCredentials(ctx *gin.Context, data dto.User) *exceptions.Exception
	LoginCredentials(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception)
	VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception
	
	CreateGoogleOAuth(ctx *gin.Context, data dto.UserGoogle) (*string, *exceptions.Exception)
	LoginGoogleOAuth(ctx *gin.Context, email string, googleSUB string) (*string, *exceptions.Exception)
}
