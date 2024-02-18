package dot

import (
	"fmt"
	"os"
	"path"
	"piotrek813/dot/config"
	"slices"

	"github.com/BurntSushi/toml"
)

type Dot struct {
	Name string
}

func (d *Dot) Install() {
	fmt.Printf("installing %v", d.Name)
}

func Get(f string) *Dot {
	var d Dot

	if _, err := os.Stat(f); err != nil {
		f = path.Join(config.Gc.DotsPath, f, "config.toml")
	}
	_, err := toml.DecodeFile(f, &d)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return nil
	}

	return &d
}

func GetAll() ([]string, error) {
	ignore := []string{".git"}
	d := config.Gc.DotsPath
	entries, err := os.ReadDir(d)
	if err != nil {
		return []string{}, err
	}

	dots := []string{}
	for _, e := range entries {
		if !e.Type().IsDir() {
			continue
		}
		n := e.Name()

		if slices.Contains[[]string](ignore, n) {
			continue
		}

		dots = append(dots, e.Name())
	}

	return dots, nil
}
