package cmd

import (
	"piotrek813/jebudu/cmd/validation"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var edit *cli.Command = &cli.Command{
	Name:      "edit",
	ArgsUsage: "[jebud]",
	Before:    validation.ArgsPresent,
	Action: func(ctx *cli.Context) error {
		return jebud.Edit(ctx.Args().First())

	},
}
