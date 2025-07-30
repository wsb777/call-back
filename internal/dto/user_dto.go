package dto

type UserSignUpDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserSignInDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
