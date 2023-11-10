package entities

import (
	"github.com/google/uuid"
)

// Room is a room
type Room struct {
	Id        uuid.UUID
	RoomName  string
	IsPlaying bool
	User1     User
	User2     *User
}
