package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{Name: "dot", ArgsUsage: "[package]"}

	(*app).Commands = []*cli.Command{list, get}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
