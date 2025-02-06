package helpers

import (
	"net/http"
	"nganterin-go/partners/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

func GetPartnerData(c *gin.Context) (dto.Partner, *exceptions.Exception) {
	var result dto.Partner
	partner_data, _ := c.Get("partner")

	result, ok := partner_data.(dto.Partner)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}

	return result, nil
}
