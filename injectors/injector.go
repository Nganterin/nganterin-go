// go:build wireinject
//go:build wireinject
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
	
	reservationRepositories "nganterin-go/reservations/repositories"
	
	storageControllers "nganterin-go/storages/controllers"
	storageRepositories "nganterin-go/storages/repositories"
	storageServices "nganterin-go/storages/services"
	
	partnerControllers "nganterin-go/partners/controllers"
	partnerRepositories "nganterin-go/partners/repositories"
	partnerServices "nganterin-go/partners/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var userFeatureSet = wire.NewSet(
	userRepositories.NewComponentRepository,
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
	partnerServices.NewComponentServices,
	partnerControllers.NewCompController,
)

func InitializeUserController(db *gorm.DB) userControllers.CompControllers {
	wire.Build(userFeatureSet)
	return nil
}

func InitializeHotelController(db *gorm.DB) hotelControllers.CompControllers {
	wire.Build(hotelFeatureSet)
	return nil
}

func InitializeOrderController(db *gorm.DB) orderControllers.CompControllers {
	wire.Build(orderFeatureSet)
	return nil
}

func InitializeMidtransController(db *gorm.DB) midtransControllers.CompControllers {
	wire.Build(midtransFeatureSet)
	return nil
}

func InitializeStorageController(db *gorm.DB) storageControllers.CompControllers {
	wire.Build(storageFeatureSet)
	return nil
}

func InitializePartnerController(db *gorm.DB) partnerControllers.CompControllers {
	wire.Build(partnerFeatureSet)
	return nil
}
