package adapter

import (
	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	authService "github.com/aman-singh7/loyalty-blockchain/application/core/auth"
	userRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	authController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/auth"
)

func AuthAdapter(api *apiService.Service) *authController.Controller {
	uRepo := userRepository.NewRepository()
	aService := authService.NewService(uRepo)
	return &authController.Controller{AuthService: aService}
}
