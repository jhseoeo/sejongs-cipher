package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/request"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
)

type AuthUsecase struct {
	userService port.UserServicePort
	jwtSecret   string
}

func NewAuthService(userService port.UserServicePort) *AuthUsecase {
	return &AuthUsecase{
		userService: userService,
		jwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func (s *AuthUsecase) Login(ctx context.Context, req request.AuthLoginRequest) (response.BaseResponse[*response.AuthLoginResponse], error) {
	// check user exists and password is correct
	user, err := s.userService.GetByUserId(ctx, req.UserId)
	if err != nil {
		return response.NewErrorResponse[*response.AuthLoginResponse](401, "wrong userId or password"), nil
	}
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	if user.Password != passwordHash {
		return response.NewErrorResponse[*response.AuthLoginResponse](401, "wrong userId or password"), nil
	}

	// generate jwt token
	claims := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return response.NewErrorResponse[*response.AuthLoginResponse](500, "internal server error"), nil
	}

	return response.NewAuthLoginResponse(t), nil
}

func (s *AuthUsecase) Register(ctx context.Context, req request.AuthRegisterRequest) (response.BaseResponse[*response.Empty], error) {
	// check user exists
	_, err := s.userService.GetByUserId(ctx, req.UserId)
	if err == nil {
		return response.NewErrorResponse[*response.Empty](409, "user already exists"), nil
	}

	// register user
	err = s.userService.Create(ctx, req.UserId, req.Password, req.UserName)
	if err != nil {
		return response.NewErrorResponse[*response.Empty](500, "internal server error"), nil
	}

	return response.NewEmptyBaseResponse[*response.Empty](), nil
}