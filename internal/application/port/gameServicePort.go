package port

import (
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type GameServicePort interface {
	JoinTetrisUser(sessionName string, conn ws.Ws) error
	JoinWordGuessUser(sessionName string, conn ws.Ws) error
	LeaveUser(session string, conn ws.Ws) error
	RelayMessage(session string, conn ws.Ws, message entities.GameMessage) error
}
