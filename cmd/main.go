package cmd

import (
	"fmt"
	"log"
	"os"
	dot "piotrek813/dot/t"

	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{Name: "dot", ArgsUsage: "[package]"}

	(*app).Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Action: func(ctx *cli.Context) error {
				dots, err := dot.GetAll()
				if err != nil {
					return cli.Exit(err, 1)
				}

				fmt.Printf("dots: %v\n", dots)

				return nil
			},
		},
		{
			Name: "get",
			Action: func(ctx *cli.Context) error {

				if !ctx.Args().Present() {
					return cli.Exit("What ?", 1)
				}
				n := ctx.Args().Get(0)
				fmt.Printf("dots: %v\n", dot.Get(n))

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
