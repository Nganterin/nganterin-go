// go:build wireinject
// go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "nganterin-go/users/controllers"
	userRepositories "nganterin-go/users/repositories"
	userServices "nganterin-go/users/services"

	hotelControllers "nganterin-go/hotels/controllers"
	hotelRepositories "nganterin-go/hotels/repositories"
	hotelServices "nganterin-go/hotels/services"

	orderControllers "nganterin-go/orders/controllers"
	orderRepositories "nganterin-go/orders/repositories"
	orderServices "nganterin-go/orders/services"

	midtransControllers "nganterin-go/midtrans/controllers"
	midtransServices "nganterin-go/midtrans/services"

	emailServices "nganterin-go/emails/services"

	storageControllers "nganterin-go/storages/controllers"
	storageRepositories "nganterin-go/storages/repositories"
	storageServices "nganterin-go/storages/services"
	
	partnerControllers "nganterin-go/partners/controllers"
	partnerRepositories "nganterin-go/partners/repositories"
	partnerServices "nganterin-go/partners/services"
	
	reservationControllers "nganterin-go/reservations/controllers"
	reservationRepositories "nganterin-go/reservations/repositories"
	reservationServices "nganterin-go/reservations/services"

	"github.com/google/wire"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	emailServices.NewComponentServices,
	userServices.NewComponentServices,
	userControllers.NewCompController,
)

var hotelFeatureSet = wire.NewSet(
	hotelRepositories.NewComponentRepository,
	hotelServices.NewComponentServices,
	hotelControllers.NewCompController,
)

var orderFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
	hotelRepositories.NewComponentRepository,
	orderRepositories.NewComponentRepository,
	orderServices.NewComponentServices,
	orderControllers.NewCompController,
)

var midtransFeatureSet = wire.NewSet(
	orderRepositories.NewComponentRepository,
	reservationRepositories.NewComponentRepository,
	midtransServices.NewComponentServices,
	midtransControllers.NewCompController,
)

var storageFeatureSet = wire.NewSet(
	storageRepositories.NewComponentRepository,
	storageServices.NewComponentServices,
	storageControllers.NewCompController,
)

var partnerFeatureSet = wire.NewSet(
	partnerRepositories.NewComponentRepository,
	emailServices.NewComponentServices,
	partnerServices.NewComponentServices,
	partnerControllers.NewCompController,
)

var reservationFeatureSet = wire.NewSet(
	reservationRepositories.NewComponentRepository,
	reservationServices.NewComponentServices,
	reservationControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB, validate *validator.Validate) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeHotelController(db *gorm.DB, validate *validator.Validate) hotelControllers.CompControllers {
	wire.Build(hotelFeatureSet)
	return nil
}

func InitializeOrderController(db *gorm.DB, validate *validator.Validate) orderControllers.CompControllers {
	wire.Build(orderFeatureSet)
	return nil
}

func InitializeMidtransController(db *gorm.DB, validate *validator.Validate) midtransControllers.CompControllers {
	wire.Build(midtransFeatureSet)
	return nil
}

func InitializeStorageController(db *gorm.DB, validate *validator.Validate) storageControllers.CompControllers {
	wire.Build(storageFeatureSet)
	return nil
}

func InitializePartnerController(db *gorm.DB, validate *validator.Validate) partnerControllers.CompControllers {
	wire.Build(partnerFeatureSet)
	return nil
}

func InitializeReservationController(db *gorm.DB, validate *validator.Validate) reservationControllers.CompControllers {
	wire.Build(reservationFeatureSet)
	return nil
}
