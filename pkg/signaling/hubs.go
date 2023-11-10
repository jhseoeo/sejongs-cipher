package signaling

import (
	"errors"
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
	mutex   sync.RWMutex
	session map[SessionName]Session
}

type MessageData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func CreateHub() *Hub {
	hub := Hub{
		session: make(map[SessionName]Session),
	}

	return &hub
}

func (h *Hub) JoinTetrisUser(sessionName SessionName, conn *websocket.Conn) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	session, ok := h.session[sessionName]
	if !ok {
		session = Session{
			TetrisUser: conn,
		}
	} else {
		session.TetrisUser = conn
		h.session[sessionName] = session
	}

	return nil
}

func (h *Hub) JoinWordGuessUser(sessionName SessionName, conn *websocket.Conn) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	session, ok := h.session[sessionName]
	if !ok {
		session = Session{
			WordGuessUser: conn,
		}
	} else {
		session.WordGuessUser = conn
		h.session[sessionName] = session
	}

	return nil
}

func (h *Hub) SendTetrisSignalingMessage(session SessionName, message MessageData) error {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	sessionData, ok := h.session[session]
	if !ok {
		return errors.New("session not found")
	}
	if sessionData.WordGuessUser == nil {
		return errors.New("wordguess user not found")
	}
	err := sessionData.WordGuessUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hub) SendWordGuessSignalingMessage(session SessionName, message MessageData) error {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	sessionData, ok := h.session[session]
	if !ok {
		return errors.New("session not found")
	}
	if sessionData.TetrisUser == nil {
		return errors.New("tetris user not found")
	}
	err := sessionData.TetrisUser.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hub) LeaveUser(sessionName SessionName) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	session, ok := h.session[sessionName]
	if !ok {
		return errors.New("session not found")
	}
	session.TetrisUser.Close()
	session.WordGuessUser.Close()
	delete(h.session, sessionName)
	return nil
}
