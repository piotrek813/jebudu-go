package main

import (
	"piotrek813/jebudu/cmd"
	"piotrek813/jebudu/config"
)

// type Config struct {
// 	Name string
// 	Dependencies []string
// 	Scope []string
// }

func main() {
	config.New("config.toml")

	cmd.Run()
}
