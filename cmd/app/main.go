package main

import (
	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
	"github.com/jhseoeo/khuthon2023/internal/domain/service"
	"github.com/jhseoeo/khuthon2023/internal/infrastructure/persistence"
	"github.com/jhseoeo/khuthon2023/internal/interfaces/routes"
	"github.com/jhseoeo/khuthon2023/pkg/database"
	"github.com/joho/godotenv"
)

func wireUp(app *fiber.App) error {
	db, err := database.ConnectToDB()
	if err != nil {
		return errors.Errorf("failed to connect to database: %v", err)
	}
	userRepo := persistence.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	authService := usecase.NewAuthService(userService)
	authRoutes := routes.NewAuthRoutes(authService)

	app.Post("/auth/login", authRoutes.Login)
	app.Post("/auth/register", authRoutes.Register)

	return nil
}

func main() {
	godotenv.Load()
	app := fiber.New()
	wireUp(app)
	app.Listen(":8000")
}
