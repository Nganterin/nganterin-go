package handlers

import (
	"net/http"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterPartner(c *gin.Context) {
	var data dto.Partner

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "all fields can't be null"})
		return
	}

	err = helpers.ValidateURL(data.LegalityFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "invalid url for legality_file"})
		return
	}

	err = helpers.ValidateURL(data.MOUFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "invalid url for mou_file"})
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

func (h *compHandlers) LoginPartner(c *gin.Context) {
	type Credentials struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var data Credentials

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}

	token, err := h.service.LoginPartner(data.Email, data.Password)
	if err != nil {
		if err.Error() == "403" {
			c.JSON(http.StatusForbidden, dto.Response{Status: http.StatusForbidden, Error: "email not verified, check your email include spam folder"})
			return
		} else if err.Error() == "401" {
			c.JSON(http.StatusUnauthorized, dto.Response{Status: http.StatusUnauthorized, Error: "invalid email or password"})
			return
		} else if err.Error() == "404" {
			c.JSON(http.StatusNotFound, dto.Response{Status: http.StatusNotFound, Error: "email not found, please register"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "login successfully", Data: token})
}

func (h *compHandlers) VerifyPartnerEmail(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "token required"})
		return
	}

	err := h.service.VerifyPartnerEmail(token)
	if err != nil {
		if err.Error() == "404" {
			c.JSON(http.StatusNotFound, dto.Response{Status: http.StatusNotFound, Error: "token not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "partner email successfully verified"})
}