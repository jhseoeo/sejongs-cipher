package port

import (
	"context"

	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
)

type ScoreServicePort interface {
	CreateScore(ctx context.Context, tetrisUsername string, wordguessUsername string, score int) error
	GetScoreList(ctx context.Context, page int) ([]entities.Score, error)
}
