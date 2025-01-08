package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	FindByUserID(ctx *gin.Context, id string) ([]dto.HotelOrderDetailsOutput, *exceptions.Exception)
	FindByHotelID(ctx *gin.Context, hotelID string) ([]dto.HotelOrderDetailsOutput, *exceptions.Exception)
	FindByReservationKey(ctx *gin.Context, reservationKey string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception)
	CheckIn(ctx *gin.Context, reservationKey string) *exceptions.Exception
	CheckOut(ctx *gin.Context, reservationKey string) *exceptions.Exception
	YearlyReservationAnalytic(ctx *gin.Context, partnerID string) (*dto.HotelYearlyReservationAnalytic, *exceptions.Exception)
}
