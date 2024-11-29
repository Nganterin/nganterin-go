package handlers

import (
	"net/http"
	"nganterin-go/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterUserCredential(c *gin.Context) {
	var data dto.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "all fields can't be null"})
		return
	}

	err = h.service.RegisterUserCredential(data)
	if err != nil {
		if err.Error() == "409" {
			c.JSON(http.StatusConflict, dto.Response{Status: http.StatusConflict, Error: "email already exists, please login"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "successfully register user"})
}

func (h *compHandlers) LoginUserCredentials(c *gin.Context) {
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

	token, err := h.service.LoginUserCredentials(data.Email, data.Password)
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

func (h *compHandlers) VerifyUserEmail(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "token required"})
		return
	}

	err := h.service.VerifyUserEmail(token)
	if err != nil {
		if err.Error() == "404" {
			c.JSON(http.StatusNotFound, dto.Response{Status: http.StatusNotFound, Error: "token not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "user email successfully verified"})
}
