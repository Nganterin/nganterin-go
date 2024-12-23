package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"nganterin-go/dto"
	"nganterin-go/mapper"
	"os"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (s *compServices) RegisterHotelOrder(data dto.HotelOrderInput) (*dto.HotelOrderOutput, error) {
	roomData, err := s.repo.GetHotelRoomByID(data.HotelRoomID)
	if err != nil {
		return nil, errors.New("400")
	}

	userData, err := s.repo.GetUserDetailsByID(data.UserID)
	if err != nil {
		return nil, errors.New("401")
	}

	input := mapper.MapHotelOrderInputToModel(data)

	duration := input.CheckOutDate.Sub(input.CheckInDate)
	days := int(duration.Hours()/24)

	if days <= 0 {
		return nil, errors.New("400")
	}

	input.TotalPrice = roomData.OvernightPrice * int64(days)
	input.ID = uuid.NewString()

	var m snap.Client
	m.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	secdat := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{\n  id: %s\n}", input.ID)))

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
		Callbacks: &snap.Callbacks{
			Finish: os.Getenv("FRONT_END_BASE_URL") + "/payment/hotel?secdat=" + secdat,
		},
	}

	snapResp, _ := m.CreateTransaction(req)
	input.SnapToken = snapResp.Token

	err = s.repo.RegisterHotelOrder(input)
	if err != nil {
		return nil, err
	}

	output := dto.HotelOrderOutput{
		ID:          input.ID,
		Token:       snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
	}

	return &output, nil
}

func (s *compServices) GetHotelOrderByID(id string) (*dto.HotelOrderDetailsOutput, error) {
	data, err := s.repo.GetHotelOrderByID(id)
	if err != nil {
		return nil, err
	}

	result := mapper.MapHotelOrderModelToOutput(*data)

	return &result, nil
}

func (s *compServices) GetAllHotelOrderByUserID(id string) ([]dto.HotelOrderDetailsOutput, error) {
	data, err := s.repo.GetAllHotelOrderByUserID(id)
	if err != nil {
		return nil, err
	}

	var result []dto.HotelOrderDetailsOutput
	for _, item := range data {
		result = append(result, mapper.MapHotelOrderModelToOutput(item))
	}

	return result, nil
}
