package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/domain/repository"
)

// UserService is a service for user
type UserService struct {
	repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

// Create creates a user
func (s *UserService) Create(ctx context.Context, userId string, password string, userName string) error {
	id := uuid.New()
	pwHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password))) // sha256 hash
	user := entities.User{
		Id:       id,
		Username: userName,
		UserId:   userId,
		Password: pwHash,
	}
	err := s.UserRepository.Create(ctx, &user)
	if err != nil {
		return errors.Errorf("failed to create user: %v", err)
	}
	return nil
}

// Get gets a user
func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		return nil, errors.Errorf("failed to get user: %v", err)
	}
	return user, nil
}

// GetByUserId gets a user by user id
func (s *UserService) GetByUserId(ctx context.Context, userId string) (*entities.User, error) {
	user, err := s.UserRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, errors.Errorf("failed to get user: %v", err)
	}
	return user, nil
}
