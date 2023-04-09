package signup

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/usecase"
	"github.com/HynDuf/tasks-go-clean-architecture/pkg/tokenutil"
)

type signupUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(userRepository repository.UserRepository) usecase.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) Create(user *entity.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByEmail(email string) (entity.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *entity.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *entity.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
