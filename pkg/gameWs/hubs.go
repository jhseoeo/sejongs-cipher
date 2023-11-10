package gameWs

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type SessionName string

type Session struct {
	TetrisUser    *websocket.Conn
	WordGuessUser *websocket.Conn
}

// Set of channels
type Hub struct {
	session sync.Map
}

type MessageData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func CreateHub() *Hub {
	hub := Hub{
		session: sync.Map{},
	}

	return &hub
}

func (h *Hub) JoinTetrisUser(sessionName SessionName, conn *websocket.Conn) error {
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

func (h *Hub) JoinWordGuessUser(sessionName SessionName, conn *websocket.Conn) error {
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

func (h *Hub) SendTetrisMessage(session SessionName, message MessageData) error {
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

func (h *Hub) SendWordGuessMessage(session SessionName, message MessageData) error {
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

func (h *Hub) LeaveUser(sessionName SessionName) error {
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
