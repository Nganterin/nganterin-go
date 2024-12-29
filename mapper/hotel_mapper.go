package mapper

import (
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

	"github.com/go-viper/mapstructure/v2"
)

func MapHotelInputToModel(input dto.HotelInputDTO) database.Hotels {
	var hotel database.Hotels
	
	mapstructure.Decode(input, &hotel)
	return hotel
}

func MapHotelModelToOutput(model database.Hotels) dto.HotelOutputDTO {
	var hotelOutput dto.HotelOutputDTO

	mapstructure.Decode(model, &hotelOutput)
	return hotelOutput
}

func MapHotelOrderInputToModel(input dto.HotelOrderInput) database.HotelOrders {
	var hotelOrder database.HotelOrders

	mapstructure.Decode(input, &hotelOrder)
	hotelOrder.CheckInDate = input.CheckInDate
	hotelOrder.CheckOutDate = input.CheckOutDate
	
	return hotelOrder
}

func MapHotelOrderModelToOutput(model database.HotelOrders) dto.HotelOrderDetailsOutput {
	var orderOutput dto.HotelOrderDetailsOutput

	mapstructure.Decode(model, &orderOutput)
	orderOutput.CheckInDate = model.CheckInDate
	orderOutput.CheckOutDate = model.CheckOutDate
	orderOutput.CreatedAt = model.CreatedAt
	orderOutput.UpdatedAt = model.UpdatedAt

	return orderOutput
}