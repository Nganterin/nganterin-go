package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	hotelRepo "nganterin-go/hotels/repositories"
	"nganterin-go/mapper"
	"nganterin-go/models/dto"
	"nganterin-go/orders/repositories"
	userRepo "nganterin-go/users/repositories"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo      repositories.CompRepositories
	hotelRepo hotelRepo.CompRepositories
	userRepo  userRepo.CompRepositories
	DB        *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, compHotelRepositories hotelRepo.CompRepositories, compUserRepositories userRepo.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		repo:      compRepositories,
		hotelRepo: compHotelRepositories,
		userRepo:  compUserRepositories,
		DB:        db,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.HotelOrderInput) (*dto.HotelOrderOutput, *exceptions.Exception) {
	roomData, err := s.hotelRepo.FindRoomByID(ctx, s.DB, data.HotelRoomID)
	if err != nil {
		return nil, err
	}

	userData, err := s.userRepo.FindByID(ctx, s.DB, data.UserID)
	if err != nil {
		return nil, err
	}

	input := mapper.MapHotelOrderInputToModel(data)

	days := helpers.GetDaysFromCheckInCheckOut(data.CheckInDate, data.CheckOutDate)

	if days <= 0 {
		return nil, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest)
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
			Finish: os.Getenv("WEBCLIENT_BASE_URL") + "/payment/hotel?secdat=" + secdat,
		},
	}

	snapResp, _ := m.CreateTransaction(req)
	input.SnapToken = snapResp.Token

	err = s.repo.Create(ctx, s.DB, input)
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

func (s *CompServicesImpl) FindByID(ctx *gin.Context, id string) (*dto.HotelOrderDetailsOutput, *exceptions.Exception) {
	data, err := s.repo.FindByID(ctx, s.DB, id)
	if err != nil {
		return nil, err
	}

	result := mapper.MapHotelOrderModelToOutput(*data)
	result.TotalDays = helpers.GetDaysFromCheckInCheckOut(data.CheckInDate, data.CheckOutDate)

	return &result, nil
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
