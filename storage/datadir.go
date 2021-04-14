package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDataDirIfNotExists(datadir string) error {
	dir, err := os.Stat(datadir)

	if err != nil {
		if !os.IsExist(err) {
			return createDatadir(datadir)
		}

		return err
	}

	if !dir.IsDir() {
		return fmt.Errorf("datadir cannot be a file")
	}

	return nil
}

func createDatadir(datadir string) error {
	return os.Mkdir(filepath.Clean(datadir), os.ModeDir|os.ModePerm)
}
