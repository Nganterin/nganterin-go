package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/hotels/repositories"
	"nganterin-go/mapper"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
	DB   *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB) CompService {
	return &CompServicesImpl{
		repo: compRepositories,
		DB:   db,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.HotelInputDTO) (*string, *exceptions.Exception) {
	model_data := mapper.MapHotelInputToModel(data)
	return s.repo.Create(ctx, s.DB, model_data)
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) (*[]dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		pricingStart := s.GetPricingStartHotelRooms(ctx, hotels[i].HotelRooms)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
		result = append(result, output)
	}
	return &result, nil
}

func (s *CompServicesImpl) FindByKeyword(ctx *gin.Context, keyword string) (*[]dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.FindByKeyword(ctx, s.DB, keyword)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOutputDTO
	for i := range hotels {
		pricingStart := s.GetPricingStartHotelRooms(ctx, hotels[i].HotelRooms)

		output := mapper.MapHotelModelToOutput(hotels[i])
		output.PricingStart = pricingStart
		result = append(result, output)
	}

	return &result, nil
}

func (s *CompServicesImpl) FindByID(ctx *gin.Context, id string) (*dto.HotelOutputDTO, *exceptions.Exception) {
	hotels, err := s.repo.FindByID(ctx, s.DB,id)
	if err != nil {
		return nil, err
	}

	pricingStart := s.GetPricingStartHotelRooms(ctx, hotels.HotelRooms)
	result := mapper.MapHotelModelToOutput(*hotels)
	result.PricingStart = pricingStart
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
