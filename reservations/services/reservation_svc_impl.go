package services

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/mapper"
	"nganterin-go/models/dto"
	"nganterin-go/reservations/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
	DB   *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		repo: compRepositories,
		DB:   db,
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

func (s *CompServicesImpl) FindByReservationKey(ctx *gin.Context, reservationKey string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception) {
	data, err := s.repo.FindByReservationKey(ctx, s.DB, reservationKey)
	if err != nil {
		return nil, err
	}

	result := mapper.MapHotelOrderModelToOutput(*data)

	return &result, nil
}

func (s *CompServicesImpl) CheckIn(ctx *gin.Context, reservationKey string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.repo.FindByReservationKey(ctx, tx, reservationKey)
	if err != nil {
		return err
	}

	if data.HotelReservations.ReservationStatus != "confirmed" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrAlreadyCheckedIn)
	}

	currentTime := time.Now()

	if !(currentTime.After(data.CheckInDate) && currentTime.Before(data.CheckOutDate)) {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidDate)
	}

	return s.repo.CheckIn(ctx, tx, data.ID)
}

func (s *CompServicesImpl) CheckOut(ctx *gin.Context, reservationKey string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.repo.FindByReservationKey(ctx, tx, reservationKey)
	if err != nil {
		return err
	}

	if data.HotelReservations.ReservationStatus != "CheckIn" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrNotCheckedInYet)
	}

	currentTime := time.Now()

	if !currentTime.After(data.CheckInDate) {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidDate)
	}

	return s.repo.CheckOut(ctx, tx, data.ID)
}