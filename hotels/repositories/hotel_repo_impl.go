package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

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

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]database.Hotels, *exceptions.Exception) {
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

func (r *CompRepositoriesImpl) SearchEngine(ctx *gin.Context, tx *gorm.DB, searchInput dto.HotelSearch) ([]database.Hotels, *exceptions.Exception) {
	var data []database.Hotels

	query := tx.
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
		Joins("LEFT JOIN hotel_rooms ON hotel_rooms.hotel_id = hotels.id")

	if searchInput.Keyword != "" {
		query = query.Where("LOWER(hotels.name) LIKE LOWER(?) OR LOWER(hotels.description) LIKE LOWER(?)",
			"%"+searchInput.Keyword+"%", "%"+searchInput.Keyword+"%")
	}

	if searchInput.Name != "" {
		query = query.Where("LOWER(hotels.name) LIKE LOWER(?)", "%"+searchInput.Name+"%")
	}

	if searchInput.City != "" {
		query = query.Where("EXISTS (SELECT 1 FROM hotels_locations WHERE hotels_locations.hotel_id = hotels.id AND LOWER(hotels_locations.city) LIKE LOWER(?))",
			"%"+searchInput.City+"%")
	}

	if searchInput.Country != "" {
		query = query.Where("EXISTS (SELECT 1 FROM hotels_locations WHERE hotels_locations.hotel_id = hotels.id AND LOWER(hotels_locations.country) LIKE LOWER(?))",
			"%"+searchInput.Country+"%")
	}

	if searchInput.PriceStart > 0 {
		query = query.Where("hotel_rooms.overnight_price >= ?", searchInput.PriceStart)
	}

	if searchInput.PriceEnd > 0 {
		query = query.Where("hotel_rooms.overnight_price <= ?", searchInput.PriceEnd)
	}

	// if searchInput.MinimumStars > 0 {
	// 	query = query.Where("", searchInput.MinimumStars)
	// }

	if searchInput.MinimumVisitor > 0 {
		query = query.Where("hotel_rooms.max_visitor >= ?", searchInput.MinimumVisitor)
	}

	query = query.Group("hotels.id")

	result := query.Find(&data)

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
