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
