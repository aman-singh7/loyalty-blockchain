package auth

import (
	domain "github.com/aman-singh7/loyalty-blockchain/domain/auth"
)

func toDomainCreateUserRequest(request *CreateUserRequest) *domain.CreateUserRequest {
	return &domain.CreateUserRequest{
		UID:   request.UID,
		Token: request.Token,
	}
}

func toDomainGenerateTokenRequest(request *GenerateTokenRequest) *domain.GenerateTokenRequest {
	return &domain.GenerateTokenRequest{
		Token: request.Token,
	}
}