package response

import "github.com/jhseoeo/khuthon2023/internal/domain/entities"

type GetRanksResponse struct {
	Scores []entities.Score `json:"scores"`
}

func NewGetRanksResponse(scores []entities.Score) BaseResponse[*GetRanksResponse] {
	return BaseResponse[*GetRanksResponse]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data: &GetRanksResponse{
			Scores: scores,
		},
	}
}

type GameTestWordResponse struct {
	IsCorrect   bool   `json:"isCorrect"`
	Description string `json:"description"`
}

func NewGameTestWordResponse(isCorrect bool, description string) BaseResponse[*GameTestWordResponse] {
	return BaseResponse[*GameTestWordResponse]{
		Ok:      true,
		Status:  200,
		Message: "success",
		Data: &GameTestWordResponse{
			IsCorrect:   isCorrect,
			Description: description,
		},
	}
}
