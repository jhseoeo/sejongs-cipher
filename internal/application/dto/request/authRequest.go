package request

type AuthLoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
