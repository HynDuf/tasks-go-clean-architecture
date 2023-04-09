package repository

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email string) (entity.User, error)
	GetByID(id string) (entity.User, error)
}
