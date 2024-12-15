package hotels

import (
	"net/http"

	"nganterin-go/dto"
	"nganterin-go/services/hotels"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	service hotels.HotelService
}

func NewHotelHandler(s hotels.HotelService) *HotelHandler {
	return &HotelHandler{
		service: s,
	}
}

func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var hotelInput dto.HotelInputDTO

	if err := c.ShouldBindJSON(&hotelInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	hotelID, err := h.service.CreateHotel(hotelInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create hotel",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Hotel created successfully",
		Data:    hotelID,
	})
}
