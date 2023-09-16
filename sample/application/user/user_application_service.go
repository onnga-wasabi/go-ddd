package user

type UserID string

type User interface {
	Name() string
}

type user struct {
	ID UserID
}

func NewUser() User {
	return user{
		ID: UserID("id"),
	}
}

func (u user) Name() string {
	return "Name"
}
