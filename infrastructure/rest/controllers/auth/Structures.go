package auth

type CreateUserRequest struct {
	UID   string `json:"uid"`
	Token string `json:"tokenId"`
}

type GenerateTokenRequest struct {
	Token string `json:"refreshToken"`
}
