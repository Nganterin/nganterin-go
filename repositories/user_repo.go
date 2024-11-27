package repositories

import (
	"errors"
	"log"
	"nganterin-go/models"
	"strings"

	"github.com/google/uuid"
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

	log.Println("INSERTED USER ID: ", data.ID)

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
