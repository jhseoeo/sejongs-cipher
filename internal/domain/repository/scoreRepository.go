package repository

import (
	"context"

	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

// UserRepository is a repository for user
type ScoreRepository interface {
	Create(ctx context.Context, score *entities.Score) error
	GetListWithOffset(ctx context.Context, offset, count int) ([]entities.Score, error)
}
