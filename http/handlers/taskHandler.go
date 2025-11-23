package handlers

import (
	"crud/domain/task"
	"crud/http/dto"
	"crud/infrastructure/database"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	persistence task.Repository
}

func NewTaskHandler(databaseConnection *sql.DB) *TaskHandler {
	return &TaskHandler{
		persistence: database.NewTaskDb(databaseConnection),
	}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.persistence.FindAll()
	if err != nil {
		http.Error(w, "Erro ao listar tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateTaskRequest

	json.NewDecoder(r.Body).Decode(&request)

	task := task.NewTask(
		request.Title,
		request.Description,
		false,
	)

	err := h.persistence.Create(task)
	if err != nil {
		http.Error(w, "Erro ao criar task", http.StatusInternalServerError)
		return
	}
    response := map[string]interface{}{
        "message": "Task criada",
        "id":      task.ID,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    json.NewEncoder(w).Encode(response)
}
