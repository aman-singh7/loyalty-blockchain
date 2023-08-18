package user

import userRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/user"

type Service struct {
	repo *userRepo.Repository
}

func NewService(repo *userRepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
