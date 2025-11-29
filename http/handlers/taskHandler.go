package handlers

import (
	"crud/http/dto"
	"crud/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	usecase *usecase.TaskUsecase
}

func NewTaskHandler(usecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		usecase: usecase,
	}
}

func (h *TaskHandler) List(context *gin.Context) {
	tasks, err := h.usecase.ListTasks()
	if err != nil {
		context.JSON(500, gin.H{"error": "Erro ao listar tasks"})
		return
	}

	context.JSON(200, tasks)
}

func (h *TaskHandler) Create(context *gin.Context) {
	var request dto.CreateTaskRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	id, err := h.usecase.CreateTask(request.Title, request.Description)
	if err != nil {
		context.JSON(500, gin.H{"error": "Erro ao criar task"})
		return
	}

	context.JSON(201, gin.H{"id": id})
}
