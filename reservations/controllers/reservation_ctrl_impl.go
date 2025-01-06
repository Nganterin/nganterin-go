package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
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

func (h *CompControllersImpl) FindByReservationKey(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.FindByReservationKey(ctx, reservationKey)
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

func (h *CompControllersImpl) CheckIn(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CheckIn(ctx, reservationKey)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "reservation checked in successfully",
	})
}

func (h *CompControllersImpl) CheckOut(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CheckOut(ctx, reservationKey)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "reservation checked out successfully",
	})
}

func (h *CompControllersImpl) FindLast12MonthReservationCount(ctx *gin.Context) {
	partnerData := helpers.GetPartnerData(ctx)

	result, err := h.services.FindLast12MonthReservationCount(ctx, partnerData.ID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Data:    result,
	})
}
