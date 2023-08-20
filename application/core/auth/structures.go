package auth

import "time"

type AuthenticatedUserData struct {
	UID           string `json:"uid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PlatformUID   string `json:"platformUid"`
	AccountType   uint8  `json:"accountType"`
	WalletAddress string `json:"walletAddress"`
}

type AuthenticatedUserToken struct {
	AccessToken     string    `json:"accessToken"`
	RefreshToken    string    `json:"refreshToken"`
	ExpAccessToken  time.Time `json:"expAccessToken"`
	ExpRefreshToken time.Time `json:"expRefreshToken"`
}

type AuthenticatedUser struct {
	Data     AuthenticatedUserData  `json:"data"`
	Security AuthenticatedUserToken `json:"token"`
}
