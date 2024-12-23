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
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
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
		return nil, result.Error
	}

	return data, nil
}

func (r *compRepository) GetHotelByID(id string) (*models.Hotels, error) {
	var data models.Hotels
	result := r.DB.
		Preload("HotelRooms").
		Preload("HotelRooms.HotelRoomPhotos").
		Preload("HotelsLocation").
		Preload("HotelPhotos").
		Preload("HotelFacilities").
		Where("id = ?", id).
		First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (r *compRepository) GetHotelRoomByID(id uint) (*models.HotelRooms, error) {
	var data models.HotelRooms
	result := r.DB.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}