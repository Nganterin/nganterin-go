package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Partner) *exceptions.Exception
	Login(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception)
	VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception
}
