package dto

type AuthRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
