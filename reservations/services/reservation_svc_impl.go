package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/mapper"
	"nganterin-go/models/dto"
	"nganterin-go/reservations/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo         repositories.CompRepositories
	DB           *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		repo:         compRepositories,
		DB:           db,
	}
}

func (s *CompServicesImpl) FindByUserID(ctx *gin.Context, id string) ([]dto.HotelOrderDetailsOutput, *exceptions.Exception) {
	data, err := s.repo.FindByUserID(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOrderDetailsOutput
	for _, item := range data {
		result = append(result, mapper.MapHotelOrderModelToOutput(item))
	}

	return result, nil
}