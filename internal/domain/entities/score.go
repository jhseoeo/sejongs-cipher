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
	User1Id   uuid.UUID `gorm:"type:uuid;"`
	User2Id   uuid.UUID `gorm:"type:uuid;"`
	User1     User      `gorm:"foreignKey:User1Id;"`
	User2     User      `gorm:"foreignKey:User2Id;"`
	Score     int
}
