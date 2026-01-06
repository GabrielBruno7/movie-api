package handlers

import (
	"crud/domain/errs"
	"crud/domain/group"
	"crud/domain/invite"
	"crud/domain/user"
	"crud/http/dto"
	"crud/http/response"
	"crud/usecase"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	usecase *usecase.GroupUsecase
}

func NewGroupHandler(usecase *usecase.GroupUsecase) *GroupHandler {
	return &GroupHandler{usecase: usecase}
}

func (groupHandler *GroupHandler) ActionCreateGroup(context *gin.Context) {
	var request dto.CreateGroupRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		response.BadRequest(context, errs.ErrInvalidBody)
		return
	}

	user := user.User{
		Email: context.GetString("email"),
	}

	group := &group.Group{
		Name: request.Name,
		User: user,
	}

	err := groupHandler.usecase.CreateGroup(group)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(202, gin.H{"id": group.ID})
}

func (groupHandler *GroupHandler) ActionSendGroupInvite(context *gin.Context) {
	var request dto.SendInviteRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		response.BadRequest(context, errs.ErrInvalidBody)
		return
	}

	sender := user.User{
		Email: context.GetString("email"),
	}

	receiver := user.User{
		Email: request.Email,
	}

	invite := invite.Invite{
		Sender:   &sender,
		Receiver: &receiver,
	}

	group := group.Group{
		ID:     context.Param("id"),
		Invite: &invite,
	}

	err := groupHandler.usecase.SendInvite(&group)
	if err != nil {
		response.HandleError(context, err)
		return
	}

	context.JSON(202, gin.H{"id": invite.ID})
}
