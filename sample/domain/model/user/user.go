package user

type UserID string

type User interface {
	ID() UserID
	Name() string
}

type user struct {
	id   UserID
	name string
}

func NewUser(id UserID, name string) User {
	return user{
		id:   id,
		name: name,
	}
}

func (u user) ID() UserID {
	return u.id
}

func (u user) Name() string {
	return u.name
}
