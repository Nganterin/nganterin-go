package repositories

import (
	"errors"
	"nganterin-go/models"
	"strings"
)

func (r *compRepository) RegisterHotelOrder(data models.HotelOrders) error {
	result := r.DB.Create(&data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return errors.New("409")
		}
		return result.Error
	}

	return nil
}

func (r *compRepository) GetHotelOrderByID(id string) (*models.HotelOrders, error) {
	var data models.HotelOrders

	result := r.DB.
		Preload("HotelReservations").
		Where("id = ?", id).
		First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (s *compRepository) UpdateHotelOrderPaymentStatus(id string, status string) error {
	var data models.HotelOrders

	result := s.DB.Model(&data).Where("id = ?", id).Update("payment_status",
		status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *compRepository) GetAllHotelOrderByUserID(id string) ([]models.HotelOrders, error) {
	var data []models.HotelOrders

	result := r.DB.
		Preload("HotelReservations").
		Where("user_id = ?", id).
		Order("created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
