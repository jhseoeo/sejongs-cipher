package persistence

import (
	"context"

	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"gorm.io/gorm"
)

// ScoreRepositoryImpl is a implementation of ScoreRepository
type ScoreRepositoryImpl struct {
	db *gorm.DB
}

// NewScoreRepositoryImpl creates a new score repository
func NewScoreRepositoryImpl(db *gorm.DB) *ScoreRepositoryImpl {
	db.AutoMigrate(&entities.Score{})
	return &ScoreRepositoryImpl{db: db}
}

// Create creates a score
func (r *ScoreRepositoryImpl) Create(ctx context.Context, score *entities.Score) error {
	return r.db.WithContext(ctx).Create(score).Error
}

// GetList gets a list of scores
func (r *ScoreRepositoryImpl) GetListWithOffset(ctx context.Context, offset, count int) ([]entities.Score, error) {
	var scores []entities.Score
	err := r.db.WithContext(ctx).Offset(offset).Limit(count).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}
