package services

import (
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.HotelReviewInput) *exceptions.Exception
}
