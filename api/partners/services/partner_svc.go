package services

import (
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Partner) *exceptions.Exception
	Login(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception)
	VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception
	ApprovalCheck(ctx *gin.Context, id string) (*string, *exceptions.Exception)
}
