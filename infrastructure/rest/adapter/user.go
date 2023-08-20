package adapter

import (
	"database/sql"

	apiService "github.com/aman-singh7/loyalty-blockchain/application/api"
	userService "github.com/aman-singh7/loyalty-blockchain/application/core/user"
	userRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	userController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/user"
)

func UserAdapter(api *apiService.Service, db *sql.DB) *userController.Controller {
	uRepo := userRepository.NewRepository(db)
	uService := userService.NewService(uRepo, api)
	return &userController.Controller{UserService: uService}
}
