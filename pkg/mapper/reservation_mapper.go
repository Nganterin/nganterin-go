package mapper

import (
	"nganterin-go/api/reservations/dto"
	"nganterin-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapHotelOrderReservationModelToOutput(model models.HotelOrders) dto.HotelOrderDetailsOutput {
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