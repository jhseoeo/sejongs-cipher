package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/gofiber/contrib/websocket"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	wsmessage "github.com/jhseoeo/khuthon2023/internal/application/dto/wsMessage"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
	"github.com/jhseoeo/khuthon2023/internal/domain/entities"
	"github.com/jhseoeo/khuthon2023/internal/infrastructure/rest"
)

// GameUsecase is a usecase for game
type GameUsecase struct {
	restClient   *rest.Client
	scoreService port.ScoreServicePort
	gameService  port.GameServicePort
}

// NewGameUsecase creates a new game usecase
func NewGameUsecase(
	scoreService port.ScoreServicePort,
	gameService port.GameServicePort,
) *GameUsecase {
	return &GameUsecase{
		restClient:   rest.NewClient(),
		scoreService: scoreService,
		gameService:  gameService,
	}
}

// Join joins a game
func (g *GameUsecase) Join(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	if message.UserType == "tetris" {
		err := g.gameService.JoinTetrisUser(sessionId, conn)
		if err != nil {
			return &wsmessage.WriteMessage{
				Status: 500,
				Ok:     false,
				Type:   "join",
				Data:   fmt.Sprintf("failed to join tetris user: %+v", err),
			}, errors.Errorf("failed to join tetris user: %w", err)
		}
	} else if message.UserType == "wordguess" {
		err := g.gameService.JoinWordGuessUser(sessionId, conn)
		if err != nil {
			return &wsmessage.WriteMessage{
				Status: 500,
				Ok:     false,
				Type:   "join",
				Data:   fmt.Sprintf("failed to join wordguess user: %+v", err),
			}, errors.Errorf("failed to join wordguess user: %w", err)
		}
	} else {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "join",
			Data:   fmt.Sprintf("invalid user type: %s", message.UserType),
		}, errors.Errorf("invalid user type: %s", message.UserType)
	}

	return &wsmessage.WriteMessage{
		Status: 200,
		Ok:     true,
		Type:   "join",
	}, nil
}

// Start starts a game
func (g *GameUsecase) Start(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	return &wsmessage.WriteMessage{}, nil
}

// Leave leaves a game
func (g *GameUsecase) Leave(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	err := g.gameService.LeaveUser(sessionId, conn)
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "leave",
			Data:   fmt.Sprintf("failed to leave user: %+v", err),
		}, errors.Errorf("failed to leave user: %w", err)
	}
	return &wsmessage.WriteMessage{
		Status: 200,
		Ok:     true,
		Type:   "leave",
	}, nil
}

func (g *GameUsecase) Relay(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	err := g.gameService.RelayMessage(sessionId, conn, entities.GameMessage{
		Type: message.Type,
		Data: message.Data,
	})
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "relay",
			Data:   fmt.Sprintf("failed to relay message: %+v", err),
		}, errors.Errorf("failed to relay message: %w", err)
	}

	return &wsmessage.WriteMessage{}, nil
}

// // End ends a game
// func (g *GameUsecase) End(ctx context.Context, req request.GameEndRequest, userId uuid.UUID) (response.BaseResponse[*response.Empty], error) {
// 	// get user
// 	user, err := g.userService.Get(ctx, userId)
// 	if err != nil {
// 		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
// 			errors.Errorf("failed to get user: %w", err)
// 	}

// 	// get room and check if user is allowed to end game
// 	room, err := g.roomService.Get(ctx, req.RoomId)
// 	if err != nil {
// 		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
// 			errors.Errorf("failed to get room: %w", err)
// 	}

// 	// end game
// 	err = g.roomService.EndGame(ctx, req.RoomId, user)
// 	if err != nil {
// 		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
// 			errors.Errorf("failed to end game: %w", err)
// 	}

// 	// create score
// 	err = g.scoreService.CreateScore(ctx, &room.User1, room.User2, req.Score)
// 	if err != nil {
// 		return response.NewErrorResponse[*response.Empty](500, "Internal Server Error"),
// 			errors.Errorf("failed to create score: %w", err)
// 	}
// 	return response.NewEmptyBaseResponse[*response.Empty](), nil
// }

// VerifyWord verifies a word
func (g *GameUsecase) VerifyWord(
	ctx context.Context,
	conn *websocket.Conn,
	sessionId string,
	userId string,
	message wsmessage.ReadMessage,
) (*wsmessage.WriteMessage, error) {
	url := "https://ko.dict.naver.com/api3/koko/search?query=" + message.Word + "&m=pc&range=word&page=1"
	respBytes, err := g.restClient.Get(ctx, url, "")
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "verify",
			Data:   fmt.Sprintf("failed to get response: %+v", err),
		}, errors.Errorf("failed to get response: %w", err)
	}

	var respMap map[string]interface{}
	err = json.Unmarshal(respBytes, &respMap)
	if err != nil {
		return &wsmessage.WriteMessage{
			Status: 500,
			Ok:     false,
			Type:   "verify",
		}, errors.Errorf("failed to get response: %w", err)
	}

	searchResultMap := respMap["searchResultMap"].(map[string]interface{})
	searchResultList := searchResultMap["searchResultListMap"].(map[string]interface{})
	searchResult := searchResultList["WORD"].(map[string]interface{})
	searchResultItems := searchResult["items"].([]interface{})
	if len(searchResultItems) == 0 {
		return &wsmessage.WriteMessage{
			Status:    200,
			Ok:        true,
			Type:      "verify",
			IsCorrect: false,
		}, nil
	}

	// compare word with first item
	item := searchResultItems[0].(map[string]interface{})
	itemWord := item["handleEntry"].(string)
	if itemWord != message.Word {
		return &wsmessage.WriteMessage{
			Status:    200,
			Ok:        true,
			Type:      "verify",
			IsCorrect: false,
		}, nil
	}

	// get word's meaning
	itemMeaning := item["meansCollector"].([]interface{})[0].(map[string]interface{})
	means := itemMeaning["means"].([]interface{})[0].(map[string]interface{})["value"].(string)
	return &wsmessage.WriteMessage{
		Status:      200,
		Ok:          true,
		Type:        "verify",
		IsCorrect:   true,
		Description: means,
	}, nil
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
