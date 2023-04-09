package usecase

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
)

type TaskUsecase interface {
	Create(task *entity.Task) error
	ToggleStatus(task *entity.Task) error
	FetchByUserID(userID int) ([]entity.Task, error)
	FetchByUserIDAndStatus(userID int, completed bool) ([]entity.Task, error)
}
