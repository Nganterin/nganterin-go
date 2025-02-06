package services

import (
	"net/http"
	"nganterin-go/api/reviews/repositories"
	"nganterin-go/api/reviews/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"
	"nganterin-go/pkg/mapper"

	orderRepo "nganterin-go/api/orders/repositories"
	reservationRepo "nganterin-go/api/reservations/repositories"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo            repositories.CompRepositories
	orderRepo       orderRepo.CompRepositories
	reservationRepo reservationRepo.CompRepositories
	DB              *gorm.DB
	validate        *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, orderRepo orderRepo.CompRepositories, reservationRepo reservationRepo.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:            compRepositories,
		orderRepo:       orderRepo,
		reservationRepo: reservationRepo,
		DB:              db,
		validate:        validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.HotelReviewInput) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderData, err := s.orderRepo.FindByID(ctx, tx, data.HotelOrdersID)
	if err != nil {
		return err
	}

	if orderData.UserID != data.UserID {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden)
	}

	if orderData.HotelReservations.ReservationStatus != "CheckedOut" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrNotCheckedOutYet)
	}

	input := mapper.MapHotelReviewInputToModel(data)
	input.HotelID = orderData.HotelID

	err = s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	err = s.reservationRepo.Reviewed(ctx, tx, orderData.ID)
	if err != nil {
		return err
	}

	return nil
}
