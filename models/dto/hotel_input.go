package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type HotelInputDTO struct {
	PartnerID   string
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`

	HotelRooms      []HotelRoomInput       `json:"hotel_rooms" validate:"required,dive"`
	HotelsLocation  HotelsLocationInput    `json:"hotels_location" validate:"required"`
	HotelPhotos     []HotelPhotoInput      `json:"hotel_photos" validate:"required,dive"`
	HotelFacilities []HotelFacilitiesInput `json:"hotel_facilities" validate:"required,dive"`
}

type HotelRoomInput struct {
	Type           string `json:"type" validate:"required"`
	MaxVisitor     int    `json:"max_visitor" validate:"required,number"`
	BedType        string `json:"bed_type" validate:"required"`
	RoomSize       int    `json:"room_size" validate:"required,number"`
	SmokingAllowed bool   `json:"smoking_allowed" validate:"boolean"`
	OvernightPrice int64  `json:"overnight_price" validate:"required,number,min=0"`
	TotalRoom      int    `json:"total_room" validate:"required,number,min=0"`
	TotalBooked    int    `json:"total_booked" validate:"number,min=0"`

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" validate:"required,dive" mapstructure:"hotel_room_photos"`
}

type HotelsLocationInput struct {
	Country         string `json:"country" validate:"required"`
	State           string `json:"state" validate:"required"`
	City            string `json:"city" validate:"required"`
	ZipCode         string `json:"zip_code" validate:"required"`
	CompleteAddress string `json:"complete_address" validate:"required"`
	Gmaps           string `json:"gmaps" validate:"required,url"`
}

type HotelPhotoInput struct {
	URL string `json:"url" validate:"required,url"`
}

type HotelRoomPhotoInput struct {
	URL string `json:"url" validate:"required,url"`
}

type HotelFacilitiesInput struct {
	Facility string `json:"facility" validate:"required"`
}

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

type HotelReviewInput struct {
	HotelOrdersID  string `json:"hotel_order_id" validate:"required"`
	UserID         string
	Review         string `json:"review" validate:"required"`
	Cleanliness    int    `json:"cleanliness" validate:"required,min=1,max=5"`
	Comfort        int    `json:"comfort" validate:"required,min=1,max=5"`
	ServiceQuality int    `json:"service_quality" validate:"required,min=1,max=5"`
	Facilities     int    `json:"facilities" validate:"required,min=1,max=5"`
	ValueForMoney  int    `json:"value_for_money" validate:"required,min=1,max=5"`
	Rating         int    `json:"rating" validate:"required,min=1,max=5"`
}
