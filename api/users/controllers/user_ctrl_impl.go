package controllers

import (
	"net/http"
	"nganterin-go/api/users/services"
	"nganterin-go/api/users/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) CreateCredentials(ctx *gin.Context) {
	var data dto.User

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CreateCredentials(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "successfully register user",
	})
}

func (h *CompControllersImpl) CreateGoogleOAuth(ctx *gin.Context) {
	var data dto.UserGoogle

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	token, err := h.services.CreateGoogleOAuth(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    token,
		Message: "successfully register user",
	})
}

func (h *CompControllersImpl) LoginCredentials(ctx *gin.Context) {
	type Credentials struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var data Credentials

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	token, err := h.services.LoginCredentials(ctx, data.Email, data.Password)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status: http.StatusOK,
		Data:   token, Message: "successfully login user",
	})
}

func (h *CompControllersImpl) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.VerifyEmail(ctx, token)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "user email successfully verified",
	})
}

func (h *CompControllersImpl) LoginGoogleOAuth(ctx *gin.Context) {
	type Credentials struct {
		Email     string `json:"email" binding:"required"`
		GoogleSUB string `json:"google_sub" binding:"required"`
	}

	var data Credentials

	jsonErr := ctx.ShouldBindBodyWithJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "invalid request"})
		return
	}

	token, err := h.services.LoginGoogleOAuth(ctx, data.Email, data.GoogleSUB)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "login successfully",
		Data:    token,
	})
}

func (h *CompControllersImpl) AuthTest(ctx *gin.Context) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusAccepted, dto.Response{
		Status:  http.StatusAccepted,
		Message: "Test Auth Success",
		Data:    userData,
	})
}
