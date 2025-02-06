package services

import (
	orderRepo "nganterin-go/api/orders/repositories"
	reservationRepo "nganterin-go/api/reservations/repositories"
	"nganterin-go/models"
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	orderRepo       orderRepo.CompRepositories
	reservationRepo reservationRepo.CompRepositories
	DB              *gorm.DB
}

func NewComponentServices(compOrderRepo orderRepo.CompRepositories, compReservationRepo reservationRepo.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		orderRepo:       compOrderRepo,
		reservationRepo: compReservationRepo,
		DB:              db,
	}
}

func (s *CompServicesImpl) Notification(ctx *gin.Context, data dto.MidtransNotification) *exceptions.Exception {
	if data.TransactionStatus == "deny" || data.TransactionStatus == "cancel" || data.TransactionStatus == "expire" || data.TransactionStatus == "failure" {
		err := s.orderRepo.UpdatePaymentStatus(ctx, s.DB, data.OrderID, data.TransactionStatus)
		if err != nil {
			return err
		}

		return nil
	} else if data.TransactionStatus != "settlement" && data.TransactionStatus != "capture" {
		return nil
	}

	orderData, err := s.orderRepo.FindByID(ctx, s.DB, data.OrderID)
	if err != nil {
		return err
	}

	reservationKey, err := helpers.GenerateSecret(32)
	if err != nil {
		return err
	}

	reservationData := models.HotelReservations{
		ID:                uuid.NewString(),
		HotelOrdersID:     data.OrderID,
		UserID:            orderData.UserID,
		ReservationKey:    reservationKey,
		ReservationStatus: "Confirmed",
	}

	err = s.reservationRepo.Create(ctx, s.DB, reservationData)
	if err != nil {
		return err
	}

	err = s.orderRepo.UpdatePaymentStatus(ctx, s.DB, data.OrderID, "paid")
	if err != nil {
		return err
	}

	return nil
}
