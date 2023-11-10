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

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body request.AuthLoginRequest true "Login"
// @Success 200 {object} response.BaseResponse[response.AuthLoginResponse]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /auth/login [post]
func (r *AuthRoutes) Login(c *fiber.Ctx) error {
	var req request.AuthLoginRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	res, err := r.authUsecase.Login(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body request.AuthRegisterRequest true "Register"
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /auth/register [post]
func (r *AuthRoutes) Register(c *fiber.Ctx) error {
	var req request.AuthRegisterRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	res, err := r.authUsecase.Register(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Check godoc
// @Summary Check
// @Description Check
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 401 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /auth/check [get]
func (r *AuthRoutes) Check(c *fiber.Ctx) error {
	return c.JSON(response.NewEmptyBaseResponse[*response.Empty]())
}
