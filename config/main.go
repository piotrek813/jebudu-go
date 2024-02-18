package config

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

type GlobalConfig struct {
	Scope    string
	AppImage string `toml:"app_image"`
	BasePath string `toml:"base_path"`
	DotsPath string `toml:"dots_path"`
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

func (gc *GlobalConfig) setPath() {
	_, b, _, _ := runtime.Caller(0)
	gc.BasePath = strings.TrimSuffix(b, "config/main.go")
	gc.DotsPath = path.Join(gc.BasePath, "dots")
}

var Gc *GlobalConfig

func New(f string) *GlobalConfig {
	if Gc == nil {
		Gc = &GlobalConfig{}
		Gc.setPath()

		f := path.Join(Gc.BasePath, "config.toml")
		_, err := toml.DecodeFile(f, Gc)

		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			return nil
		}

		Gc.setUser()

		fmt.Printf("Gc: %v\n", Gc)
	}

	return Gc
}
