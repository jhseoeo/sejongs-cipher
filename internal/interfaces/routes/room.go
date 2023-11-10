package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/usecase"
)

type RoomRoutes struct {
	roomUsecase *usecase.RoomUsecase
}

func NewRoomRoutes(roomUsecase *usecase.RoomUsecase) *RoomRoutes {

	return &RoomRoutes{
		roomUsecase: roomUsecase,
	}
}

// Room Create godoc
// @Summary Room Create
// @Description Room Create
// @Tags Room
// @Accept json
// @Produce json
// @Param body body request.RoomCreateRequest true "Room Create"
// @Success 200 {object} response.BaseResponse[response.RoomCreateResponse]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /room/create [post]
func (r *RoomRoutes) Create(c *fiber.Ctx) error {
	var req request.RoomCreateRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uuid.UUID)
	res, err := r.roomUsecase.Create(c.UserContext(), req, userId)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Room GetList godoc
// @Summary Room GetList
// @Description Room GetList
// @Tags Room
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse[response.RoomGetListResponse]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /room [get]
func (r *RoomRoutes) GetList(c *fiber.Ctx) error {
	res, err := r.roomUsecase.GetList(c.UserContext())
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Room Join godoc
// @Summary Room Join
// @Description Room Join
// @Tags Room
// @Accept json
// @Produce json
// @Param body body request.RoomJoinRequest true "Room Join"
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /room/join [post]
func (r *RoomRoutes) Join(c *fiber.Ctx) error {
	var req request.RoomJoinRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uuid.UUID)
	res, err := r.roomUsecase.Join(c.UserContext(), req, userId)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

// Room Leave godoc
// @Summary Room Leave
// @Description Room Leave
// @Tags Room
// @Accept json
// @Produce json
// @Param body body request.RoomLeaveRequest true "Room Leave"
// @Success 200 {object} response.BaseResponse[response.Empty]
// @Failure 400 {object} response.BaseResponse[response.Empty]
// @Failure 500 {object} response.BaseResponse[response.Empty]
// @Router /room/leave [post]
func (r *RoomRoutes) Leave(c *fiber.Ctx) error {
	var req request.RoomLeaveRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewErrorResponse[*response.Empty](400, "Bad Request"))
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(uuid.UUID)
	res, err := r.roomUsecase.Leave(c.UserContext(), req, userId)
	if err != nil {
		return c.Status(500).JSON(response.NewErrorResponse[*response.Empty](500, "Internal Server Error"))
	}
	return c.JSON(res)
}

func (r *RoomRoutes) WS(c *fiber.Ctx) error {
	return websocket.New(func(ws *websocket.Conn) {
		// TODO: implement
	})(c)
}
