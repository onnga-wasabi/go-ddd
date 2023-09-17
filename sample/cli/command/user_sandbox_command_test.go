package command

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/assert"
)

func TestUserSandboxCommand(t *testing.T) {
	tests := []struct {
		name string
		want *cli.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UserSandboxCommand())
		})
	}
}
