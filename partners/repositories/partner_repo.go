package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data database.Partners) (*string, *exceptions.Exception)
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.Partners, *exceptions.Exception)
	FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*database.Partners, *exceptions.Exception)
	VerifyEmail(ctx *gin.Context, tx *gorm.DB, token string) *exceptions.Exception
	CreateVerificationToken(ctx *gin.Context, tx *gorm.DB, id string) (*string, *exceptions.Exception)
}
