package user

type UserRepository interface {
	Save(User) error
}
