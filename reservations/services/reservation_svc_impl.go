package services

import (
	"fmt"
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
		output := mapper.MapHotelOrderModelToOutput(item)
		output.TotalDays = helpers.GetDaysFromCheckInCheckOut(item.CheckInDate, item.CheckOutDate)

		result = append(result, output)
	}

	return result, nil
}

func (s *CompServicesImpl) FindByReservationKey(ctx *gin.Context, reservationKey string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception) {
	data, err := s.repo.FindByReservationKey(ctx, s.DB, reservationKey)
	if err != nil {
		return nil, err
	}

	result := mapper.MapHotelOrderModelToOutput(*data)
	result.TotalDays = helpers.GetDaysFromCheckInCheckOut(data.CheckInDate, data.CheckOutDate)

	return &result, nil
}

func (s *CompServicesImpl) CheckIn(ctx *gin.Context, reservationKey string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.repo.FindByReservationKey(ctx, tx, reservationKey)
	if err != nil {
		return err
	}

	if data.HotelReservations.ReservationStatus != "Confirmed" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrAlreadyCheckedIn)
	}

	currentTime := time.Now()

	if !(currentTime.After(data.CheckInDate) && currentTime.Before(data.CheckOutDate)) {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidDate)
	}

	return s.repo.CheckIn(ctx, tx, reservationKey)
}

func (s *CompServicesImpl) CheckOut(ctx *gin.Context, reservationKey string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.repo.FindByReservationKey(ctx, tx, reservationKey)
	if err != nil {
		return err
	}

	if data.HotelReservations.ReservationStatus != "CheckedIn" {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrNotCheckedInYet)
	}

	currentTime := time.Now()

	if !currentTime.After(data.CheckInDate) {
		return exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidDate)
	}

	return s.repo.CheckOut(ctx, tx, reservationKey)
}

func (s *CompServicesImpl) YearlyReservationAnalytic(ctx *gin.Context, partnerID string) (*dto.HotelYearlyReservationAnalytic, *exceptions.Exception) {
	reservationData, err := s.repo.FindLast12MonthReservationCount(ctx, s.DB, partnerID)
	if err != nil {
		return nil, err
	}

	lastMonth := reservationData[len(reservationData)-2]
	currentMonth := reservationData[len(reservationData)-1]
	difference := currentMonth.ReservationCount - lastMonth.ReservationCount
	percentageChange := 0.0

	if lastMonth.ReservationCount == 0 {
		percentageChange = 100 * float64(difference)
	} else {
		percentageChange = (float64(difference) / float64(lastMonth.ReservationCount)) * 100
	}

	result := dto.HotelYearlyReservationAnalytic{
		Period:             fmt.Sprintf("%s - %s", lastMonth.MonthYear, currentMonth.MonthYear),
		TrendPercentage:    float32(percentageChange),
		MonthlyReservation: reservationData,
	}

	return &result, nil
}
