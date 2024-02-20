package cmd

import (
	"fmt"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var get *cli.Command = &cli.Command{
	Name:      "get",
	Args:      true,
	ArgsUsage: "[jebud]",
	Before: func(ctx *cli.Context) error {
		if !ctx.Args().Present() {
			n := ctx.App.Name + " " + ctx.Command.FullName()
			fmt.Printf("\"%v\" requires 1 or more arguments\n", n)
			cli.ShowSubcommandHelpAndExit(ctx, 1)
		}

		return nil
	},
	Action: func(ctx *cli.Context) error {
		n := ctx.Args().Get(0)
		fmt.Printf("j: %v\n", jebud.Get(n))

		return nil
	},
}
