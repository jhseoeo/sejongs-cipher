package entities

import "github.com/google/uuid"

type Room struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	RoomName string
	RoomType string
	User1    User
	User2    User
}
