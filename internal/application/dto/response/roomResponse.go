package response

import "github.com/jhseoeo/khuthon2023/internal/domain/entities"

type RoomCreateResponse struct {
	entities.Room
}

func NewRoomCreateResponse(room *entities.Room) BaseResponse[*RoomCreateResponse] {
	return BaseResponse[*RoomCreateResponse]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data: &RoomCreateResponse{
			Room: *room,
		},
	}
}

type RoomGetListResponse struct {
	Rooms []entities.Room `json:"rooms"`
}

func NewRoomGetListResponse(rooms []entities.Room) BaseResponse[*RoomGetListResponse] {
	return BaseResponse[*RoomGetListResponse]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data: &RoomGetListResponse{
			Rooms: rooms,
		},
	}
}
