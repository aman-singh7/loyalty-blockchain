package adapter

import (
	"database/sql"

	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	authService "github.com/aman-singh7/loyalty-blockchain/application/core/auth"
	userRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	authController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/auth"
)

func AuthAdapter(api *apiService.Service, db *sql.DB) *authController.Controller {
	uRepo := userRepository.NewRepository(db)
	aService := authService.NewService(uRepo)
	return &authController.Controller{AuthService: aService}
}
