package config

import (
	"fmt"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	Scope    string
	AppImage string `toml:"app_image"`
	Home     string
	User     *user.User
}

func (gc *GlobalConfig) setUser() {
	u, err := user.Current()

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	gc.User = u
}

func New(f string) *GlobalConfig {
	var gc GlobalConfig

	_, err := toml.DecodeFile("config.toml", &gc)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return nil
	}

	gc.setUser()

	return &gc
}
