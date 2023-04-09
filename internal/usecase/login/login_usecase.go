package login

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/usecase"
	"github.com/HynDuf/tasks-go-clean-architecture/pkg/tokenutil"
)

type loginUsecase struct {
	userRepository repository.UserRepository
}

func NewLoginUsecase(userRepository repository.UserRepository) usecase.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}
func (lu *loginUsecase) GetUserByEmail(email string) (entity.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
