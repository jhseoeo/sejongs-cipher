package wsmessage

import "github.com/jhseoeo/khuthon2023/internal/domain/entities"

type ReadMessage struct {
	Type     string      `json:"type"`
	UserType string      `json:"userType,omitempty"`
	Word     string      `json:"word,omitempty"`
	Score    string      `json:"score,omitempty"`
	Data     interface{} `json:"data"`
}

type WriteMessage struct {
	Status int         `json:"status"`
	Ok     bool        `json:"ok"`
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`

	// VerifyWord
	IsCorrect   bool   `json:"isCorrect,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m *ReadMessage) ToGameMessage() *entities.GameMessage {
	return &entities.GameMessage{
		Type: m.Type,
		Data: m.Data,
	}
}

func NewWriteMessageFromGameMessage(m *entities.GameMessage) *WriteMessage {
	return &WriteMessage{
		Type: m.Type,
		Data: m.Data,
	}
}
