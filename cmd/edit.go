package cmd

import (
	"os"
	"os/exec"
	"piotrek813/jebudu/cmd/validation"
	"piotrek813/jebudu/config"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var edit *cli.Command = &cli.Command{
	Name:      "edit",
	ArgsUsage: "[jebud]",
	Before:    validation.ArgsPresent,
	Action: func(ctx *cli.Context) error {
		editor := config.Gc.Editor
		j := jebud.Get(ctx.Args().First())

		cmd := exec.Command(editor, j.Path)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		return err
	},
}
