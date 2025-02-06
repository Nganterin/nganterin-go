// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	userControllers "nganterin-go/api/users/controllers"
	userRepositories "nganterin-go/api/users/repositories"
	userServices "nganterin-go/api/users/services"

	hotelControllers "nganterin-go/api/hotels/controllers"
	hotelRepositories "nganterin-go/api/hotels/repositories"
	hotelServices "nganterin-go/api/hotels/services"

	orderControllers "nganterin-go/api/orders/controllers"
	orderRepositories "nganterin-go/api/orders/repositories"
	orderServices "nganterin-go/api/orders/services"

	emailServices "nganterin-go/emails/services"

	storageControllers "nganterin-go/api/storages/controllers"
	storageRepositories "nganterin-go/api/storages/repositories"
	storageServices "nganterin-go/api/storages/services"

	partnerControllers "nganterin-go/api/partners/controllers"
	partnerRepositories "nganterin-go/api/partners/repositories"
	partnerServices "nganterin-go/api/partners/services"

	reservationControllers "nganterin-go/api/reservations/controllers"
	reservationRepositories "nganterin-go/api/reservations/repositories"
	reservationServices "nganterin-go/api/reservations/services"

	reviewControllers "nganterin-go/api/reviews/controllers"
	reviewRepositories "nganterin-go/api/reviews/repositories"
	reviewServices "nganterin-go/api/reviews/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
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
	hotelRepositories.NewComponentRepository,
	reservationServices.NewComponentServices,
	reservationControllers.NewCompController,
)

var reviewFeatureSet = wire.NewSet(
	orderRepositories.NewComponentRepository,
	reservationRepositories.NewComponentRepository,
	reviewRepositories.NewComponentRepository,
	reviewServices.NewComponentServices,
	reviewControllers.NewCompController,
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

func InitializeReviewController(db *gorm.DB, validate *validator.Validate) reviewControllers.CompControllers {
	wire.Build(reviewFeatureSet)
	return nil
}
