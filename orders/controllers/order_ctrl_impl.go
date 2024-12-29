package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/dto"
	"nganterin-go/orders/services"

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
	var orderInput dto.HotelOrderInput

	if jsonErr := ctx.ShouldBindJSON(&orderInput); jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	if orderInput.IsForSomeoneElse {
		if orderInput.SomeoneName == "" || orderInput.SomeoneRegion == "" {
			ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, "someone_name and someone_region required for is_for_someone_else true"))
			return
		}
	}

	userData := helpers.GetUserData(ctx)

	orderInput.UserID = userData.ID

	result, err := h.services.Create(ctx, orderInput)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "order successfully registered",
	})
}

func (h *CompControllersImpl) FindByID(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.FindByID(ctx, id)
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
