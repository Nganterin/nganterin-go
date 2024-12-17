package repositories

import (
	"nganterin-go/models"

	"github.com/google/uuid"
)

func (r *compRepository) FileUpload(data models.Files) (*models.Files, error) {
	data.ID = uuid.NewString()

	result := r.DB.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}