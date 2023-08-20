package auth

import (
	"firebase.google.com/go/v4/auth"
	userDomain "github.com/aman-singh7/loyalty-blockchain/domain/user"
)

func newFirebaseUserMapper(user *auth.UserRecord) *userDomain.User {
	domainUser := userDomain.User{}

	domainUser.Name = user.DisplayName
	domainUser.Email = user.Email
	domainUser.AccountType = 1
	domainUser.PlatformUID = user.UID

	return &domainUser
}

func authUserMapper(domainUser *userDomain.User, auth *Auth) *AuthenticatedUser {
	return &AuthenticatedUser{
		Data: AuthenticatedUserData{
			UID:           domainUser.UID,
			Name:          domainUser.Name,
			Email:         domainUser.Email,
			PlatformUID:   domainUser.PlatformUID,
			AccountType:   domainUser.AccountType,
			WalletAddress: domainUser.WalletAddress,
		},
		Security: AuthenticatedUserToken{
			AccessToken:     auth.AccessToken,
			RefreshToken:    auth.RefreshToken,
			ExpAccessToken:  auth.ExpAccessToken,
			ExpRefreshToken: auth.ExpRefreshToken,
		},
	}
}
