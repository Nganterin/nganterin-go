package services

import (
	"nganterin-go/dto"
	"nganterin-go/repositories"
)

type CompService interface {
	RegisterUserCredential(data dto.User) error
	LoginUserCredentials(email string, password string) (*string, error)

	SendEmail(data dto.Email) error
	VerifyUserEmail(token string) error

	RegisterPartner(data dto.Partner) error

	RegisterHotel(data dto.HotelInputDTO) (*string, error)
	GetAllHotels() (*[]dto.HotelOutputDTO, error)
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
