package services

import (
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Notification(ctx *gin.Context, data dto.MidtransNotification) *exceptions.Exception
}
