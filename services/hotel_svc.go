package services

import (
	"nganterin-go/dto"
	"nganterin-go/mapper"
)

func (s *compServices) RegisterHotel(data dto.HotelInputDTO) (*string, error) {
	model_data := mapper.MapHotelInputToModel(data)

	return s.repo.RegisterHotel(model_data)
}
