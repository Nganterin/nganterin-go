package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Notification(ctx *gin.Context)
}