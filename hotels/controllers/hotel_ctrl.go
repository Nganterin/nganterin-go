package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	SearchEngine(ctx *gin.Context)
	FindByPartnerID(ctx *gin.Context)
}
