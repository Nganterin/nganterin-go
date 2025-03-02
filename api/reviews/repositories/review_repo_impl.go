package repositories

import (
	"nganterin-go/models"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.HotelReviews) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByHotelID(ctx *gin.Context, tx *gorm.DB, id string) ([]models.HotelReviews, *exceptions.Exception) {
	var data []models.HotelReviews

	result := tx.Where("hotel_id = ?", id).Preload("User").Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}
