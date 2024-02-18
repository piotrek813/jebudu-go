package main

import (
	"piotrek813/dot/cmd"
	"piotrek813/dot/config"
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
