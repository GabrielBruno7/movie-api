package invite

import (
	"crud/domain/user"
)

type Invite struct {
	ID       string
	Sender   *user.User
	Receiver *user.User
}
