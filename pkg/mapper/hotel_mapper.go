package mapper

import (
	"nganterin-go/models"
	"nganterin-go/models/dto"

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

func MapHotelReviewInputToModel(input dto.HotelReviewInput) models.HotelReviews {
	var hotelReview models.HotelReviews

	mapstructure.Decode(input, &hotelReview)

	return hotelReview
}

func MapHotelReviewModelToOutput(model models.HotelReviews) dto.HotelReviewOutput {
	var reviewOutput dto.HotelReviewOutput

	mapstructure.Decode(model, &reviewOutput)
	reviewOutput.CreatedAt = model.CreatedAt
	reviewOutput.UpdatedAt = model.UpdatedAt

	return reviewOutput
}
