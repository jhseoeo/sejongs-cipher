package entities

import (
	"time"

	"github.com/google/uuid"
)

// Score is a score
type Score struct {
	Id        uuid.UUID `gorm:"type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User1     User
	User2     User
	Score     int
}
