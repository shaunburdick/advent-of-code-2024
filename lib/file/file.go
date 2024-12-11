package file

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func LoadRelativeFile(path string) (content string, err error) {
	content = ""

	// get the current source filename
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		err = errors.New("Unable to file caller")
	} else {
		absPath, pathErr := filepath.Abs(filename)
		if pathErr != nil {
			err = pathErr
		} else {
			file, fileErr := os.ReadFile(filepath.Join(filepath.Dir(absPath), path))
			if fileErr != nil {
				err = fileErr
			} else {
				content = string(file)
			}
		}
	}

	return content, err
}

func ExistsRelativeFile(path string) bool {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return false
	} else {
		absPath, pathErr := filepath.Abs(filename)
		if pathErr != nil {
			return false
		} else {
			_, fileErr := os.Stat(filepath.Join(filepath.Dir(absPath), path))
			if fileErr != nil {
				return false
			} else {
				return true
			}
		}
	}
}
