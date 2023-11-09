package request

type AuthLoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type AuthRegisterRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}
