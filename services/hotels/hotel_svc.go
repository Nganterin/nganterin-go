package hotels

import (
	"nganterin-go/dto"
	"nganterin-go/mapper"
	"nganterin-go/repositories/hotels"
)

type HotelService interface {
	CreateHotel(data dto.HotelInputDTO) (*string, error)
}

type hotelService struct {
	repo hotels.HotelRepository
}

func NewHotelService(repo hotels.HotelRepository) *hotelService {
	return &hotelService{
		repo: repo,
	}
}

func (s *hotelService) CreateHotel(data dto.HotelInputDTO) (*string, error) {
	model_data := mapper.MapHotelInputToModel(data)

	return s.repo.CreateHotel(model_data)
}
