package usecase

import "github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"

type RefreshTokenUsecase interface {
	GetUserByID(id string) (entity.User, error)
	CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
