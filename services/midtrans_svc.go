package services

import (
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"nganterin-go/models"

	"github.com/google/uuid"
)

func (s *compServices) MidtransNotification(data dto.MidtransNotification) error {
	if data.TransactionStatus != "settlement" && data.TransactionStatus != "capture" {
		return nil
	}

	orderData, err := s.repo.GetHotelOrderByID(data.OrderID)
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
		ReservationStatus: "confirmed",
	}

	err = s.repo.RegisterHotelReservation(reservationData)
	if err != nil {
		return err
	}

	err = s.repo.UpdateHotelOrderPaymentStatus(data.OrderID, "paid")
	if err != nil {
		return err
	}

	return nil
}
