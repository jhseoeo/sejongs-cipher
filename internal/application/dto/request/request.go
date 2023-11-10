package request

type WebSocketRequestMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
