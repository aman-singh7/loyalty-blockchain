package auth

type CreateUserRequest struct {
	UID   string `json:"uid"`
	Token string `json:"token"`
}

type GenerateTokenRequest struct {
	Token string `json:"refreshToken"`
}