package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Notification(ctx *gin.Context, data dto.MidtransNotification) *exceptions.Exception
}