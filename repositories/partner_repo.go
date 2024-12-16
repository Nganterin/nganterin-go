package repositories

import (
	"errors"
	"nganterin-go/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *compRepository) GetPartnerDetailsByEmail(email string) (*models.Partners, error) {
	var partner_data models.Partners
	result := r.DB.Where("email = ?", email).First(&partner_data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &partner_data, nil
}

func (r *compRepository) VerifyPartnerEmail(token string) error {
	var token_data models.Tokens
	result := r.DB.Where("token = ?", token).First(&token_data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("404")
		}
		return result.Error
	}

	partner_model := models.Partners{
		ID: token_data.UserID,
	}

	result = r.DB.Delete(&token_data)
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Model(&partner_model).Select("email_verified_at").Updates(map[string]interface{}{"email_verified_at": time.Now()})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
