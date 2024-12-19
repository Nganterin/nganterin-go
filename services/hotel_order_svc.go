package services

import (
	"errors"
	"nganterin-go/dto"
	"nganterin-go/mapper"
)


func (s *compServices) RegisterHotelOrder(data dto.HotelOrderInput) error {
	roomData, err := s.repo.GetHotelRoomByID(data.HotelRoomID)
	if err != nil {
		return errors.New("400")
	}

	input := mapper.MapHotelOrderInputToModel(data)
	
	duration := input.CheckOutDate.Sub(input.CheckInDate)
	days := int(duration.Hours() / 24)

	input.TotalPrice = roomData.OvernightPrice * int64(days)

	_, err = s.repo.RegisterHotelOrder(input)
	if err != nil {
		return err
	}

	return nil
}