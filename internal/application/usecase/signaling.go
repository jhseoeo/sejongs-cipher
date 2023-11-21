package usecase

import (
	"context"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/gofiber/contrib/websocket"
	wsmessage "github.com/jhseoeo/khuthon2023/internal/application/dto/wsMessage"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type SignalingUsecase struct {
	signalingService port.SignalingServicePort
}

func NewSignalingUsecase(signalingService port.SignalingServicePort) *SignalingUsecase {
	return &SignalingUsecase{
		signalingService: signalingService,
	}
}

// Join joins a signaling channel
func (s *SignalingUsecase) Join(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	if message.UserType == "tetris" {
		err := s.signalingService.JoinTetrisUser(sessionId, conn)
		if err != nil {
			return &wsmessage.WriteMessage{
				Status: 500,
				Ok:     false,
				Type:   "join",
				Data:   fmt.Sprintf("failed to join tetris user: %+v", err),
			}, errors.Errorf("failed to join tetris user: %w", err)
		}
	} else if message.UserType == "wordguess" {
		err := s.signalingService.JoinWordGuessUser(sessionId, conn)
		if err != nil {
			return &wsmessage.WriteMessage{
				Status: 500,
				Ok:     false,
				Type:   "join",
				Data:   fmt.Sprintf("failed to join wordguess user: %+v", err),
			}, errors.Errorf("failed to join wordguess user: %w", err)
		}
	} else {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "join",
			Data:   fmt.Sprintf("invalid user type: %s", message.UserType),
		}, errors.Errorf("invalid user type: %s", message.UserType)
	}

	return &wsmessage.WriteMessage{
		Status: 200,
		Ok:     true,
		Type:   "join",
	}, nil
}

// Leave leaves a signaling channel
func (s *SignalingUsecase) Leave(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	err := s.signalingService.LeaveUser(sessionId, conn)
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "leave",
			Data:   fmt.Sprintf("failed to leave user: %+v", err),
		}, errors.Errorf("failed to leave user: %w", err)
	}
	return &wsmessage.WriteMessage{
		Status: 200,
		Ok:     true,
		Type:   "leave",
	}, nil
}

func (s *SignalingUsecase) Relay(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	err := s.signalingService.RelayMessage(sessionId, conn, entities.SignalingMessage{
		Type: message.Type,
		Data: message.Data,
	})
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "relay",
			Data:   fmt.Sprintf("failed to relay message: %+v", err),
		}, errors.Errorf("failed to relay message: %w", err)
	}

	return &wsmessage.WriteMessage{
		Status: 200,
		Ok:     true,
		Type:   "relay",
	}, nil
}
