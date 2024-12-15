package handlers

import (
	"net/http"

	"nganterin-go/dto"
	"nganterin-go/helpers"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterHotel(c *gin.Context) {
	var hotelInput dto.HotelInputDTO

	if err := c.ShouldBindJSON(&hotelInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	partnerData := helpers.GetPartnerData(c)

	hotelInput.PartnerID = partnerData.ID

	hotelID, err := h.service.RegisterHotel(hotelInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "failed to create hotel",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "data retrieved successfully",
		Data:    hotelID,
	})
}

func (h *compHandlers) GetAllHotels(c *gin.Context) {
	result, err := h.service.GetAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Data:    result,
	})
}

func (h *compHandlers) SearchHotels(c *gin.Context) {
	keyword := c.Query("q")

	result, err := h.service.SearchHotels(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Data:    result,
	})
}

func (h *compHandlers) GetHotelByID(c *gin.Context) {
	hotelID := c.Query("id")

	if hotelID == "" {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status: http.StatusBadRequest,
			Error:  "id is required",
		})
		return
	}

	result, err := h.service.GetHotelByID(hotelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Data:    result,
	})
}
