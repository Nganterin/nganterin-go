package handlers

import (
	"bytes"
	"crypto/hmac"
	"io/ioutil"
	"log"
	"net/http"
	"nganterin-go/dto"
	"nganterin-go/helpers"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) MidtransNotification(c *gin.Context) {
    bodyBytes, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        log.Println("Error reading request body:", err)
        c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "Invalid request"})
        return
    }

    c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

    var data dto.MidtransNotification
    if err := c.ShouldBindJSON(&data); err != nil {
        log.Println("JSON Binding Error:", err)
        c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "Invalid JSON"})
        return
    }

    if data.SignatureKey == "" {
        log.Println("Missing signature in body")
        c.JSON(http.StatusUnauthorized, dto.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
        return
    }

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

    dataString := data.OrderID + data.StatusCode + data.GrossAmount + serverKey

    calculatedSignature := helpers.EncryptToSHA512(dataString)

	log.Println("CALCULATED SIGNATURE: ", calculatedSignature)
	log.Println("RECIEVED SIGNATURE: ", data.SignatureKey)

    if !hmac.Equal([]byte(calculatedSignature), []byte(data.SignatureKey)) {
        log.Println("Invalid signature")
        c.JSON(http.StatusUnauthorized, dto.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
        return
    }

    if err := h.service.MidtransNotification(data); err != nil {
        log.Println("Service Error:", err)
        c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Success"})
}