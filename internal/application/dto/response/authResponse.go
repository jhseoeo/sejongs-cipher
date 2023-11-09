package response

type AuthLoginResponse struct {
	Token string `json:"token,omitempty"`
}

func NewAuthLoginResponse(token string) BaseResponse[*AuthLoginResponse] {
	return BaseResponse[*AuthLoginResponse]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data:    &AuthLoginResponse{Token: token},
	}
}
