package usecase

import (
	"context"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
)

// RoomUsecase is a usecase for room
type RoomUsecase struct {
	userService port.UserServicePort
	roomService port.RoomServicePort

	subscriber []func(context.Context) error
}

// NewRoomUsecase creates a new room usecase
func NewRoomUsecase(
	userService port.UserServicePort,
	roomService port.RoomServicePort,
) *RoomUsecase {
	return &RoomUsecase{
		userService: userService,
		roomService: roomService,

		subscriber: make([]func(context.Context) error, 0),
	}
}

// Create creates a room
func (r *RoomUsecase) Create(ctx context.Context, req request.RoomCreateRequest, userId uuid.UUID) (response.BaseResponse[*response.RoomCreateResponse], error) {
	// get user
	user, err := r.userService.Get(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.RoomCreateResponse](500, "Internal Server Error"),
			errors.Errorf("failed to get user: %w", err)
	}

	// create room
	room, err := r.roomService.Create(ctx, user, req.RoomName)
	if err != nil {
		return response.NewErrorResponse[*response.RoomCreateResponse](500, "Internal Server Error"),
			errors.Errorf("failed to create room: %w", err)
	}

	return response.NewRoomCreateResponse(room), nil
}

// GetList gets a list of rooms
func (r *RoomUsecase) GetList(ctx context.Context) (response.BaseResponse[*response.RoomGetListResponse], error) {
	// get room list
	rooms, err := r.roomService.GetList(ctx)
	if err != nil {
		return response.NewErrorResponse[*response.RoomGetListResponse](500, "Internal Server Error"),
			errors.Errorf("failed to get room list: %w", err)
	}

	return response.NewRoomGetListResponse(rooms), nil
}

// Join joins a room
func (r *RoomUsecase) Join(ctx context.Context, req request.RoomJoinRequest, userId uuid.UUID) (response.BaseResponse[*response.Empty], error) {
	// get user
	user, err := r.userService.Get(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to get user: %w", err)
	}

	// join room
	err = r.roomService.Join(ctx, req.RoomId, user)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to join room: %w", err)
	}

	return response.NewEmptyBaseResponse[*response.Empty](), nil
}

// Leave leaves a room
func (r *RoomUsecase) Leave(ctx context.Context, req request.RoomLeaveRequest, userId uuid.UUID) (response.BaseResponse[*response.Empty], error) {
	// get user
	user, err := r.userService.Get(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to get user: %w", err)
	}

	// leave room
	_, err = r.roomService.Leave(ctx, req.RoomId, user)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to leave room: %w", err)
	}

	// TODO: let user know if room is deleted

	return response.NewEmptyBaseResponse[*response.Empty](), nil
}

func (r *RoomUsecase) SubscribeUserLeave(ctx context.Context, callback func(context.Context) error) {
	r.subscriber = append(r.subscriber, callback)
}
