package user

import (
	"github.com/onnga-wasabi/go-ddd/sample/application/user/register"
	"github.com/onnga-wasabi/go-ddd/sample/domain/model/user"
)

type UserApplicationService interface {
	Register(register.UserRegisterCommand) register.UserRegisterResult
}

type userApplicationService struct {
	userRepository user.UserRepository
}

func NewUserApplicationService(
	userRepository user.UserRepository,
) UserApplicationService {
	return userApplicationService{
		userRepository: userRepository,
	}
}

func (s userApplicationService) Register(command register.UserRegisterCommand) register.UserRegisterResult {
	u := user.NewUser(user.UserID("id"), command.Name)
	err := s.userRepository.Save(u)
	if err != nil {
		return register.UserRegisterResult{}
	}
	return register.UserRegisterResult{CreatedUserID: u.ID()}
}
