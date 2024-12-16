package services

import (
	"nganterin-go/dto"
	"nganterin-go/mapper"
	"nganterin-go/models"
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
		pricingStart := s.GetPricingStartHotelDetails(hotels[i].HotelDetails)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
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
		pricingStart := s.GetPricingStartHotelDetails(hotels[i].HotelDetails)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
		result = append(result, output)
	}

	return &result, nil
}

func (s *compServices) GetHotelByID(id string) (*dto.HotelOutputDTO, error) {
	hotels, err := s.repo.GetHotelByID(id)
	if err != nil {
		return nil, err
	}

	pricingStart := s.GetPricingStartHotelDetails(hotels.HotelDetails)
	result := mapper.MapHotelModelToOutput(*hotels)
	result.PricingStart = pricingStart
	return &result, nil
}

func (s *compServices) GetPricingStartHotelDetails(data []models.HotelDetails) (int64) {
	var pricingStart int64
	
	if len(data) > 0 {
		pricingStart = data[0].OvernightPrice
		for _, detail := range data {
			if detail.OvernightPrice < pricingStart {
				pricingStart = detail.OvernightPrice
			}
		}
	}


	return pricingStart
}
