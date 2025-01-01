package controllers

import (
	"net/http"
	"nganterin-go/helpers"
	"nganterin-go/models/dto"
	"nganterin-go/reservations/services"

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

func (h *CompControllersImpl) FindByUserID(ctx *gin.Context) {
	userData := helpers.GetUserData(ctx)

	result, err := h.services.FindByUserID(ctx, userData.ID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "data retrieved successfully",
	})
}
