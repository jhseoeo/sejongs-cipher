package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jhseoeo/khuthon2023/internal/application/dto/response"
	"github.com/jhseoeo/khuthon2023/internal/application/port"
)

type AuthService struct {
	userService port.UserServicePort
	jwtSecret   string
}

func NewAuthService(userService port.UserServicePort) *AuthService {
	return &AuthService{
		userService: userService,
		jwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

func (s *AuthService) Login(ctx context.Context, userId string, password string) (response.BaseResponse[*response.AuthLoginResponse], error) {
	// check user exists and password is correct
	user, err := s.userService.GetByUserId(ctx, userId)
	if err != nil {
		return response.NewErrorResponse[*response.AuthLoginResponse](401, "wrong userId or password"), nil
	}
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
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
