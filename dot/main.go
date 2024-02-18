package dot

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Dot struct {
	name string
}

func (d *Dot) Install() {
	fmt.Printf("installing %v", d.name)
}

func New(f string) *Dot {
	var d Dot

	_, err := toml.DecodeFile(f, &d)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return nil
	}

	return &d
}
