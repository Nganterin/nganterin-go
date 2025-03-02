package services

import (
	"nganterin-go/api/hotels/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompService interface {
	Create(ctx *gin.Context, data dto.HotelInputDTO) (*string, *exceptions.Exception)
	FindAll(ctx *gin.Context) (*[]dto.HotelOutputDTO, *exceptions.Exception)
	FindByID(ctx *gin.Context, id string) (*dto.HotelOutputDTO, *exceptions.Exception)
	SearchEngine(ctx *gin.Context, searchInput dto.HotelSearch) (*[]dto.HotelOutputDTO, *exceptions.Exception)
	FindByPartnerID(ctx *gin.Context, partnerID string) ([]dto.HotelOutputDTO, *exceptions.Exception)
}
