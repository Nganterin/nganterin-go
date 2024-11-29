package repositories

import (
	"nganterin-go/models"

	"github.com/google/uuid"
)

func (r *compRepository) RegisterPartner(data models.Partners) (*string, error) {
	data.ID = uuid.NewString()

	result := r.DB.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data.ID, nil
}

func (r *compRepository) GetPartnerDetailsByID(id string) (*models.Partners, error) {
	var partner_data models.Partners
	result := r.DB.Where("id = ?", id).First(&partner_data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &partner_data, nil
}