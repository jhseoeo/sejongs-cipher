package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"gorm.io/gorm"
)

// UserRepositoryImpl is a implementation of UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepositoryImpl creates a new user repository
func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&entities.User{})
	return &UserRepositoryImpl{db: db}
}

// Create creates a user
func (r *UserRepositoryImpl) Create(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindById finds a user by id
func (r *UserRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUserId finds a user by user id
func (r *UserRepositoryImpl) FindByUserId(ctx context.Context, userId string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
