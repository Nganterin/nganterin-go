package controllers

import (
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/hotels/services"
	"nganterin-go/models/dto"
	"strconv"

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

func (h *CompControllersImpl) SearchEngine(ctx *gin.Context) {
	priceStart, inputErr := strconv.ParseInt(ctx.Query("priceStart"), 10, 64)
	if inputErr != nil {
		priceStart = 0
	}

	priceEnd, inputErr := strconv.ParseInt(ctx.Query("priceEnd"), 10, 64)
	if inputErr != nil {
		priceEnd = 64000000
	}

	minStars, inputErr := strconv.Atoi(ctx.Query("minStars"))
	if inputErr != nil {
		minStars = 0
	}

	minVisitor, inputErr := strconv.Atoi(ctx.Query("minVisitor"))
	if inputErr != nil {
		minVisitor = 0
	}

	searchInput := dto.HotelSearch{
		Keyword:        ctx.Query("keyword"),
		Name:           ctx.Query("name"),
		City:           ctx.Query("city"),
		Country:        ctx.Query("country"),
		PriceStart:     priceStart,
		PriceEnd:       priceEnd,
		MinimumStars:   minStars,
		MinimumVisitor: minVisitor,
	}

	result, err := h.services.SearchEngine(ctx, searchInput)
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

func (h *CompControllersImpl) FindByPartnerID(ctx *gin.Context) {
	partnerData := helpers.GetPartnerData(ctx)

	result, err := h.services.FindByPartnerID(ctx, partnerData.ID)
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