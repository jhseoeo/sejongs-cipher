package gameWs

import (
	"fmt"

	ws "github.com/gofiber/contrib/websocket"
)

type userType string // tetris, wordguess

type userJoinMessage struct {
	UserType string `json:"userType"`
}

// When a client has connected, client send its uuid.
// Get uuid from this message.
func handleUserJoin(hub *Hub, conn *ws.Conn, sessionName SessionName) (userType, error) {
	var messageData userJoinMessage
	err := conn.ReadJSON(&messageData)
	if err != nil {
		if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
			fmt.Println("read error:", err)
		}
		return "", err
	}

	if messageData.UserType == "tetris" {
		hub.JoinTetrisUser(sessionName, conn)
	} else if messageData.UserType == "wordguess" {
		hub.JoinWordGuessUser(sessionName, conn)
	} else {
		return "", fmt.Errorf("invalid user type: %s", messageData.UserType)
	}

	return userType(messageData.UserType), nil
}

// Websocket Session Loop for each client
func WebsocketConnectionLoop(hub *Hub, conn *ws.Conn) {
	session := SessionName(conn.Params("session"))
	user, err := handleUserJoin(hub, conn, session)
	if err != nil {
		fmt.Println("an error occurred while handling user join:", err)
		conn.Close()
	}

	fmt.Printf("%s user joined on %s\n", user, session)

	defer func() { // when user leaves
		fmt.Printf("%s user leaved from %s\n", user, session)
		if user == "tetris" {
			hub.LeaveUser(session)
		}
		conn.Close()
	}()

	for {
		var messageData MessageData
		err := conn.ReadJSON(&messageData)

		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				fmt.Println("read error:", err)
			}
			return
		}

		hub.SendTetrisMessage(session, messageData)
	}
}
