package repositories

import (
	"nganterin-go/models"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Users) (*string, *exceptions.Exception)
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*models.Users, *exceptions.Exception)
	FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*models.Users, *exceptions.Exception)

	CreateVerificationToken(ctx *gin.Context, tx *gorm.DB, id string) (*string, *exceptions.Exception)
	VerifyEmail(ctx *gin.Context, tx *gorm.DB, token string) *exceptions.Exception
}
