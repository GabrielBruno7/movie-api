package database

import (
	"crud/debug"
	"crud/domain/group"
	"database/sql"
)

type GroupDb struct {
	db *sql.DB
}

func NewGroupDb(db *sql.DB) *GroupDb {
	return &GroupDb{db: db}
}

func (persistence *GroupDb) CreateGroup(group *group.Group) error {
	query := `
	    INSERT INTO groups (
			id,
			name,
			owner_id
	    ) VALUES ($1, $2, $3)
	    RETURNING id
	`

	_, err := persistence.db.Exec(query,
		group.ID,
		group.Name,
		group.User.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (persistence *GroupDb) CreateMemberGroup(group *group.Group) error {
	query := `
	    INSERT INTO group_members (
			group_id,
			user_id
	    ) VALUES ($1, $2)
	    RETURNING id
	`

	_, err := persistence.db.Exec(query,
		group.ID,
		group.User.Id,
	)

	if err != nil {
		debug.DD(err)
		return err
	}

	return nil
}

func (g *GroupDb) LoadGroupByName(group *group.Group) (*group.Group, error) {
	query := `
		SELECT id, name, owner_id
		FROM groups
		WHERE name = $1
		AND owner_id = $2
		LIMIT 1
	`

	err := g.db.QueryRow(query, group.Name, group.User.Id).Scan(
		&group.ID,
		&group.Name,
		&group.User.Id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return group, nil
}
