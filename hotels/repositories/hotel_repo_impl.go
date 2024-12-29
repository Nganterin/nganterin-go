package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data database.Hotels) (*string, *exceptions.Exception) {
	data.ID = uuid.NewString()

	result := tx.Create(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data.ID, nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB, ) ([]database.Hotels, *exceptions.Exception) {
	var data []database.Hotels

	result := tx.
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
		Find(&data)

	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindByKeyword(ctx *gin.Context, tx *gorm.DB, keyword string) ([]database.Hotels, *exceptions.Exception) {
	var data []database.Hotels
	searchKeyword := "%" + keyword + "%"

	result := tx.
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
		Where("LOWER(name) LIKE LOWER(?) OR LOWER(description) LIKE LOWER(?) OR "+
			"EXISTS (SELECT 1 FROM hotels_locations WHERE hotels_locations.hotel_id = hotels.id AND "+
			"(LOWER(hotels_locations.city) LIKE LOWER(?) OR LOWER(hotels_locations.state) LIKE LOWER(?) OR LOWER(hotels_locations.country) LIKE LOWER(?)))",
			searchKeyword, searchKeyword, searchKeyword, searchKeyword, searchKeyword).
		Find(&data)

	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.Hotels, *exceptions.Exception) {
	var data database.Hotels
	result := tx.
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
		Where("id = ?", id).
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindRoomByID(ctx *gin.Context, tx *gorm.DB, id uint) (*database.HotelRooms, *exceptions.Exception) {
	var data database.HotelRooms
	result := tx.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}