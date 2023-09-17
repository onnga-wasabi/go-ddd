package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id   UserID
		name string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewUser(tt.args.id, tt.args.name))
		})
	}
}

func Test_user_ID(t *testing.T) {
	tests := []struct {
		name           string
		generateObject func() user
		want           UserID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.generateObject()
			assert.Equal(t, tt.want, u.ID())
		})
	}
}

func Test_user_Name(t *testing.T) {
	tests := []struct {
		name           string
		generateObject func() user
		want           string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.generateObject()
			assert.Equal(t, tt.want, u.Name())
		})
	}
}
