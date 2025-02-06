package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	CreateCredentials(ctx *gin.Context)
	CreateGoogleOAuth(ctx *gin.Context)
	LoginCredentials(ctx *gin.Context)
	LoginGoogleOAuth(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	AuthTest(ctx *gin.Context)
}
