package repositories

import (
	"errors"
	"nganterin-go/models"
	"strings"
)

func (r *compRepository) RegisterUserCredential(data models.Users) (string, error) {
	result := r.DB.Create(&data)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return "", errors.New("409")
		}
		return "", result.Error
	}

	return data.ID, nil
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

