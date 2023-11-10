package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
	"github.com/jhseoeo/khuthon2023/pkg/gameWs"
	"github.com/jhseoeo/khuthon2023/pkg/signaling"
)

type GameRoutes struct {
	signalingHub *signaling.Hub
	gameWsHub    *gameWs.Hub
	gameUsecase  *usecase.GameUsecase
}

func NewGameRoutes(gameUsecase *usecase.GameUsecase) *GameRoutes {
	return &GameRoutes{
		signalingHub: signaling.CreateHub(),
		gameWsHub:    gameWs.CreateHub(),
		gameUsecase:  gameUsecase,
	}
}

// Game Start godoc
// @Summary Game Start
// @Description Game Start
// @Tags Game
// @Accept json
// @Produce json
// @Param body body request.GameStartRequest true "Game Start"
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /game/start [post]
func (r *GameRoutes) Start(c *fiber.Ctx) error {
	var req request.GameStartRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uuid.UUID)
	res, err := r.gameUsecase.Start(c.UserContext(), req, userId)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Game End godoc
// @Summary Game End
// @Description Game End
// @Tags Game
// @Accept json
// @Produce json
// @Param body body request.GameEndRequest true "Game End"
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /game/end [post]
func (r *GameRoutes) End(c *fiber.Ctx) error {
	var req request.GameEndRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uuid.UUID)
	res, err := r.gameUsecase.End(c.UserContext(), req, userId)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Game GetRanks godoc
// @Summary Game GetRanks
// @Description Game GetRanks
// @Tags Game
// @Accept json
// @Produce json
// @Param body body request.GetRanksRequest true "Game GetRanks"
// @Success 200 {object} response.BaseResponse[response.GetRanksResponse]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /game/ranks [post]
func (r *GameRoutes) Ranks(c *fiber.Ctx) error {
	var req request.GetRanksRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	res, err := r.gameUsecase.GetRanks(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Game VerifyWord godoc
// @Summary Game VerifyWord
// @Description Game VerifyWord
// @Tags Game
// @Accept json
// @Produce json
// @Param body body request.GameTestWordRequest true "Game VerifyWord"
// @Success 200 {object} response.BaseResponse[response.GameTestWordResponse]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /game/verify [post]
func (r *GameRoutes) VerifyWord(c *fiber.Ctx) error {
	var req request.GameTestWordRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	res, err := r.gameUsecase.VerifyWord(c.UserContext(), req)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res.Data)
}

func (r *GameRoutes) Signaling() func(c *fiber.Ctx) error {
	return websocket.New(func(ws *websocket.Conn) {
		signaling.WebsocketConnectionLoop(r.signalingHub, ws)
	})
}

func (r *GameRoutes) WS() func(c *fiber.Ctx) error {
	return websocket.New(func(ws *websocket.Conn) {
		gameWs.WebsocketConnectionLoop(r.gameWsHub, ws)
	})
}
