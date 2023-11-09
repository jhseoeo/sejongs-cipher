package entities

import (
	"time"

	"github.com/google/uuid"
)

type Score struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User1     User
	User2     User
	Score     int
}
