package userrepo

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *entity.User) error {
	return nil
}

func (ur *userRepository) GetByEmail(email string) (entity.User, error) {
	return entity.User{}, nil
}

func (ur *userRepository) GetByID(id string) (entity.User, error) {
	return entity.User{}, nil
}
