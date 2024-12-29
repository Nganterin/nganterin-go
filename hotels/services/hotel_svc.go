package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompService interface {
	Create(ctx *gin.Context, data dto.HotelInputDTO) (*string, *exceptions.Exception)
	FindAll(ctx *gin.Context, ) (*[]dto.HotelOutputDTO, *exceptions.Exception)
	FindByID(ctx *gin.Context, id string) (*dto.HotelOutputDTO, *exceptions.Exception)
	FindByKeyword(ctx *gin.Context, keyword string) (*[]dto.HotelOutputDTO, *exceptions.Exception)
}
