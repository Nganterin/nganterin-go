package hotels

import (
	"errors"
	"nganterin-go/config"
	"nganterin-go/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HotelRepository interface {
	CreateHotel(data models.Hotels) (*string, error)
}

type hotelRepository struct {
	DB *gorm.DB
}

func NewHotelRepository(DB *gorm.DB) *hotelRepository {
	db := config.InitDB()
	return &hotelRepository{
		DB: db,
	}
}

func (r *hotelRepository) CreateHotel(data models.Hotels) (*string, error) {
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
