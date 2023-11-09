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

type UserService struct {
	repository.UserRepository
}

func NewUserService(
	userRepository repository.UserRepository,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Create(ctx context.Context, userId string, password string, userName string) error {
	id := uuid.New()
	pwHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	user := entities.User{
		Id:       id,
		Username: userName,
		UserId:   userId,
		Password: pwHash,
	}
	err := s.UserRepository.CreateUser(ctx, &user)
	if err != nil {
		return errors.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		return nil, errors.Errorf("failed to get user: %v", err)
	}
	return user, nil
}

func (s *UserService) GetByUserId(ctx context.Context, userId string) (*entities.User, error) {
	user, err := s.UserRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, errors.Errorf("failed to get user: %v", err)
	}
	return user, nil
}
