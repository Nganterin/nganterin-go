package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/dto"
	"nganterin-go/reviews/services"

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
	var data dto.HotelReviewInput

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	data.UserID = userData.ID

	err = h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Review created successfully",
	})
}
