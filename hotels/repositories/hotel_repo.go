package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data database.Hotels) (*string, *exceptions.Exception)
	FindAll(ctx *gin.Context, tx *gorm.DB, ) ([]database.Hotels, *exceptions.Exception)
	SearchEngine(ctx *gin.Context, tx *gorm.DB, searchInput dto.HotelSearch) ([]database.Hotels, *exceptions.Exception)
	FindByID(ctx *gin.Context, tx *gorm.DB, id string)(*database.Hotels, *exceptions.Exception)
	FindRoomByID(ctx *gin.Context, tx *gorm.DB, id uint) (*database.HotelRooms, *exceptions.Exception)
}
