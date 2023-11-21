package service

import (
	"context"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/repository"
)

const scoresPerPage = 20

// ScoreService is a service for score
type ScoreService struct {
	scoreRepository repository.ScoreRepository
}

// NewScoreService creates a new score service
func NewScoreService(scoreRepository repository.ScoreRepository) *ScoreService {
	return &ScoreService{
		scoreRepository: scoreRepository,
	}
}

// CreateScore creates a score
func (s *ScoreService) CreateScore(
	ctx context.Context,
	tetrisUsername string,
	wordguessUsername string,
	score int) error {
	id := uuid.New()
	score_ := &entities.Score{
		Id:                id,
		TetrisUsername:    tetrisUsername,
		WordguessUsername: wordguessUsername,
		Score:             score,
	}
	err := s.scoreRepository.Create(ctx, score_)
	if err != nil {
		return errors.Errorf("failed to create score: %v", err)
	}
	return nil
}

// GetScoreList gets a list of scores
func (s *ScoreService) GetScoreList(ctx context.Context, page int) ([]entities.Score, error) {
	offset := (page - 1) * scoresPerPage
	scores, err := s.scoreRepository.GetListWithOffset(ctx, offset, scoresPerPage)
	if err != nil {
		return nil, errors.Errorf("failed to get score list: %v", err)
	}
	return scores, nil
}
