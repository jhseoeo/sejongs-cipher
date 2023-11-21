package port

import (
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type SignalingServicePort interface {
	JoinTetrisUser(sessionId string, conn ws.Ws) error
	JoinWordGuessUser(sessionId string, conn ws.Ws) error
	LeaveUser(session string, conn ws.Ws) error
	RelayMessage(sessionId string, conn ws.Ws, message entities.SignalingMessage) error
}
