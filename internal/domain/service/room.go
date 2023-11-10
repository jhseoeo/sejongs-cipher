package service

import (
	"context"
	"sync"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

// RoomService is a service for room
type RoomService struct {
	rooms sync.Map
}

// NewRoomService creates a new room service
func NewRoomService() *RoomService {
	return &RoomService{}
}

// Create creates a room
func (s *RoomService) Create(ctx context.Context, user1 *entities.User, roomName string) (*entities.Room, error) {
	id := uuid.New()
	room := entities.Room{
		Id:       id,
		RoomName: roomName,
		User1:    *user1,
	}
	s.rooms.Store(id, room)
	return &room, nil
}

// Get gets a room
func (s *RoomService) Get(ctx context.Context, id uuid.UUID) (*entities.Room, error) {
	room, ok := s.rooms.Load(id)
	if !ok {
		return &entities.Room{}, errors.New("room not found")
	}
	ret := room.(entities.Room)
	return &ret, nil
}

// GetList gets a list of rooms
func (s *RoomService) GetList(ctx context.Context) ([]entities.Room, error) {
	roomsList := []entities.Room{}
	s.rooms.Range(func(key, value any) bool {
		roomsList = append(roomsList, value.(entities.Room))
		return true
	})
	return roomsList, nil
}

// Join joins a room
func (s *RoomService) Join(ctx context.Context, id uuid.UUID, user2 *entities.User) error {
	room, ok := s.rooms.Load(id)
	if !ok {
		return errors.New("room not found")
	}

	// if user2 is already in a room, it is not allowed to join
	roomToJoin := room.(entities.Room)
	if roomToJoin.User2 != nil {
		return errors.New("room is full")
	}
	roomToJoin.User2 = user2
	s.rooms.Store(id, roomToJoin)
	return nil
}

// Leave leaves a room
func (s *RoomService) Leave(ctx context.Context, id uuid.UUID, user *entities.User) (bool, error) {
	room, ok := s.rooms.Load(id)
	if !ok {
		return false, errors.New("room not found")
	}
	roomToLeave := room.(entities.Room)
	if roomToLeave.User1.Id == user.Id {
		// if user is user1, delete the room
		s.rooms.Delete(id)
		return true, nil

	} else if roomToLeave.User2.Id == user.Id {
		// if user is user2, delete user2
		roomToLeave.User2 = nil
		s.rooms.Store(id, roomToLeave)
		return false, nil

	} else {
		return false, errors.New("user is not in room")
	}
}

// StartGame starts a game
func (s *RoomService) StartGame(ctx context.Context, roomId uuid.UUID, user *entities.User) error {
	room, ok := s.rooms.Load(roomId)
	if !ok {
		return errors.New("room not found")
	}
	roomToStart := room.(entities.Room)
	if roomToStart.User1.Id != user.Id {
		// if user is not user1, it is not allowed to start game
		return errors.New("user is not allowed to start game")
	} else if roomToStart.User2 == nil {
		return errors.New("room is not full")
	}
	roomToStart.IsPlaying = true
	s.rooms.Store(roomId, roomToStart)
	return nil
}

// EndGame ends a game
func (s *RoomService) EndGame(ctx context.Context, roomId uuid.UUID, user *entities.User) error {
	room, ok := s.rooms.Load(roomId)
	if !ok {
		return errors.New("room not found")
	}
	roomToEnd := room.(entities.Room)
	if roomToEnd.User1.Id != user.Id {
		// if user is not user1, it is not allowed to end game
		return errors.New("user is not allowed to end game")
	}
	roomToEnd.IsPlaying = false
	s.rooms.Store(roomId, roomToEnd)
	return nil
}
