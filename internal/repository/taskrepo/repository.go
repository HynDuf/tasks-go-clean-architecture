package taskrepo

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"gorm.io/gorm"
)

type taskRepository struct {
	db         *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Create(task *entity.Task) error {
	return nil
}

func (tr *taskRepository) ToggleStatus(task *entity.Task) error {
	return nil
}
func (tr *taskRepository) FetchByUserID(userID int) ([]entity.Task, error) {
	return nil, nil
}

func (tr *taskRepository) FetchByUserIDAndStatus(userID int, completed bool) ([]entity.Task, error) {
	return nil, nil
}
