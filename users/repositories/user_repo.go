package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	CreateCredentials(ctx *gin.Context, tx *gorm.DB, data database.Users) (*string, *exceptions.Exception)
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.Users, *exceptions.Exception)
	FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*database.Users, *exceptions.Exception)

	CreateVerificationToken(ctx *gin.Context, tx *gorm.DB, id string) (*string, *exceptions.Exception) 
	VerifyEmail(ctx *gin.Context, tx *gorm.DB, token string) *exceptions.Exception
}
