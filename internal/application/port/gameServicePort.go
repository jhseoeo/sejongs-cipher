package port

import (
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type GameServicePort interface {
	JoinTetrisUser(sessionName string, username string, conn ws.Ws) error
	JoinWordGuessUser(sessionName string, username string, conn ws.Ws) error
	GetUserName(sessionId string, conn ws.Ws) (tetrisUsername string, wordguessUsername string, err error)
	LeaveUser(session string, conn ws.Ws) error
	RelayMessage(session string, conn ws.Ws, message entities.GameMessage) error
}
