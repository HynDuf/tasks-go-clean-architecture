package userrepo

import (
	"time"

	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
)

type UserModel struct {
	ID             int        `gorm:"primary_key;column:id"`
	Name           string     `gorm:"column:name"`
	Email          string     `gorm:"column:email"`
	HashedPassword string     `gorm:"column:hashed_password"`
	CreatedAt      *time.Time `gorm:"column:created"`
}

func (UserModel) FromUserEntity(u entity.User) *UserModel {
	return &UserModel{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
        HashedPassword: u.HashedPassword,
	}
}

func (u *UserModel) ToUserEntity() *entity.User {
	return &entity.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
        HashedPassword: u.HashedPassword,
	}
}
