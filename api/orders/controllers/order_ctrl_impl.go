package controllers

import (
	"net/http"
	"nganterin-go/api/orders/services"
	"nganterin-go/models/dto"
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

	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

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
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

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

func (h *CompControllersImpl) YearlyOrderAnalytic(ctx *gin.Context) {
	partnerData, err := helpers.GetPartnerData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	result, err := h.services.YearlyOrderAnalytic(ctx, partnerData.ID)
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
