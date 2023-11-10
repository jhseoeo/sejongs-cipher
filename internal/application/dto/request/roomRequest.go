package request

import "github.com/google/uuid"

type RoomCreateRequest struct {
	RoomName string `json:"roomName"`
}

type RoomJoinRequest struct {
	RoomId uuid.UUID `json:"roomId"`
}

type RoomLeaveRequest struct {
	RoomId uuid.UUID `json:"roomId"`
}
