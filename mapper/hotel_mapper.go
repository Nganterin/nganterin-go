package mapper

import (
	"nganterin-go/dto"
	"nganterin-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapHotelInputToModel(input dto.HotelInputDTO) models.Hotels {
	var hotel models.Hotels
	
	mapstructure.Decode(input, &hotel)
	return hotel
}

func MapHotelModelToOutput(model models.Hotels) dto.HotelOutputDTO {
	var hotelOutput dto.HotelOutputDTO

	mapstructure.Decode(model, &hotelOutput)
	return hotelOutput
}

func MapHotelOrderInputToModel(input dto.HotelOrderInput) models.HotelOrders {
	var hotelOrder models.HotelOrders

	mapstructure.Decode(input, &hotelOrder)
	hotelOrder.CheckInDate = input.CheckInDate
	hotelOrder.CheckOutDate = input.CheckOutDate
	
	return hotelOrder
}