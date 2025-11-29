package user

type Repository interface {
	Create(user *User) error
	LoadUserByEmail(user *User) (*User, error)
}
