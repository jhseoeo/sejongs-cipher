package middlewares

import (
	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
)

func NewWebsocketUpgradeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if ws.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}
