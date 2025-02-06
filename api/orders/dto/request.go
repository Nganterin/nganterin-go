package dto

import (
	"encoding/json"
	"fmt"
	"time"
)


type HotelOrderInput struct {
	UserID           string
	HotelID          string    `json:"hotel_id" validate:"required"`
	HotelRoomID      uint      `json:"room_id" validate:"required"`
	CheckInDate      time.Time `json:"check_in_date" validate:"required"`
	CheckOutDate     time.Time `json:"check_out_date" validate:"required"`
	SpecialRequest   string    `json:"special_request"`
	IsForSomeoneElse bool      `json:"is_for_someone_else"`
	SomeoneName      string    `json:"someone_name"`
	SomeoneRegion    string    `json:"someone_region"`
}

func (h *HotelOrderInput) UnmarshalJSON(data []byte) error {
	type Alias HotelOrderInput
	aux := &struct {
		CheckInDate  string `json:"check_in_date" binding:"required"`
		CheckOutDate string `json:"check_out_date" binding:"required"`
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	layout := "2006-01-02"
	if aux.CheckInDate != "" {
		parsedCheckInDate, err := time.Parse(layout, aux.CheckInDate)
		if err != nil {
			return fmt.Errorf("invalid format for check_in_date: %v", err)
		}
		h.CheckInDate = parsedCheckInDate
	}

	if aux.CheckOutDate != "" {
		parsedCheckOutDate, err := time.Parse(layout, aux.CheckOutDate)
		if err != nil {
			return fmt.Errorf("invalid format for check_out_date: %v", err)
		}
		h.CheckOutDate = parsedCheckOutDate
	}

	return nil
}