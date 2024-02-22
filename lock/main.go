package lock

import (
	"os"
	"path/filepath"
	"piotrek813/jebudu/config"
)

func ensureExists() error {
	p := filepath.Join(config.Gc.DotsPath, "jebudu-lock.json")
	_, err := os.Stat(p)
	os.IsNotExist(err)

	if !os.IsNotExist(err) {
		return nil
	}

	f, err := os.Create(p)
	f.Close()

	return err
}

func put() {
	ensureExists()

}
