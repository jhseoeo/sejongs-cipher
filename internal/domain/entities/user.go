package entities

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `gorm:"unique"`
	UserId   string    `gorm:"unique"`
	Password string
}
