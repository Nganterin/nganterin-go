package helpers

import (
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

func GetPartnerData(c *gin.Context) dto.Partner {
	partner_data, _ := c.Get("partner")

	partnerDTO := partner_data.(dto.Partner)

	return partnerDTO
}