package dto

type AuthDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RefreshTokenDto struct {
	Token string `json:"refresh_token"`
}
