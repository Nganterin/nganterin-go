package services

import (
	"nganterin-go/api/reviews/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.HotelReviewInput) *exceptions.Exception
}
