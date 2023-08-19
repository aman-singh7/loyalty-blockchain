package adapter

import (
	"github.com/aman-singh7/loyalty-blockchain/application/api"
	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	brandService "github.com/aman-singh7/loyalty-blockchain/application/core/brand"
	brandRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/brand"
	brandController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/brand"
)

func BrandAdapter(api *api.Api) *brandController.Controller {
	bRepo := brandRepository.NewRepository()
	aService := apiService.NewService(api)
	bService := brandService.NewService(bRepo, aService)
	return &brandController.Controller{BrandService: bService}
}
