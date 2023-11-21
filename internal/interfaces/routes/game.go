package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
	"github.com/jhseoeo/khuthon2023/internal/interfaces/websocketHandler"
)

type GameRoutes struct {
	gameUsecase      *usecase.GameUsecase
	signalingUsecase *usecase.SignalingUsecase
}

func NewGameRoutes(gameUsecase *usecase.GameUsecase, signalingUsecase *usecase.SignalingUsecase) *GameRoutes {
	return &GameRoutes{
		gameUsecase:      gameUsecase,
		signalingUsecase: signalingUsecase,
	}
}

// Ranks godoc
// @Summary Get ranks
// @Description Get ranks
// @Tags Game
// @Accept json
// @Produce json
// @Param body body request.GetRanksRequest true "body"
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

// Signaling godoc
// @Summary Signaling Channel
// @Description Websocket Signaling Server for WebRTC Communication
// @Description Use this endpoint as signaling server url
// @Tags Game
// @Router /game/signaling [get]
func (r *GameRoutes) Signaling() func(c *fiber.Ctx) error {
	return websocket.New(func(ws *websocket.Conn) {
		sessionId := ws.Params("sessionId")
		userId := ws.Query("userId")
		handler := websocketHandler.NewWebsocketHandlerBuilder().
			OnJoin(r.signalingUsecase.Join).
			OnLeave(r.signalingUsecase.Leave).
			OnDefault(r.signalingUsecase.Relay).
			Build()

		handler(sessionId, userId, ws)

	})
}

// Game-Server godoc
// @Summary Game Server
// @Description Websocket Endpoint for In-game Communication
// @Description Use this endpoint as signaling server url
// @Tags Game
// @Router /game/ws [get]
func (r *GameRoutes) WS() func(c *fiber.Ctx) error {
	return websocket.New(func(ws *websocket.Conn) {
		sessionId := ws.Params("sessionId")
		userId := ws.Query("userId")

		handler := websocketHandler.NewWebsocketHandlerBuilder().

			// Req: { "type": "join", "userType": "tetris" | "wordguess" }
			// tetris 유저는 wordguess 유저가 들어오면 받음 -> { "type": "join", "data": "wordguess user has joined" }
			// wordguess 유저가 이미 들어와 있으면, 들어옴과 동시에 받음 -> { "type": "join", "data": "wordguess user has joined" }
			OnJoin(r.gameUsecase.Join).

			// 그냥 나가면 됨
			// 상대편 유저에게 전달: { "type": "leave", "data": "tetris|wordguess user has left" }
			OnLeave(r.gameUsecase.Leave).

			// Req: { "type": "verify", "word": "단어" }
			// Res: { "type": "verify", "isCorrect": true | false, "description": "설명" }
			OnMessageType("verify", r.gameUsecase.VerifyWord).

			// Req: { "type": "end", score: 점수(숫자) }
			OnMessageType("end", r.gameUsecase.VerifyWord).

			// Req: { "type": "위 제외 아무거나", "data": 아무거나 } -> 그대로 relay
			OnDefault(r.gameUsecase.Relay).
			Build()

		handler(sessionId, userId, ws)
	})
}
