package mapper

import (
	"nganterin-go/api/orders/dto"
	"nganterin-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapHotelOrderInputToModel(input dto.HotelOrderInput) models.HotelOrders {
	var hotelOrder models.HotelOrders

	mapstructure.Decode(input, &hotelOrder)
	hotelOrder.CheckInDate = input.CheckInDate
	hotelOrder.CheckOutDate = input.CheckOutDate

	return hotelOrder
}

func MapHotelOrderModelToOutput(model models.HotelOrders) dto.HotelOrderDetailsOutput {
	var orderOutput dto.HotelOrderDetailsOutput

	mapstructure.Decode(model, &orderOutput)
	orderOutput.CheckInDate = model.CheckInDate
	orderOutput.CheckOutDate = model.CheckOutDate
	orderOutput.CreatedAt = model.CreatedAt
	orderOutput.UpdatedAt = model.UpdatedAt
	orderOutput.HotelReservations.CreatedAt = model.HotelReservations.CreatedAt
	orderOutput.HotelReservations.UpdatedAt = model.HotelReservations.UpdatedAt

	return orderOutput
}