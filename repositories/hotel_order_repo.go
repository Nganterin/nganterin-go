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