package websocketHandler

import (
	"context"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	wsmessage "github.com/jhseoeo/khuthon2023/internal/application/dto/wsMessage"
)

type wsHandler func(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error)

type WebsocketHandlerBuilder struct {
	joinHandler    wsHandler
	leaveHandler   wsHandler
	handler        map[string]wsHandler
	defaultHandler wsHandler
}

func NewWebsocketHandlerBuilder() *WebsocketHandlerBuilder {
	return &WebsocketHandlerBuilder{}
}

func (b *WebsocketHandlerBuilder) OnJoin(handler wsHandler) *WebsocketHandlerBuilder {
	b.joinHandler = handler
	return b
}

func (b *WebsocketHandlerBuilder) OnLeave(handler wsHandler) *WebsocketHandlerBuilder {
	b.leaveHandler = handler
	return b
}

func (b *WebsocketHandlerBuilder) OnMessageType(messageType string, handler wsHandler) *WebsocketHandlerBuilder {
	b.handler[messageType] = handler
	return b
}

func (b *WebsocketHandlerBuilder) OnDefault(handler wsHandler) *WebsocketHandlerBuilder {
	b.defaultHandler = handler
	return b
}

func (b *WebsocketHandlerBuilder) Build() func(sessionId string, userId string, ws *websocket.Conn) {
	return func(sessionId string, userId string, ws *websocket.Conn) {
		ctx, cancel := context.WithCancel(context.Background())

		var messageData wsmessage.ReadMessage
		err := ws.ReadJSON(&messageData)
		if err != nil {
			fmt.Println("read error:", err)
			ws.Close()
		}
		_, err = b.joinHandler(ctx, ws, sessionId, userId, messageData)
		if err != nil {
			fmt.Println("an error occurred while handling user join:", err)
			ws.Close()
		}

		fmt.Printf("%s user joined on %s\n", userId, sessionId)

		defer func() {
			fmt.Printf("%s user leaved from %s\n", userId, sessionId)
			_, err := b.leaveHandler(ctx, ws, sessionId, userId, messageData)
			if err != nil {
				fmt.Println("an error occurred while handling user leave:", err)
			}
			ws.Close()
			cancel()
		}()

		for {
			var message wsmessage.ReadMessage
			err := ws.ReadJSON(&message)
			if err != nil {
				fmt.Println("read error:", err)
				return
			}

			for messageType, handler := range b.handler {
				if message.Type == messageType {
					res, err := handler(ctx, ws, sessionId, userId, message)
					if err != nil {
						if res != nil {
							err = ws.WriteJSON(res)
							if err != nil {
								fmt.Println("an error occurred while writing message:", err)
								return
							}
						}
						fmt.Println("an error occurred while handling message:", err)
						return
					}

					if res != nil {
						err = ws.WriteJSON(res)
						if err != nil {
							fmt.Println("an error occurred while writing message:", err)
							return
						}
					}

					break
				}
			}
		}
	}
}
