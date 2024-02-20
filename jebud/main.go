package jebud

import (
	"fmt"
	"os"
	"path"
	"piotrek813/jebudu/config"
	"slices"

	"github.com/BurntSushi/toml"
)

type Jebud struct {
	Name string
	Path string
}

func (j *Jebud) Install() {
	fmt.Printf("installing %v", j.Name)
}

func Get(f string) *Jebud {
	var j Jebud
	p := path.Join(config.Gc.DotsPath, f)

	if _, err := os.Stat(f); err != nil {
		f = path.Join(p, "config.toml")
	}
	_, err := toml.DecodeFile(f, &j)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return nil
	}

	j.Path = p

	return &j
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
