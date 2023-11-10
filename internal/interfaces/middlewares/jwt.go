package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func NewJWTMiddleware(jwtSecret string) fiber.Handler {
	// return jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{
	// 		Key: []byte(jwtSecret),
	// 	},
	// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 		fmt.Printf("jwt error: %v\n", err)
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 			"message": "Unauthorized",
	// 		})
	// 	},
	// })
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
