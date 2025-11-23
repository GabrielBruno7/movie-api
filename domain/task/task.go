package task

type Task struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsCompleted bool    `json:"is_completed"`
}

func NewTask(title string, description *string, isCompleted bool) *Task {
	return &Task{
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
	}
}
