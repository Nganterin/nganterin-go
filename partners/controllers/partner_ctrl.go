package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	ApprovalCheck(ctx *gin.Context)
}
