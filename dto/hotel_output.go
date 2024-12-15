package dto

type HotelOutputDTO struct {
	ID             string                 `json:"id"`
	PartnerID      string                 `json:"partner_id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	HotelDetails   []HotelDetailOutput    `json:"hotel_details"`
	HotelsLocation []HotelsLocationOutput `json:"hotels_location"`
	HotelPhotos    []HotelPhotoOutput     `json:"hotel_photos"`
}

type HotelDetailOutput struct {
	ID             uint   `json:"id"`
	Type           string `json:"type"`
	MaxVisitor     int    `json:"max_visitor"`
	BedType        string `json:"bed_type"`
	RoomSize       int    `json:"room_size"`
	SmokingAllowed bool   `json:"smoking_allowed"`
	OvernightPrice int64  `json:"overnight_price"`
	TotalRoom      int    `json:"total_room"`
	TotalBooked    int    `json:"total_booked"`
}

type HotelsLocationOutput struct {
	ID              uint   `json:"id"`
	Country         string `json:"country"`
	State           string `json:"state"`
	City            string `json:"city"`
	ZipCode         string `json:"zip_code"`
	CompleteAddress string `json:"complete_address"`
	Gmaps           string `json:"gmaps"`
}

type HotelPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}
