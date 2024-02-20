package validation

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ArgsPresent(ctx *cli.Context) error {
	if !ctx.Args().Present() {
		n := ctx.App.Name + " " + ctx.Command.FullName()
		fmt.Printf("\"%v\" requires 1 or more arguments\n", n)
		cli.ShowSubcommandHelpAndExit(ctx, 1)
	}

	return nil
}
