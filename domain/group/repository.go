package group

type Repository interface {
	CreateGroup(group *Group) error
	CreateMemberGroup(group *Group) error
	LoadGroupById(group *Group) (*Group, error)
	LoadGroupByName(group *Group) (*Group, error)
	CreateInvite(group *Group) error
}
