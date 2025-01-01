package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	FindByUserID(ctx *gin.Context)
	FindByReservationKey(ctx *gin.Context)
	CheckIn(ctx *gin.Context)
	CheckOut(ctx *gin.Context)
}