package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data database.HotelReservations) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelOrders, *exceptions.Exception) {
	var data []database.HotelOrders

	result := tx.
		Joins("JOIN hotel_reservations ON hotel_reservations.hotel_orders_id = hotel_orders.id").
		Preload("HotelReservations").
		Preload("Hotel").
		Preload("Hotel.HotelsLocation").
		Preload("HotelRoom").
		Preload("HotelRoom.HotelRoomPhotos").
		Where("hotel_orders.user_id = ?", id).
		Order("hotel_orders.created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}