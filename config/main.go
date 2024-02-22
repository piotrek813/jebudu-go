package config

import (
	"fmt"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	Scope       string
	AppImage    string `toml:"app_image"`
	DotsPath    string `toml:"dots_path"`
	ScriptsPath string `toml:"scripts_path"`
	User        *user.User
	Editor      string
}

func (gc *GlobalConfig) setUser() {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}

	gc.User = u
}

var Gc *GlobalConfig

func New() *GlobalConfig {
	if Gc == nil {
		Gc = &GlobalConfig{}

		_, err := toml.DecodeFile("config.toml", Gc)

		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			return nil
		}

		Gc.setUser()

		Gc.Dump()
	}

	return Gc
}

func (gc *GlobalConfig) Dump() {
	toml.NewEncoder(os.Stdout).Encode(Gc)
}
