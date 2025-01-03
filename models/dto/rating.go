package dto

type HotelAverageRating struct {
	Rating         float32 `json:"rating"`
	Cleanliness    float32 `json:"cleanliness"`
	Comfort        float32 `json:"comfort"`
	ServiceQuality float32 `json:"service_quality"`
	Facilities     float32 `json:"facilities"`
	ValueForMoney  float32 `json:"value_for_money"`
}

type HotelReviewStatistic struct {
	TotalReviews  int     `json:"total_reviews"`
	AverageRating float32 `json:"average_rating"`
	Percentage5   int     `json:"percentage_5"`
	Percentage4   int     `json:"percentage_4"`
	Percentage3   int     `json:"percentage_3"`
	Percentage2   int     `json:"percentage_2"`
	Percentage1   int     `json:"percentage_1"`
}
