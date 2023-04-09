package tasks

import (
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/usecase"
)

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) usecase.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

func (tu *taskUsecase) Create(task *entity.Task) error {
	return tu.taskRepository.Create(task)
}

func (tu *taskUsecase) ToggleStatus(task *entity.Task) error {
	return tu.taskRepository.ToggleStatus(task)
}

func (tu *taskUsecase) FetchByUserID(userID int) ([]entity.Task, error) {
	return tu.taskRepository.FetchByUserID(userID)
}

func (tu *taskUsecase) FetchByUserIDAndStatus(userID int, completed bool) ([]entity.Task, error) {
	return tu.taskRepository.FetchByUserIDAndStatus(userID, completed)
}
