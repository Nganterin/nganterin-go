package services

import (
	"errors"
	"nganterin-go/dto"
	"nganterin-go/mapper"
	"os"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (s *compServices) RegisterHotelOrder(data dto.HotelOrderInput) (*snap.Response, error) {
	roomData, err := s.repo.GetHotelRoomByID(data.HotelRoomID)
	if err != nil {
		return nil, errors.New("400")
	}

	userData, err := s.repo.GetUserDetailsByID(data.UserID)
	if err != nil {
		return nil, errors.New("403")
	}

	input := mapper.MapHotelOrderInputToModel(data)

	duration := input.CheckOutDate.Sub(input.CheckInDate)
	days := int(duration.Hours() / 24)

	input.TotalPrice = roomData.OvernightPrice * int64(days)
	input.ID = uuid.NewString()

	var m snap.Client
	m.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  input.ID,
			GrossAmt: input.TotalPrice,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: userData.Name,
			Email: userData.Email,
			Phone: userData.PhoneNumber,
		},
	}

	snapResp, _ := m.CreateTransaction(req)
	input.SnapToken = snapResp.Token

	err = s.repo.RegisterHotelOrder(input)
	if err != nil {
		return nil, err
	}

	return snapResp, nil
}