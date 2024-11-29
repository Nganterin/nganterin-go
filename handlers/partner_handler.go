package handlers

import (
	"net/http"
	"nganterin-go/dto"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) PartnerRegister(c *gin.Context) {
	var data dto.Partner

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "all fields can't be null"})
		return
	}

	err = h.service.RegisterPartner(data)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			var error_message string

			if strings.Contains(err.Error(), "company_name") {
				error_message = "company name already exists"
			} else if strings.Contains(err.Error(), "company_email") {
				error_message = "company email already exists"
			} else {
				error_message = "email already exists, please login"
			}

			c.JSON(http.StatusConflict, dto.Response{Status: http.StatusConflict, Error: error_message})
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "partner successfully registered"})
}
