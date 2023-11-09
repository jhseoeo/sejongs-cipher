package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByUserId(ctx context.Context, userid string) (*entities.User, error)
}
