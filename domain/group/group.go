package group

import "crud/domain/user"

type Group struct {
	ID   string
	Name string
	User user.User
}
