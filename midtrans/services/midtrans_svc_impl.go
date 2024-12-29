package services

import (
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"
	orderRepo "nganterin-go/orders/repositories"
	reservationRepo "nganterin-go/reservations/repositories"

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

	reservationData := database.HotelReservations{
		ID:                uuid.NewString(),
		HotelOrdersID:     data.OrderID,
		UserID:            orderData.UserID,
		ReservationKey:    reservationKey,
		ReservationStatus: "confirmed",
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
