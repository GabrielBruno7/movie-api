package dto

type CreateTaskRequest struct {
	Title       string  `json:"titulo"`
	Description *string `json:"descricao"`
}
