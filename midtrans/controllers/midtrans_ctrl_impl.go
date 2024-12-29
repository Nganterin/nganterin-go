package controllers

import (
	"crypto/hmac"
	"net/http"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/midtrans/services"
	"nganterin-go/models/dto"
	"os"

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


func (h *CompControllersImpl) Notification(ctx *gin.Context) {
    var data dto.MidtransNotification
    if jsonErr := ctx.ShouldBindJSON(&data); jsonErr != nil {
        ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
    }

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
    dataString := data.OrderID + data.StatusCode + data.GrossAmount + serverKey
    calculatedSignature := helpers.EncryptToSHA512(dataString)

    if !hmac.Equal([]byte(calculatedSignature), []byte(data.SignatureKey)) {
        ctx.JSON(http.StatusUnauthorized, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrUnauthorized))
        return
    }

    if err := h.services.Notification(ctx, data); err != nil {
        ctx.JSON(err.Status, err)
        return
    }

    ctx.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "success"})
}