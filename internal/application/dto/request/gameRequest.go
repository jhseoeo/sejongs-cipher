package request

import "github.com/google/uuid"

type GameStartRequest struct {
	RoomId uuid.UUID `json:"roomId"`
}

type GameEndRequest struct {
	RoomId uuid.UUID `json:"roomId"`
	Score  int       `json:"score"`
}

type GetRanksRequest struct {
	Page int `json:"page"`
}

type GameTestWordRequest struct {
	Word string `json:"word"`
}
