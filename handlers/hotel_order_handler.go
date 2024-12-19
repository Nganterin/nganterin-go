package handlers

import (
	"net/http"
	"nganterin-go/dto"
	"nganterin-go/helpers"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterHotelOrder(c *gin.Context) {
	var orderInput dto.HotelOrderInput

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	if orderInput.IsForSomeoneElse {
		if orderInput.SomeoneName == "" || orderInput.SomeoneRegion == "" {
			c.JSON(http.StatusBadRequest, dto.Response{
				Status: http.StatusBadGateway,
				Error:  "someone_name and someone_region required for is_for_someone_else true",
			})
			return
		}
	}

	userData := helpers.GetUserData(c)

	orderInput.UserID = userData.ID

	result, err := h.service.RegisterHotelOrder(orderInput)
	if err != nil {
		if err.Error() == "400" {
			c.JSON(http.StatusBadRequest, dto.Response{
				Status: http.StatusBadRequest,
				Error:  "bad request",
			})
		} else if err.Error() == "401" {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Status: http.StatusUnauthorized,
				Error:  "user not found, who are you?",
			})
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "order successfully registered",
	})
}
