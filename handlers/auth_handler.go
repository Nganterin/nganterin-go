package handlers

import (
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) AuthTest(c *gin.Context) {
	user_data := helpers.GetUserData(c)

	c.JSON(http.StatusAccepted, dto.Response{Status: http.StatusAccepted, Message: "Test Auth Success", Data: user_data})
}