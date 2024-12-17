package dto

type HotelInputDTO struct {
	PartnerID   string `json:"partner_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`

	HotelRooms      []HotelRoomInput       `json:"hotel_rooms"`
	HotelsLocation  HotelsLocationInput    `json:"hotels_location"`
	HotelPhotos     []HotelPhotoInput      `json:"hotel_photos"`
	HotelFacilities []HotelFacilitiesInput `json:"hotel_facilities"`
}

type HotelRoomInput struct {
	Type           string `json:"type" validate:"required"`
	MaxVisitor     int    `json:"max_visitor" validate:"required"`
	BedType        string `json:"bed_type" validate:"required"`
	RoomSize       int    `json:"room_size" validate:"required"`
	SmokingAllowed bool   `json:"smoking_allowed" validate:"required"`
	OvernightPrice int64  `json:"overnight_price" validate:"required"`
	TotalRoom      int    `json:"total_room" validate:"required"`
	TotalBooked    int    `json:"total_booked" validate:"required"`

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" mapstructure:"hotel_room_photos"`
}

type HotelsLocationInput struct {
	Country         string `json:"country" validate:"required"`
	State           string `json:"state" validate:"required"`
	City            string `json:"city" validate:"required"`
	ZipCode         string `json:"zip_code" validate:"required"`
	CompleteAddress string `json:"complete_address" validate:"required"`
	Gmaps           string `json:"gmaps" validate:"required"`
}

type HotelPhotoInput struct {
	URL string `json:"url" validate:"required"`
}

type HotelRoomPhotoInput struct {
	URL string `json:"url" validate:"required"`
}

type HotelFacilitiesInput struct {
	Facility string `json:"facility" validate:"required"`
}
