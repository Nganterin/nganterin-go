package repositories

import (
	"errors"
	"nganterin-go/models"
	"strings"
)

func (r *compRepository) RegisterHotelReservation(data models.HotelReservations) error {
	result := r.DB.Create(&data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return errors.New("409")
		}
		return result.Error
	}

	return nil
}