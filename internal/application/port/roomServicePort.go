package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type RoomServicePort interface {
	Create(ctx context.Context, user1 *entities.User, roomName string) (*entities.Room, error)
	Get(ctx context.Context, id uuid.UUID) (*entities.Room, error)
	GetList(ctx context.Context) ([]entities.Room, error)
	Join(ctx context.Context, id uuid.UUID, user2 *entities.User) error
	Leave(ctx context.Context, id uuid.UUID, user *entities.User) (bool, error)
	StartGame(ctx context.Context, id uuid.UUID, user *entities.User) error
	EndGame(ctx context.Context, id uuid.UUID, user *entities.User) error
}
