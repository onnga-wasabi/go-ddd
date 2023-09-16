package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func UserSandboxCommand() *cli.Command {
	return &cli.Command{
		Name:  "boom",
		Usage: "test cli v2",
		Action: func(ctx *cli.Context) error {
			fmt.Println("boom! I say!")
			fmt.Printf("Hello %q", ctx.Args().Get(0))
			fmt.Println()
			return nil
		},
	}
}
