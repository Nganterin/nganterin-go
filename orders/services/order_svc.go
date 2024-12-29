package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.HotelOrderInput) (*dto.HotelOrderOutput, *exceptions.Exception)
	FindByID(ctx *gin.Context, id string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception)
	FindByUserID(ctx *gin.Context, id string) ([]dto.HotelOrderDetailsOutput, *exceptions.Exception)
}
