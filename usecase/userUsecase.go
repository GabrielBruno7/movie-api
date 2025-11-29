package usecase

import (
	"crud/domain/user"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository user.Repository
}

func NewUserUsecase(repository user.Repository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (userUsecase *UserUsecase) CreateUser(name string, email string, password string) (string, error) {
	user := user.NewUser(
		uuid.New().String(),
		name,
		email,
		password,
	)

	err := userUsecase.checkIfUserAlreadyExists(user)
	if err != nil {
		return "", err
	}

	err = userUsecase.hashPassword(user)
	if err != nil {
		return "", err
	}

	err = userUsecase.repository.Create(user)

	return user.Id, err
}

func (userUsecase *UserUsecase) checkIfUserAlreadyExists(user *user.User) error {
	user, err := userUsecase.repository.LoadUserByEmail(user)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("O usuário com o e-mail " + user.Email + " já existe")
	}

	return nil
}

func (userUsecase *UserUsecase) hashPassword(user *user.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedBytes)

	return nil
}
