package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/service"
)

func SetAuthRoutes(authService service.AuthService, app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/login", func(c *fiber.Ctx) error {
		var req request.AuthLoginRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(400).JSON(response.NewErrorResponse[*struct{}](400, "Bad Request"))
		}
		res, err := authService.Login(c.UserContext(), req.UserId, req.Password)
		if err != nil {
			return c.Status(500).JSON(response.NewErrorResponse[*struct{}](500, "Internal Server Error"))
		}
		return c.JSON(res)
	})

	api.Post("/register", func(c *fiber.Ctx) error {
		return c.SendString("Register")
	})
}
