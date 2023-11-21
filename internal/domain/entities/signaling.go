package entities

import (
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type SignalingSession struct {
	TetrisUser    ws.Ws
	WordGuessUser ws.Ws
}

type SignalingMessage struct {
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	UserType string      `json:"userType,omitempty"`
}
