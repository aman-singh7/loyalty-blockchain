package adapter

import (
	"github.com/aman-singh7/loyalty-blockchain/application/api"
	userService "github.com/aman-singh7/loyalty-blockchain/application/core/user"
	userRepository "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
	userController "github.com/aman-singh7/loyalty-blockchain/infrastructure/rest/controllers/user"
)

func UserAdapter(api *api.Api) *userController.Controller {
	uRepo := userRepository.NewRepository()
	uService := userService.NewService(uRepo, api)
	return &userController.Controller{UserService: uService}
}
