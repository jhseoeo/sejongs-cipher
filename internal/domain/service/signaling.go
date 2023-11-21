package service

import (
	"errors"
	"sync"

	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

// Set of channels
type SignalingService struct {
	session sync.Map
}

func NewSignalingService() *SignalingService {
	hub := SignalingService{
		session: sync.Map{},
	}

	return &hub
}

func (h *SignalingService) JoinTetrisUser(sessionName string, conn ws.Ws) error {
	sessionData, ok := h.session.Load(sessionName)
	if ok {
		session := sessionData.(*entities.SignalingSession)
		session.TetrisUser = conn
		h.session.Store(sessionName, session)
	} else {
		var newSession entities.SignalingSession
		newSession.TetrisUser = conn
		h.session.Store(sessionName, &newSession)
	}
	return nil
}

func (h *SignalingService) JoinWordGuessUser(sessionName string, conn ws.Ws) error {
	sessionData, ok := h.session.Load(sessionName)
	if ok {
		session := sessionData.(*entities.SignalingSession)
		session.WordGuessUser = conn
		h.session.Store(sessionName, session)
	} else {
		var newSession entities.SignalingSession
		newSession.WordGuessUser = conn
		h.session.Store(sessionName, &newSession)
	}

	return nil
}

func (h *SignalingService) SendTetrisSignalingMessage(sessionName string, message entities.SignalingMessage) error {
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}
	session := sessionData.(*entities.SignalingSession)
	if session.WordGuessUser == nil {
		return errors.New("wordguess user not found")
	}
	err := session.WordGuessUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *SignalingService) SendWordGuessSignalingMessage(sessionName string, message entities.SignalingMessage) error {
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}
	session := sessionData.(*entities.SignalingSession)
	if session.TetrisUser == nil {
		return errors.New("tetris user not found")
	}
	err := session.TetrisUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *SignalingService) RelayMessage(sessionName string, conn ws.Ws, message entities.SignalingMessage) error {
	// get session
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}
	session := sessionData.(*entities.SignalingSession)

	// check user type
	var userType string
	if session.TetrisUser == conn {
		userType = "tetris"
	} else if session.WordGuessUser == conn {
		userType = "wordguess"
	} else {
		return errors.New("user not in session")
	}

	// relay message according to user type
	if userType == "tetris" {
		if session.WordGuessUser == nil {
			return errors.New("wordguess user not found")
		}
		err := session.WordGuessUser.WriteJSON(message)
		if err != nil {
			return err
		}

	} else if userType == "wordguess" {
		if session.TetrisUser == nil {
			return errors.New("tetris user not found")
		}
		err := session.TetrisUser.WriteJSON(message)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid user type")
	}

	return nil
}

func (h *SignalingService) LeaveUser(sessionName string, conn ws.Ws) error {
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}
	session := sessionData.(*entities.GameSession)
	if session.TetrisUser == conn {
		if session.WordGuessUser != nil {
			err := session.WordGuessUser.WriteJSON(entities.GameMessage{
				Type: "leave",
				Data: "tetris user has left",
			})
			if err != nil {
				return err
			}
		}
		err := session.TetrisUser.Close()
		if err != nil {
			return err
		}

	} else if session.WordGuessUser == conn {
		if session.TetrisUser != nil {
			err := session.TetrisUser.WriteJSON(entities.GameMessage{
				Type: "leave",
				Data: "wordguess user has left",
			})
			if err != nil {
				return err
			}
		}
		err := session.WordGuessUser.Close()
		if err != nil {
			return err
		}

	} else {
		return errors.New("user not found")
	}
	h.session.Delete(sessionName)
	return nil
}
