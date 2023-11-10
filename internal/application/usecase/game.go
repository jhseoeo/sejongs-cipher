package usecase

import (
	"context"
	"encoding/json"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
	"github.com/jhseoeo/khuthon2023/internal/infrastructure/rest"
)

// GameUsecase is a usecase for game
type GameUsecase struct {
	restClient   *rest.Client
	userService  port.UserServicePort
	roomService  port.RoomServicePort
	scoreService port.ScoreServicePort
}

// NewGameUsecase creates a new game usecase
func NewGameUsecase(
	userService port.UserServicePort,
	roomService port.RoomServicePort,
	scoreService port.ScoreServicePort,
) *GameUsecase {
	return &GameUsecase{
		restClient:   rest.NewClient(),
		userService:  userService,
		roomService:  roomService,
		scoreService: scoreService,
	}
}

// Start starts a game
func (g *GameUsecase) Start(ctx context.Context, req request.GameStartRequest, userId uuid.UUID) (response.BaseResponse[*response.Empty], error) {
	// get user
	user, err := g.userService.Get(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to get user: %w", err)
	}

	// start game
	err = g.roomService.StartGame(ctx, req.RoomId, user)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to start game: %w", err)
	}
	return response.NewEmptyBaseResponse[*response.Empty](), nil
}

// End ends a game
func (g *GameUsecase) End(ctx context.Context, req request.GameEndRequest, userId uuid.UUID) (response.BaseResponse[*response.Empty], error) {
	// get user
	user, err := g.userService.Get(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to get user: %w", err)
	}

	// get room and check if user is allowed to end game
	room, err := g.roomService.Get(ctx, req.RoomId)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to get room: %w", err)
	}

	// end game
	err = g.roomService.EndGame(ctx, req.RoomId, user)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to end game: %w", err)
	}

	// create score
	err = g.scoreService.CreateScore(ctx, &room.User1, room.User2, req.Score)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
			errors.Errorf("failed to create score: %w", err)
	}
	return response.NewEmptyBaseResponse[*response.Empty](), nil
}

// GetRanks gets ranks
func (g *GameUsecase) GetRanks(ctx context.Context, req request.GetRanksRequest) (response.BaseResponse[*response.GetRanksResponse], error) {
	// get scores
	scores, err := g.scoreService.GetScoreList(ctx, req.Page)
	if err != nil {
		return response.NewErrorResponse[*response.GetRanksResponse](500, "Internal Server Error"),
			errors.Errorf("failed to get scores: %w", err)
	}
	return response.NewGetRanksResponse(scores), nil
}

// VerifyWord verifies a word
func (g *GameUsecase) VerifyWord(ctx context.Context, req request.GameTestWordRequest) (response.BaseResponse[*response.GameTestWordResponse], error) {
	url := "https://ko.dict.naver.com/api3/koko/search?query=" + req.Word + "&m=pc&range=word&page=1"
	respBytes, err := g.restClient.Get(ctx, url, "")
	if err != nil {
		return response.NewErrorResponse[*response.GameTestWordResponse](500, "Internal Server Error"),
			errors.Errorf("failed to get response: %w", err)
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return response.NewErrorResponse[*response.GameTestWordResponse](500, "Internal Server Error"),
			errors.Errorf("failed to unmarshal response: %w", err)
	}
	searchResultMap := respMap["searchResultMap"].(map[string]interface{})
	searchResultList := searchResultMap["searchResultListMap"].(map[string]interface{})
	searchResult := searchResultList["WORD"].(map[string]interface{})
	searchResultItems := searchResult["items"].([]interface{})
	if len(searchResultItems) == 0 {
		return response.NewGameTestWordResponse(false, ""), nil
	}

	// compare word with first item
	item := searchResultItems[0].(map[string]interface{})
	itemWord := item["handleEntry"].(string)
	if itemWord != req.Word {
		return response.NewGameTestWordResponse(false, ""), nil
	}

	// get word's meaning
	itemMeaning := item["meansCollector"].([]interface{})[0].(map[string]interface{})
	means := itemMeaning["means"].([]interface{})[0].(map[string]interface{})["value"].(string)
	return response.NewGameTestWordResponse(true, means), nil
}
