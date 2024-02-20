package cmd

import (
	"fmt"
	"piotrek813/jebudu/cmd/validation"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var get *cli.Command = &cli.Command{
	Name:      "get",
	Args:      true,
	ArgsUsage: "[jebud]",
	Before:    validation.ArgsPresent,
	Action: func(ctx *cli.Context) error {
		n := ctx.Args().Get(0)
		fmt.Printf("j: %v\n", jebud.Get(n))

		return nil
	},
}
