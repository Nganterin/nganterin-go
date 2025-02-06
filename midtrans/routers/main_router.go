package routers

import (
	"nganterin-go/midtrans/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func MidtransRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	midtransController := injectors.InitializeMidtransController(db, validate)

	MidtransRoutes(r, midtransController)
}