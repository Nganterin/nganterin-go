package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type HotelInputDTO struct {
	PartnerID   string
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`

	HotelRooms      []HotelRoomInput       `json:"hotel_rooms"`
	HotelsLocation  HotelsLocationInput    `json:"hotels_location"`
	HotelPhotos     []HotelPhotoInput      `json:"hotel_photos"`
	HotelFacilities []HotelFacilitiesInput `json:"hotel_facilities"`
}

type HotelRoomInput struct {
	Type           string `json:"type" binding:"required"`
	MaxVisitor     int    `json:"max_visitor" binding:"required"`
	BedType        string `json:"bed_type" binding:"required"`
	RoomSize       int    `json:"room_size" binding:"required"`
	SmokingAllowed bool   `json:"smoking_allowed" binding:"required"`
	OvernightPrice int64  `json:"overnight_price" binding:"required"`
	TotalRoom      int    `json:"total_room" binding:"required"`
	TotalBooked    int    `json:"total_booked" binding:"required"`

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" mapstructure:"hotel_room_photos"`
}

type HotelsLocationInput struct {
	Country         string `json:"country" binding:"required"`
	State           string `json:"state" binding:"required"`
	City            string `json:"city" binding:"required"`
	ZipCode         string `json:"zip_code" binding:"required"`
	CompleteAddress string `json:"complete_address" binding:"required"`
	Gmaps           string `json:"gmaps" binding:"required"`
}

type HotelPhotoInput struct {
	URL string `json:"url" binding:"required"`
}

type HotelRoomPhotoInput struct {
	URL string `json:"url" binding:"required"`
}

type HotelFacilitiesInput struct {
	Facility string `json:"facility" binding:"required"`
}

type HotelOrderInput struct {
	UserID           string
	HotelID          string    `json:"hotel_id" binding:"required"`
	HotelRoomID      uint      `json:"room_id" binding:"required"`
	CheckInDate      time.Time `json:"check_in_date" binding:"required"`
	CheckOutDate     time.Time `json:"check_out_date" binding:"required"`
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