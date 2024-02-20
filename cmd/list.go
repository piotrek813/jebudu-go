package cmd

import (
	"fmt"
	"piotrek813/jebudu/jebud"

	"github.com/urfave/cli/v2"
)

var list *cli.Command = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Action: func(ctx *cli.Context) error {
		jebuds, err := jebud.GetAll()
		if err != nil {
			return cli.Exit(err, 1)
		}

		if len(jebuds) == 0 {
			fmt.Println("No jebuds found")
			return nil
		}

		for _, j := range jebuds {
			fmt.Println(j)
		}

		return nil
	},
}
