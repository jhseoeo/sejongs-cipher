package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

// UserRepository is a repository for user
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByUserId(ctx context.Context, userId string) (*entities.User, error)
}
