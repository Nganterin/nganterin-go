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
