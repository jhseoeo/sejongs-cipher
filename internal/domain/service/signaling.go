package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type SessionName string

type Session struct {
	TetrisUser    *websocket.Conn
	WordGuessUser *websocket.Conn
}

// Set of channels
type SignalingService struct {
	session sync.Map
}

func CreateSignalingService() *SignalingService {
	hub := SignalingService{
		session: sync.Map{},
	}

	return &hub
}

func (h *SignalingService) JoinTetrisUser(sessionName SessionName, conn *websocket.Conn) error {
	session, ok := h.session.Load(sessionName)
	if ok {
		session.(*Session).TetrisUser = conn
		h.session.Store(sessionName, session)
	} else {
		var newSession Session
		newSession.TetrisUser = conn
		h.session.Store(sessionName, &newSession)
	}
	fmt.Println("session", session, sessionName)

	return nil
}

func (h *SignalingService) JoinWordGuessUser(sessionName SessionName, conn *websocket.Conn) error {
	session, ok := h.session.Load(sessionName)
	if ok {
		session.(*Session).WordGuessUser = conn
		h.session.Store(sessionName, session)
	} else {
		var newSession Session
		newSession.WordGuessUser = conn
		h.session.Store(sessionName, &newSession)
	}

	return nil
}

func (h *SignalingService) SendTetrisSignalingMessage(session SessionName, message entities.SignalingMessage) error {
	sessionData, ok := h.session.Load(session)
	if !ok {
		return errors.New("session not found")
	}
	if sessionData.(*Session).WordGuessUser == nil {
		return errors.New("wordguess user not found")
	}
	err := sessionData.(*Session).WordGuessUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *SignalingService) SendWordGuessSignalingMessage(session SessionName, message entities.SignalingMessage) error {
	sessionData, ok := h.session.Load(session)
	if !ok {
		return errors.New("session not found")
	}
	if sessionData.(*Session).TetrisUser == nil {
		return errors.New("tetris user not found")
	}
	err := sessionData.(*Session).TetrisUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *SignalingService) LeaveUser(sessionName SessionName) error {
	sessionData, ok := h.session.Load(sessionName)
	if !ok {
		return errors.New("session not found")
	}
	session := sessionData.(*Session)
	if session.TetrisUser != nil {
		session.TetrisUser.Close()
	}
	if session.WordGuessUser != nil {
		session.WordGuessUser.Close()
	}
	h.session.Delete(sessionName)
	return nil
}
