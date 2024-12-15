package repositories

import (
	"errors"
	"nganterin-go/models"
	"strings"

	"github.com/google/uuid"
)

func (r *compRepository) RegisterHotel(data models.Hotels) (*string, error) {
	data.ID = uuid.NewString()

	result := r.DB.Create(&data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return nil, errors.New("409")
		}
		return nil, result.Error
	}

	return &data.ID, nil
}

func (r *compRepository) GetAllHotels() ([]models.Hotels, error) {
	var data []models.Hotels

	result := r.DB.
		Preload("HotelDetails").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *compRepository) SearchHotels(keyword string) ([]models.Hotels, error) {
	var data []models.Hotels
	searchKeyword := "%" + keyword + "%"

	result := r.DB.
		Preload("HotelDetails").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Where("LOWER(name) LIKE LOWER(?) OR LOWER(description) LIKE LOWER(?) OR "+
			"EXISTS (SELECT 1 FROM hotels_locations WHERE hotels_locations.hotel_id = hotels.id AND "+
			"(LOWER(hotels_locations.city) LIKE LOWER(?) OR LOWER(hotels_locations.state) LIKE LOWER(?) OR LOWER(hotels_locations.country) LIKE LOWER(?)))",
			searchKeyword, searchKeyword, searchKeyword, searchKeyword, searchKeyword).
		Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
