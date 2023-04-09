package usecase

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
)

type SignupUsecase interface {
	Create(user *entity.User) error
	GetUserByEmail(email string) (entity.User, error)
	CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error)
}
