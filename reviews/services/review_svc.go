package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)


type CompServices interface {
	Create(ctx *gin.Context, data dto.HotelReviewInput) *exceptions.Exception
}