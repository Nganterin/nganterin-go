package services

import (
	"nganterin-go/dto"
	"nganterin-go/repositories"

	"github.com/midtrans/midtrans-go/snap"
)

type CompService interface {
	RegisterUserCredential(data dto.User) error
	LoginUserCredentials(email string, password string) (*string, error)
	VerifyUserEmail(token string) error

	RegisterPartner(data dto.Partner) error
	LoginPartner(email string, password string) (*string, error)
	VerifyPartnerEmail(token string) error
	
	SendEmail(data dto.Email) error

	RegisterHotel(data dto.HotelInputDTO) (*string, error)
	GetAllHotels() (*[]dto.HotelOutputDTO, error)
	SearchHotels(keyword string) (*[]dto.HotelOutputDTO, error)
	GetHotelByID(id string) (*dto.HotelOutputDTO, error)

	FileUpload(file []byte, data dto.FilesInputDTO) (*dto.FilesOutputDTO, error)

	RegisterHotelOrder(data dto.HotelOrderInput) (*snap.Response, error)

	MidtransNotification(data dto.MidtransNotification) error
}

type compServices struct {
	repo repositories.CompRepository
}

func NewService(r repositories.CompRepository) *compServices {
	return &compServices{
		repo: r,
	}
}
