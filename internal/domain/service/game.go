package service

import (
	"errors"
	"sync"

	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/ws"
)

type GameService struct {
	session sync.Map
}

func NewGameService() *GameService {
	GameService := GameService{
		session: sync.Map{},
	}

	return &GameService
}

// JoinTetrisUser joins tetris user to session1
func (h *GameService) JoinTetrisUser(sessionName string, conn ws.Ws) error {
	sessionData, ok := h.session.Load(sessionName)
	if ok {
		session := sessionData.(*entities.GameSession)
		session.TetrisUser = conn
		h.session.Store(sessionName, session)
	} else {
		var newSession entities.GameSession
		newSession.TetrisUser = conn
		if newSession.WordGuessUser != nil {
			newSession.TetrisUser.WriteJSON(entities.GameMessage{
				Type: "join",
				Data: "wordguess user has joined",
			})
		}
		h.session.Store(sessionName, &newSession)
	}
	return nil
}

// JoinWordGuessUser joins wordguess user to session
func (h *GameService) JoinWordGuessUser(sessionName string, conn ws.Ws) error {
	session, ok := h.session.Load(sessionName)
	if ok {
		session.(*entities.GameSession).WordGuessUser = conn
		h.session.Store(sessionName, session)
		if session.(*entities.GameSession).TetrisUser != nil {
			session.(*entities.GameSession).TetrisUser.WriteJSON(entities.GameMessage{
				Type: "join",
				Data: "word guess user has joined",
			})
		}
	} else {
		var newSession entities.GameSession
		newSession.WordGuessUser = conn
		h.session.Store(sessionName, &newSession)
	}

	return nil
}

// RelayMessage relays message to other user in session
func (h *GameService) RelayMessage(sessionName string, conn ws.Ws, message entities.GameMessage) error {
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}

	// get session
	session := sessionData.(*entities.GameSession)
	if !ok {
		return errors.New("session not found")
	}

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
		err := sessionData.(*entities.GameSession).WordGuessUser.WriteJSON(message)
		if err != nil {
			return err
		}

	} else if userType == "wordguess" {
		if session.TetrisUser == nil {
			return errors.New("tetris user not found")
		}
		err := sessionData.(*entities.GameSession).TetrisUser.WriteJSON(message)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid user type")
	}

	return nil
}

func (h *GameService) LeaveUser(sessionName string, conn ws.Ws) error {
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
