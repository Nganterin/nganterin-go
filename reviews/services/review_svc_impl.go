package services

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/mapper"
	"nganterin-go/models/dto"
	"nganterin-go/reviews/repositories"

	orderRepo "nganterin-go/orders/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo      repositories.CompRepositories
	orderRepo orderRepo.CompRepositories
	DB        *gorm.DB
	validate  *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, orderRepo orderRepo.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:      compRepositories,
		orderRepo: orderRepo,
		DB:        db,
		validate:  validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.HotelReviewInput) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	orderData, err := s.orderRepo.FindByID(ctx, s.DB, data.HotelOrdersID)
	if err != nil {
		return err
	}

	if orderData.UserID != data.UserID || orderData.HotelReservations.ReservationStatus != "CheckedOut" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden)
	}

	input := mapper.MapHotelReviewInputToModel(data)
	input.HotelID = orderData.HotelID

	err = s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}
