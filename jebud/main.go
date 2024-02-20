package jebud

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"piotrek813/jebudu/config"
	"slices"

	"github.com/BurntSushi/toml"
)

type Jebud struct {
	// TODO i think it should get added automatically to toml
	// TODO todos not showing up in telescoe move to google todo
	Name         string
	Path         string
	Dependencies []string
}

func (j *Jebud) Install() {
	j.installDependencies()
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func (j *Jebud) installDependencies() error {
	ds := ""

	for _, d := range j.Dependencies {
		if commandExists(d) {
			fmt.Printf("Skipping %v, already installed\n", d)
			continue
		}

		ds += d + " "
	}

	if ds == "" {
		return nil
	}

	// idk why this is the right way but okay
	cmd := exec.Command("bash", "-c", "sudo apt install "+ds)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
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
