package repository

import (
	"errors"

	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
)
var ErrUserNotFound = errors.New("Video not found")
type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email string) (entity.User, error)
	GetByID(id string) (entity.User, error)
}
