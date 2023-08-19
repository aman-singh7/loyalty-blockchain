package brand

import brandRepo "github.com/aman-singh7/loyalty-blockchain/infrastructure/repository/brand"

type Service struct {
	repo *brandRepo.Repository
}

func NewService(repo *brandRepo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
