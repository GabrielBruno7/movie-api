package database

import (
	"crud/domain/user"
	"database/sql"
	"time"
)

type UserDb struct {
	db *sql.DB
}

func NewUserDb(db *sql.DB) *UserDb {
	return &UserDb{db: db}
}

func (persistence *UserDb) Create(user *user.User) error {
	query := `
        INSERT INTO users (
			id,
			name,
            email,
            password,
            created_at
        ) VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	createdAt := time.Now().In(loc).Format(time.RFC3339)

	err := persistence.db.QueryRow(query, user.Id, user.Name, user.Email, user.Password, createdAt).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserDb) LoadUserByEmail(user *user.User) (*user.User, error) {
	query := `
		SELECT id, name, email, password
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	err := p.db.QueryRow(query, user.Email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
