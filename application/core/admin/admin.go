package admin

import (
	"github.com/aman-singh7/loyalty-blockchain/application/api"
	adminRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/admin"
)

type Service struct {
	repo *adminRepo.Repository
	api  *api.Service
}

func NewService(repo *adminRepo.Repository, api *api.Service) *Service {
	return &Service{
		repo: repo,
		api:  api,
	}
}
