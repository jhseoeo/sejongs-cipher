package main

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
	"github.com/jhseoeo/khuthon2023/internal/domain/service"
	"github.com/jhseoeo/khuthon2023/internal/infrastructure/persistence"
	_ "github.com/jhseoeo/khuthon2023/internal/interfaces/docs"
	"github.com/jhseoeo/khuthon2023/internal/interfaces/routes"
	"github.com/jhseoeo/khuthon2023/pkg/database"
	"github.com/joho/godotenv"
)

func wireUp(app *fiber.App) error {
	db, err := database.ConnectToDB()
	if err != nil {
		return errors.Errorf("failed to connect to database: %v", err)
	}

	basicRoutes := routes.NewBasicRoutes()
	app.Get("/api/*", basicRoutes.Docs)

	userRepo := persistence.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	authService := usecase.NewAuthService(userService)
	authRoutes := routes.NewAuthRoutes(authService)
	app.Post("/api/auth/login", authRoutes.Login)
	app.Post("/api/auth/register", authRoutes.Register)

	return nil
}

func main() {
	godotenv.Load()
	app := fiber.New()
	err := wireUp(app)
	if err != nil {
		panic(fmt.Sprintf("failed to wire up: %+v", err))
	}
	app.Listen(":8000")
}
