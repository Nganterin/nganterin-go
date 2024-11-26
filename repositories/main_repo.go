package repositories

import (
	"nganterin-go/config"
	"nganterin-go/models"

	"gorm.io/gorm"
)

type CompRepository interface {
	RegisterUserCredential(data models.Users) (string, error)
	LoginUserCredentials(email string) (*models.Users, error)
}

type compRepository struct {
	DB *gorm.DB
}

func NewComponentRepository(DB *gorm.DB) *compRepository {
	db := config.InitDB()

	return &compRepository{
		DB: db,
	}
}
