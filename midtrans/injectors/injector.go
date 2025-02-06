// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	orderRepositories "nganterin-go/api/orders/repositories"
	reservationRepositories "nganterin-go/api/reservations/repositories"
	midtransControllers "nganterin-go/midtrans/notifications/controllers"
	midtransServices "nganterin-go/midtrans/notifications/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var midtransFeatureSet = wire.NewSet(
	orderRepositories.NewComponentRepository,
	reservationRepositories.NewComponentRepository,
	midtransServices.NewComponentServices,
	midtransControllers.NewCompController,
)

func InitializeMidtransController(db *gorm.DB, validate *validator.Validate) midtransControllers.CompControllers {
	wire.Build(midtransFeatureSet)
	return nil
}
