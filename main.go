package main

import (
	"fmt"
	"piotrek813/dot/config"
)

// type Config struct {
// 	Name string
// 	Dependencies []string
// 	Scope []string
// }

func main() {
	gc := config.New("config.toml")

	fmt.Printf("gc: %v\n", gc)

}
