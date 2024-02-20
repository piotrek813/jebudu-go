package jebud

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Dependency string

func (d Dependency) isCustom(lookup string) bool {
	p := d.getPath(lookup)
	_, err := os.Stat(p)
	return !errors.Is(err, os.ErrNotExist)
}

func (d Dependency) isPresent() bool {
	_, err := exec.LookPath(string(d))
	return err == nil
}

func (d Dependency) getPath(lookup string) string {
	return path.Join(lookup, "dependencies", string(d))
}

func installCustomDepencies(lookup string, ds []Dependency) error {
	for _, d := range ds {
		p := d.getPath(lookup)
		s := fmt.Sprintf("source %v; add", p)
		cmd := exec.Command("bash", "-c", s)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func installWithPackageManager(ds []Dependency) error {
	if len(ds) == 0 {
		return nil
	}
	// idk why this is the right way but okay
	s := ""
	for _, v := range ds {
		s += string(v) + " "
	}
	cmd := exec.Command("bash", "-c", "sudo apt install "+s)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (j *Jebud) installDependencies() error {
	ds := []Dependency{}
	cs := []Dependency{}

	for _, d := range j.Dependencies {
		if d.isPresent() {
			fmt.Printf("Skipping %v, already installed\n", d)
			continue
		}

		if d.isCustom(j.Path) {
			cs = append(cs, d)
			continue
		}

		ds = append(ds, d)
	}

	err := installCustomDepencies(j.Path, cs)

	if err != nil {
		// TODO implement rollbacks
		return err
	}

	err = installWithPackageManager(ds)

	return err
}
