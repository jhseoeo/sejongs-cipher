package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type BasicRoutes struct{}

// @title KHUTHON2023 API
// @version 1.0
// @description Khuthon 2023 API Server
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host jhseodev1.site
// @BasePath /api
// @schemes https

func NewBasicRoutes() *BasicRoutes {
	return &BasicRoutes{}
}

func (r *BasicRoutes) Docs(c *fiber.Ctx) error {
	return swagger.New(swagger.ConfigDefault)(c)
}

// Echo godoc
// @Summary Echo
// @Description Echo
// @Tags Basic
// @Accept json
// @Produce json
// @Param body body string true "Echo"
// @Success 200 {string} string "Echo"
// @Router /echo [post]
func (r *BasicRoutes) Echo(c *fiber.Ctx) error {
	return c.Send(c.Body())
}
