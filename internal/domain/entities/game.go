package entities

import (
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type GameSession struct {
	TetrisUser    ws.Ws
	WordGuessUser ws.Ws
}

type GameMessage struct {
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
	UserType string      `json:"userType,omitempty"`
}
