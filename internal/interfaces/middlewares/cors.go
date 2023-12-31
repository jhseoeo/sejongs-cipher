package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS",
	})
}
