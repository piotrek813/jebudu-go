package main

import (
	"piotrek813/jebudu/cmd"
	"piotrek813/jebudu/config"
)

func main() {
	config.New()

	cmd.Run()
}
