package services

import (
	"nganterin-go/dto"
	"nganterin-go/repositories"
)

type CompService interface{
	RegisterUserCredential(data dto.User) error
	LoginUserCredentials(email string, password string) (*string, error)

	SendEmail(data dto.Email) error
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
