package task

type Repository interface {
	Create(task *Task) error
	FindAll() ([]Task, error)
}
