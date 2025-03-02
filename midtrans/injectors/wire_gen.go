// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"nganterin-go/api/orders/repositories"
	repositories2 "nganterin-go/api/reservations/repositories"
	"nganterin-go/midtrans/notifications/controllers"
	"nganterin-go/midtrans/notifications/services"
)

// Injectors from injector.go:

func InitializeMidtransController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	repositoriesCompRepositories := repositories2.NewComponentRepository()
	compServices := services.NewComponentServices(compRepositories, repositoriesCompRepositories, db)
	compControllers := controllers.NewCompController(compServices)
	return compControllers
}

// injector.go:

var midtransFeatureSet = wire.NewSet(repositories.NewComponentRepository, repositories2.NewComponentRepository, services.NewComponentServices, controllers.NewCompController)
