package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
)

type AuthRoutes struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthRoutes(authUsecase *usecase.AuthUsecase) *AuthRoutes {
	return &AuthRoutes{
		authUsecase: authUsecase,
	}
}

func (r *AuthRoutes) Login(c *fiber.Ctx) error {
	var req request.AuthLoginRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*struct{}](400, "Bad Request"))
	}
	res, err := r.authUsecase.Login(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*struct{}](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

func (r *AuthRoutes) Register(c *fiber.Ctx) error {
	var req request.AuthRegisterRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*struct{}](400, "Bad Request"))
	}
	res, err := r.authUsecase.Register(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*struct{}](500, "Internal Server Error"))
	}
	return c.JSON(res)
}
