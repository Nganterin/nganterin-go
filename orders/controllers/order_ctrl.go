package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindByUserID(ctx *gin.Context)
}
