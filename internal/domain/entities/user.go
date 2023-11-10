package entities

import (
	"time"

	"github.com/google/uuid"
)

// User is a user
type User struct {
	Id        uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"unique"`
	UserId    string `gorm:"unique"`
	Password  string
}
