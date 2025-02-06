// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	partnerControllers "nganterin-go/partners/controllers"
	partnerRepositories "nganterin-go/partners/repositories"
	partnerServices "nganterin-go/partners/services"

	emailServices "nganterin-go/emails/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var partnerFeatureSet = wire.NewSet(
	partnerRepositories.NewComponentRepository,
	emailServices.NewComponentServices,
	partnerServices.NewComponentServices,
	partnerControllers.NewCompController,
)

func InitializePartnerController(db *gorm.DB, validate *validator.Validate) partnerControllers.CompControllers {
	wire.Build(partnerFeatureSet)
	return nil
}