package services

import (
	"log"
	"nganterin-go/dto"
)

func (s *compServices) MidtransNotification(data dto.MidtransNotification) error {
	log.Println("ORDER ID: ", data.OrderID)

	return nil
}
