package group

import (
	"crud/domain/invite"
	"crud/domain/user"
)

type Group struct {
	ID     string
	Name   string
	User   user.User
	Invite *invite.Invite
}
