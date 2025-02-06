package mapper

import (
	"nganterin-go/models"
	"nganterin-go/api/hotels/dto"

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

func MapHotelReviewModelToOutput(model models.HotelReviews) dto.HotelReviewOutput {
	var reviewOutput dto.HotelReviewOutput

	mapstructure.Decode(model, &reviewOutput)
	reviewOutput.CreatedAt = model.CreatedAt
	reviewOutput.UpdatedAt = model.UpdatedAt

	return reviewOutput
}
