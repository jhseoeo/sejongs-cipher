package entities

import (
	"time"

	"github.com/google/uuid"
)

// Score is a score
type Score struct {
	Id                uuid.UUID `gorm:"type:uuid;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	TetrisUsername    string
	WordguessUsername string
	Score             int
}
