package services

import (
	"nganterin-go/api/orders/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.HotelOrderInput) (*dto.HotelOrderOutput, *exceptions.Exception)
	FindByID(ctx *gin.Context, id string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception)
	FindByUserID(ctx *gin.Context, id string) ([]dto.HotelOrderDetailsOutput, *exceptions.Exception)
	YearlyOrderAnalytic(ctx *gin.Context, partnerID string) (*dto.HotelYearlyOrderAnalytic, *exceptions.Exception)
}
