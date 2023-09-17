package user

import (
	"testing"

	"github.com/onnga-wasabi/go-ddd/sample/application/user/register"
	"github.com/onnga-wasabi/go-ddd/sample/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestNewUserApplicationService(t *testing.T) {
	type args struct {
		userRepository user.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserApplicationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUserApplicationService(tt.args.userRepository))
		})
	}
}

func Test_userApplicationService_Register(t *testing.T) {
	type args struct {
		command register.UserRegisterCommand
	}
	tests := []struct {
		name           string
		generateObject func() userApplicationService
		args           args
		want           register.UserRegisterResult
	}{
		{
			name: "登録成功",
			generateObject: func() userApplicationService {
				userRepoMock := user.NewMockUserRepository(t)
				userRepoMock.EXPECT().
					Save(user.NewUser("id", "test_name")).
					Return(nil)
				return userApplicationService{
					userRepository: userRepoMock,
				}
			},
			args: args{
				command: register.UserRegisterCommand{
					Name: "test_name",
				},
			},
			want: register.UserRegisterResult{
				CreatedUserID: user.UserID("id"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.generateObject()
			assert.Equal(t, tt.want, s.Register(tt.args.command))
		})
	}
}
