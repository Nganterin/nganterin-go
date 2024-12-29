package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/hotels/services"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompService
}

func NewCompController(compServices services.CompService) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var hotelInput dto.HotelInputDTO

	if jsonErr := ctx.ShouldBindJSON(&hotelInput); jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	partnerData := helpers.GetPartnerData(ctx)

	hotelInput.PartnerID = partnerData.ID

	hotelID, err := h.services.Create(ctx, hotelInput)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "data retrieved successfully",
		Data:    hotelID,
	})
}

func (h *CompControllersImpl) FindAll(ctx *gin.Context) {
	result, err := h.services.FindAll(ctx)
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

func (h *CompControllersImpl) FindByKeyword(ctx *gin.Context) {
	keyword := ctx.Query("q")

	result, err := h.services.FindByKeyword(ctx, keyword)
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

func (h *CompControllersImpl) FindByID(ctx *gin.Context) {
	hotelID := ctx.Query("id")

	if hotelID == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.FindByID(ctx, hotelID)
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
