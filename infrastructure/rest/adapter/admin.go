package adapter

import (
	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	adminService "github.com/aman-singh7/loyalty-blockchain/application/core/admin"
	adminRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/admin"
	adminController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/admin"
)

func AdminAdapter(api *apiService.Service) *adminController.Controller {
	aRepo := adminRepository.NewRepository()
	uService := adminService.NewService(aRepo, api)
	return &adminController.Controller{AdminService: uService}
}
