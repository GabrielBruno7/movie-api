package usecase

import (
	"crud/domain/errs"
	"crud/domain/group"

	"github.com/google/uuid"
)

type GroupUsecase struct {
	repository  group.Repository
	UserUsecase *UserUsecase
}

func NewGroupUsecase(repository group.Repository, userUsecase *UserUsecase) *GroupUsecase {
	return &GroupUsecase{
		repository:  repository,
		UserUsecase: userUsecase,
	}
}

func (groupUsecase *GroupUsecase) CreateGroup(group *group.Group) error {
	_, err := groupUsecase.loadUser(group)
	if err != nil {
		return err
	}

	existingGroup, err := groupUsecase.repository.LoadGroupByName(group)
	if err != nil {
		return err
	}

	if existingGroup != nil {
		return errs.NewWithCode(errs.ErrGroupAlreadyExists, nil)
	}

	group.ID = uuid.New().String()

	groupUsecase.repository.CreateGroup(group)
	groupUsecase.repository.CreateMemberGroup(group)

	return nil
}

func (groupUsecase *GroupUsecase) loadUser(group *group.Group) (*group.Group, error) {
	user, err := groupUsecase.UserUsecase.LoadUserByEmail(&group.User)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errs.NewWithCode(errs.ErrUserNotFound, nil)
	}

	group.User = *user

	return group, nil
}

func (groupUsecase *GroupUsecase) SendInvite(group *group.Group) error {
	sender, err := groupUsecase.UserUsecase.LoadUserByEmail(group.Invite.Sender)
	if err != nil {
		return err
	}

	group.Invite.Sender = sender

	group.User = *group.Invite.Sender

	group, err = groupUsecase.repository.LoadGroupById(group)
	if err != nil {
		return err
	}

	if group == nil {
		return errs.NewWithCode(errs.ErrGroupDoesNotBelongToSender, nil)
	}

	receiver, err := groupUsecase.UserUsecase.LoadUserByEmail(group.Invite.Receiver)
	if err != nil {
		return err
	}

	group.Invite.Receiver = receiver

	group.Invite.ID = uuid.New().String()

	err = groupUsecase.repository.CreateInvite(group)
	if err != nil {
		return err
	}

	return nil
}
