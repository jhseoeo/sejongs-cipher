package main

import (
	"fmt"

	"github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
	"github.com/jhseoeo/khuthon2023/internal/domain/service"
	"github.com/jhseoeo/khuthon2023/internal/infrastructure/persistence"
	_ "github.com/jhseoeo/khuthon2023/internal/interfaces/docs"
	"github.com/jhseoeo/khuthon2023/internal/interfaces/middlewares"
	"github.com/jhseoeo/khuthon2023/internal/interfaces/routes"
	"github.com/jhseoeo/khuthon2023/pkg/database"
	"github.com/joho/godotenv"
)

func wireUp(app *fiber.App) error {
	db, err := database.ConnectToDB()
	if err != nil {
		return errors.Errorf("failed to connect to database: %v", err)
	}
	app.Use(middlewares.NewCORSMiddleware())

	api := app.Group("/api")

	basicRoutes := routes.NewBasicRoutes()
	api.Get("/docs/*", basicRoutes.Docs)
	api.Post("/echo", basicRoutes.Echo)

	scoreRepo := persistence.NewScoreRepositoryImpl(db)
	scoreService := service.NewScoreService(scoreRepo)
	gameService := service.NewGameService()
	gameUsecase := usecase.NewGameUsecase(scoreService, gameService)
	gameRoutes := routes.NewGameRoutes(gameUsecase)
	game := api.Group("/game")
	game.Post("/ranks", gameRoutes.Ranks)
	api.Get("/game/ws/:session", middlewares.NewWebsocketUpgradeMiddleware(), gameRoutes.WS())
	api.Get("/game/signaling/:session", middlewares.NewWebsocketUpgradeMiddleware(), gameRoutes.Signaling())

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
