package services

import (
	"nganterin-go/dto"
	"nganterin-go/mapper"
)

func (s *compServices) RegisterHotel(data dto.HotelInputDTO) (*string, error) {
	model_data := mapper.MapHotelInputToModel(data)

	return s.repo.RegisterHotel(model_data)
}

func (s *compServices) GetAllHotels() (*[]dto.HotelOutputDTO, error) {
	hotels, err := s.repo.GetAllHotels()
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		output := mapper.MapHotelModelToOutput(hotels[i])
		result = append(result, output)
	}
	return &result, nil
}

func (s *compServices) SearchHotels(keyword string) (*[]dto.HotelOutputDTO, error) {
	hotels, err := s.repo.SearchHotels(keyword)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		output := mapper.MapHotelModelToOutput(hotels[i])
		result = append(result, output)
	}

	return &result, nil
}

func (s *compServices) GetHotelByID(id string) (*dto.HotelOutputDTO, error) {
	data, err := s.repo.GetHotelByID(id)
	if err != nil {
		return nil, err
	}

	result := mapper.MapHotelModelToOutput(*data)
	return &result, nil
}
