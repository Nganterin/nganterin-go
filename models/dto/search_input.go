package dto

type HotelSearch struct {
	Keyword        string
	Name           string
	PriceStart     int64
	PriceEnd       int64
	City           string
	Country        string
	MinimumStars   int
	MinimumVisitor int
}
