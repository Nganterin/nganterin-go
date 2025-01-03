package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/hotels/repositories"
	"nganterin-go/mapper"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompService {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.HotelInputDTO) (*string, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	model_data := mapper.MapHotelInputToModel(data)
	return s.repo.Create(ctx, tx, model_data)
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) (*[]dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		pricingStart := s.GetPricingStartHotelRooms(ctx, hotels[i].HotelRooms)
		averageRating := s.GetReviewAverageRating(ctx, hotels[i].HotelReviews)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
		output.Rating = averageRating
		result = append(result, output)
	}
	return &result, nil
}

func (s *CompServicesImpl) SearchEngine(ctx *gin.Context, searchInput dto.HotelSearch) (*[]dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.SearchEngine(ctx, s.DB, searchInput)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		pricingStart := s.GetPricingStartHotelRooms(ctx, hotels[i].HotelRooms)
		averageRating := s.GetReviewAverageRating(ctx, hotels[i].HotelReviews)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
		output.Rating = averageRating
		result = append(result, output)
	}

	return &result, nil
}

func (s *CompServicesImpl) FindByID(ctx *gin.Context, id string) (*dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.FindByID(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	pricingStart := s.GetPricingStartHotelRooms(ctx, hotels.HotelRooms)
	averageRating := s.GetReviewAverageRating(ctx, hotels.HotelReviews)
	reviewStatistic := s.GetReviewStatistics(ctx, hotels.HotelReviews)

	result := mapper.MapHotelModelToOutput(*hotels)
	result.PricingStart = pricingStart
	result.Rating = averageRating
	result.ReviewStatistic = reviewStatistic
	return &result, nil
}

func (s *CompServicesImpl) GetPricingStartHotelRooms(ctx *gin.Context, data []database.HotelRooms) int64 {
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

func (s *CompServicesImpl) GetReviewAverageRating(ctx *gin.Context, data []database.HotelReviews) dto.HotelAverageRating {
	var totalCleanliness, totalComfort, totalServiceQuality, totalFacilities, totalValueForMoney, totalRating int

	for _, review := range data {
		totalCleanliness += review.Cleanliness
		totalComfort += review.Comfort
		totalServiceQuality += review.ServiceQuality
		totalFacilities += review.Facilities
		totalValueForMoney += review.ValueForMoney
		totalRating += review.Rating
	}

	result := dto.HotelAverageRating{}

	if len(data) > 0 {
		count := float32(len(data))
		result.Cleanliness = float32(totalCleanliness) / count
		result.Comfort = float32(totalComfort) / count
		result.ServiceQuality = float32(totalServiceQuality) / count
		result.Facilities = float32(totalFacilities) / count
		result.ValueForMoney = float32(totalValueForMoney) / count
		result.Rating = float32(totalRating) / count
	}

	return result
}

func (s *CompServicesImpl) GetReviewStatistics(ctx *gin.Context, data []database.HotelReviews) dto.HotelReviewStatistic {
	var totalReviews int
	var totalRating int
	var count1, count2, count3, count4, count5 int

	for _, review := range data {
		totalReviews++
		totalRating += review.Rating

		switch review.Rating {
		case 5:
			count5++
		case 4:
			count4++
		case 3:
			count3++
		case 2:
			count2++
		case 1:
			count1++
		}
	}

	result := dto.HotelReviewStatistic{
		TotalReviews: totalReviews,
	}

	if totalReviews > 0 {
		result.AverageRating = float32(totalRating) / float32(totalReviews)
		result.Percentage5 = (count5 * 100) / totalReviews
		result.Percentage4 = (count4 * 100) / totalReviews
		result.Percentage3 = (count3 * 100) / totalReviews
		result.Percentage2 = (count2 * 100) / totalReviews
		result.Percentage1 = (count1 * 100) / totalReviews
	}

	return result
}
