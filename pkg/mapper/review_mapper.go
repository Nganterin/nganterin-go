package mapper

import (
	"nganterin-go/api/reviews/dto"
	"nganterin-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapHotelReviewInputToModel(input dto.HotelReviewInput) models.HotelReviews {
	var hotelReview models.HotelReviews

	mapstructure.Decode(input, &hotelReview)

	return hotelReview
}
