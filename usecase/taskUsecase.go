package usecase

import (
	"crud/domain/task"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repository task.Repository
}

func NewTaskUsecase(repository task.Repository) *TaskUsecase {
	return &TaskUsecase{
		repository: repository,
	}
}

func (taskUsecase *TaskUsecase) CreateTask(title string, description *string) (string, error) {
	task := task.NewTask(
		uuid.New().String(),
		title,
		description,
		false,
	)

	err := taskUsecase.repository.Create(task)

	return task.Id, err
}

func (taskUsecase *TaskUsecase) ListTasks() ([]task.Task, error) {
	return taskUsecase.repository.FindAll()
}
