package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFiles(srcFolder string, destFolder string) error {
	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		return err
	}
	for _, f := range files {
		err := copy(filepath.Join(srcFolder, f.Name()), filepath.Join(destFolder, f.Name()))
		if err != nil {
			return fmt.Errorf("could not copy content of folder %v to %v", srcFolder, destFolder)
		}
	}
	return nil
}

func copy(sourcePath string, destinationPath string) error {
	dir, _ := os.Getwd()

	source, err := ioutil.ReadFile(filepath.Join(dir, sourcePath))
	if err != nil {
		return errors.New("could not read source file")
	}

	err = ioutil.WriteFile(filepath.Join(dir, destinationPath), source, 0755)
	if err != nil {
		return errors.New("could not read destination file")
	}
	return nil
}
