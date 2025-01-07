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
	
func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data database.HotelOrders) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.HotelOrders, *exceptions.Exception) {
	var data database.HotelOrders

	result := tx.
		Preload("HotelReservations").
		Where("id = ?", id).
		First(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}

func (s *CompRepositoriesImpl) UpdatePaymentStatus(ctx *gin.Context, tx *gorm.DB, id string, status string) *exceptions.Exception {
	var data database.HotelOrders

	result := tx.Model(&data).Where("id = ?", id).Update("payment_status",
		status)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelOrders, *exceptions.Exception) {
	var data []database.HotelOrders

	result := tx.
		Preload("HotelReservations").
		Preload("Hotel").
		Preload("Hotel.HotelsLocation").
		Preload("HotelRoom").
		Preload("HotelRoom.HotelRoomPhotos").
		Where("user_id = ?", id).
		Order("created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindByPartnerID(ctx *gin.Context, tx *gorm.DB, partnerID string) ([]database.HotelOrders, *exceptions.Exception) {
	var data []database.HotelOrders

	result := tx.
		Joins("JOIN hotels ON hotels.id = hotel_orders.hotel_id").
		Preload("Hotel").
		Preload("HotelRoom").
		Where("hotels.partner_id = ?", partnerID).
		Order("hotel_orders.created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}