package user

import (
	"github.com/aman-singh7/loyalty-blockchain/application/api"
	userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"
)

type Service struct {
	repo *userRepo.Repository
	api  *api.Api
}

func NewService(repo *userRepo.Repository, api *api.Api) *Service {
	return &Service{
		repo: repo,
		api:  api,
	}
}
