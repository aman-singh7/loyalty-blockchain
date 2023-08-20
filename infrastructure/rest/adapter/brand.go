package adapter

import (
	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	brandService "github.com/aman-singh7/loyalty-blockchain/application/core/brand"
	brandRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/brand"
	brandController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/brand"
)

func BrandAdapter(api *apiService.Service) *brandController.Controller {
	bRepo := brandRepository.NewRepository()
	bService := brandService.NewService(bRepo, api)
	return &brandController.Controller{BrandService: bService}
}
