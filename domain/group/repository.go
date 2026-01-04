package group

type Repository interface {
	CreateGroup(group *Group) error
	CreateMemberGroup(group *Group) error
	LoadGroupByName(group *Group) (*Group, error)
}
