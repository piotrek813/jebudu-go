package jebud

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"piotrek813/jebudu/config"
	"slices"

	"github.com/BurntSushi/toml"
)

type Jebud struct {
	// TODO i think it should get added automatically to toml
	// TODO todos not showing up in telescoe move to google todo
	Name         string
	Path         string
	Dependencies []Dependency
}

func (j *Jebud) Install() {
	j.Dump()
	j.installDependencies()
	fmt.Printf("j.setupScripts(): %v\n", j.setupScripts())
}

func (j *Jebud) setupScripts() error {
	if !j.doesSubmoduleExist("scripts") {
		fmt.Printf("\"dupa\": %v\n", "dupa")
		return nil
	}

	sp := filepath.Join(j.Path, "scripts")

	scripts, err := os.ReadDir(sp)

	if err != nil {
		return err
	}

	for _, s := range scripts {
		op := filepath.Join(sp, s.Name())
		np := filepath.Join(config.Gc.ScriptsPath, s.Name())
		err := os.Symlink(op, np)

		if err != nil {
			return err
		}
	}

	return nil
}
func (j *Jebud) doesSubmoduleExist(n string) bool {
	p := path.Join(j.Path, n)
	_, err := os.Stat(p)

	return !os.IsNotExist(err)
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

func Create(n string) error {
	d := config.Gc.DotsPath

	p := path.Join(d, n)

	if err := os.Mkdir(p, 0700); err != nil {
		return err
	}

	_, err := os.Create(path.Join(p, "config.toml"))

	return err
}

func Edit(n string) error {
	editor := config.Gc.Editor
	j := Get(n)

	cmd := exec.Command(editor, j.Path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err
}

func (j *Jebud) Dump() {
	toml.NewEncoder(os.Stdout).Encode(j)
}
