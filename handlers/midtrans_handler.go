package handlers

import (
	"crypto/hmac"
	"log"
	"net/http"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) MidtransNotification(c *gin.Context) {
    var data dto.MidtransNotification
    if err := c.ShouldBindJSON(&data); err != nil {
        log.Println("JSON Binding Error:", err)
        c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "Invalid JSON"})
        return
    }

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
    dataString := data.OrderID + data.StatusCode + data.GrossAmount + serverKey
    calculatedSignature := helpers.EncryptToSHA512(dataString)

    if !hmac.Equal([]byte(calculatedSignature), []byte(data.SignatureKey)) {
        c.JSON(http.StatusUnauthorized, dto.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
        return
    }

    if err := h.service.MidtransNotification(data); err != nil {
        c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Success"})
}