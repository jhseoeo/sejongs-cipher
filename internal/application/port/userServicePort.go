package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type UserServicePort interface {
	Create(ctx context.Context, userId string, password string, userName string) error
	Get(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetByUserId(ctx context.Context, userId string) (*entities.User, error)
}
