package dto

type HotelAverageRating struct {
	Rating         float32 `json:"rating"`
	Cleanliness    float32 `json:"cleanliness"`
	Comfort        float32 `json:"comfort"`
	ServiceQuality float32 `json:"service_quality"`
	Facilities     float32 `json:"facilities"`
	ValueForMoney  float32 `json:"value_for_money"`
}
