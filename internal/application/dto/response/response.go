package response

type response interface {
	*Empty | *AuthLoginResponse | *GetRanksResponse | *RoomCreateResponse |
		*RoomGetListResponse | *GameTestWordResponse
}

type Empty struct{}

type BaseResponse[T response] struct {
	Ok      bool   `json:"ok"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func NewEmptyBaseResponse[T response]() BaseResponse[T] {
	return BaseResponse[T]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data:    nil,
	}
}

func NewErrorResponse[T response](status int, message string) BaseResponse[T] {
	return BaseResponse[T]{
		Ok:      false,
		Status:  status,
		Message: message,
		Data:    nil,
	}
}
