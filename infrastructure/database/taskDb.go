package database

import (
	"crud/domain/task"
	"database/sql"
)

type TaskDb struct {
	db *sql.DB
}

func NewTaskDb(db *sql.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (persistence *TaskDb) Create(t *task.Task) error {
    query := `
        INSERT INTO tasks (
            title,
            description,
            is_completed
        ) VALUES ($1, $2, $3)
        RETURNING id
    `

    err := persistence.db.QueryRow(query, t.Title, t.Description, t.IsCompleted).Scan(&t.ID)
    if err != nil {
        return err
    }

    return nil
}

func (persistence *TaskDb) FindAll() ([]task.Task, error) {
	rows, err := persistence.db.Query("SELECT id, title, description, is_completed FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []task.Task

	for rows.Next() {
		var task task.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
