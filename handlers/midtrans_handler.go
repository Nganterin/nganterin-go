package handlers

import (
	"log"
	"net/http"
	"nganterin-go/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) MidtransNotification(c *gin.Context) {
	var data dto.MidtransNotification

	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
	}

	err = h.service.MidtransNotification(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError})
		log.Println(err.Error())
		return
	}
	
	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK})
}
