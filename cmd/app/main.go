package main

import (
	"fmt"
	"os"

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
	app.Static("/", "./public")
	app.Static("/", "./../public")
	app.Static("/", "./../../public")

	api := app.Group("/api")
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtMiddleware := middlewares.NewJWTMiddleware(jwtSecret)

	basicRoutes := routes.NewBasicRoutes()
	api.Get("/docs/*", basicRoutes.Docs)
	api.Post("/echo", basicRoutes.Echo)

	userRepo := persistence.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	authUsecase := usecase.NewAuthService(userService)
	authRoutes := routes.NewAuthRoutes(authUsecase)
	auth := api.Group("/auth")
	auth.Post("/login", authRoutes.Login)
	auth.Post("/register", authRoutes.Register)
	auth.Get("/check", jwtMiddleware, authRoutes.Check)

	roomService := service.NewRoomService()
	roomUsecase := usecase.NewRoomUsecase(userService, roomService)
	roomRoutes := routes.NewRoomRoutes(roomUsecase)
	room := api.Group("/room", jwtMiddleware)
	room.Post("/create", roomRoutes.Create)
	room.Get("/", roomRoutes.GetList)
	room.Post("/", roomRoutes.Join)
	room.Post("/leave", roomRoutes.Leave)
	api.Get("/room/ws/:roomId", middlewares.NewWebsocketUpgradeMiddleware(), roomRoutes.WS())

	scoreRepo := persistence.NewScoreRepositoryImpl(db)
	scoreService := service.NewScoreService(scoreRepo)
	gameUsecase := usecase.NewGameUsecase(userService, roomService, scoreService)
	gameRoutes := routes.NewGameRoutes(gameUsecase)
	game := api.Group("/game", jwtMiddleware)
	game.Post("/start", gameRoutes.Start)
	game.Post("/end", gameRoutes.End)
	game.Post("/ranks", gameRoutes.Ranks)
	game.Post("/verify", gameRoutes.VerifyWord)
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
