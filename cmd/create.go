package cmd

import (
	"piotrek813/jebudu/cmd/validation"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var create *cli.Command = &cli.Command{
	Name:      "create",
	Before:    validation.ArgsPresent,
	ArgsUsage: "[name]",
	Action: func(ctx *cli.Context) error {
		n := ctx.Args().First()
		err := jebud.Create(n)

		if err != nil {
			return err
		}

		jebud.Edit(n)

		return nil
	},
}
