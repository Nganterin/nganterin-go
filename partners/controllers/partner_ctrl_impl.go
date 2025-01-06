package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/dto"
	"nganterin-go/partners/services"

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

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.Partner

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "partner successfully registered"})
}

func (h *CompControllersImpl) Login(ctx *gin.Context) {
	type Credentials struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var data Credentials

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	token, err := h.services.Login(ctx, data.Email, data.Password)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "login successfully", Data: token})
}

func (h *CompControllersImpl) VerifyEmail(ctx *gin.Context) {
	token := ctx.Query("token")

	if token == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.VerifyEmail(ctx, token)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "partner email successfully verified"})
}

func (h *CompControllersImpl) ApprovalCheck(ctx *gin.Context) {
	userData := helpers.GetUserData(ctx)

	token, err := h.services.ApprovalCheck(ctx, userData.ID)
	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "data already verified", Data: token})
}