package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewCORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowOrigins: "http://localhost:8000",
	})
}
