package refreshtkn

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/usecase"
	"github.com/HynDuf/tasks-go-clean-architecture/pkg/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository repository.UserRepository
}

func NewRefreshTokenUsecase(userRepository repository.UserRepository) usecase.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(email string) (entity.User, error) {
	return rtu.userRepository.GetByID(email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
