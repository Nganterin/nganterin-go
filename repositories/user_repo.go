package repositories

import (
	"errors"
	"nganterin-go/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *compRepository) RegisterUserCredential(data models.Users) (*string, error) {
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

func (r *compRepository) GetUserDetailsByID(id string) (*models.Users, error) {
	var user_data models.Users
	result := r.DB.Where("id = ?", id).First(&user_data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user_data, nil
}

func (r *compRepository) GetUserDetailsByEmail(email string) (*models.Users, error) {
	var user_data models.Users
	result := r.DB.Where("email = ?", email).First(&user_data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user_data, nil
}

func (r *compRepository) VerifyUserEmail(token string) error {
	var token_data models.Tokens
	result := r.DB.Where("token = ?", token).First(&token_data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("404")
		}
		return result.Error
	}

	user_model := models.Users{
		ID: token_data.UserID,
	}

	result = r.DB.Delete(&token_data)
	if result.Error != nil {
		return result.Error
	}


	result = r.DB.Model(&user_model).Select("email_verified_at").Updates(map[string]interface{}{"email_verified_at": time.Now()})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
