package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	FindByUserID(ctx *gin.Context) 
}