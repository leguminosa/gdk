package gomodfinder

import (
	"errors"
	"os"
	"path/filepath"
)

func Find() (string, error) {
	currentPath, err := filepath.Abs(".")
	if err != nil {
		return "", errors.New("cannot find the absolute path of the current directory: " + err.Error())
	}

	return recursivelyFind(currentPath, 0)
}

const maxRecursiveCount = 100

func recursivelyFind(currentPath string, recursiveCount int) (string, error) {
	if recursiveCount > maxRecursiveCount {
		return "", errors.New("cannot find the go.mod file after nesting more than 100 levels deep")
	}

	if currentPath == "/" {
		return "", errors.New("cannot find the go.mod file in any of the parent directories all the way to the root")
	}

	files, err := os.ReadDir(currentPath)
	if err != nil {
		return "", errors.New("cannot list the files in the current directory (" + currentPath + "): " + err.Error())
	}

	for _, file := range files {
		if file.Name() == "go.mod" {
			return currentPath, nil
		}
	}

	var parentPath string
	parentPath, err = filepath.Abs(currentPath + "/../")
	if err != nil {
		return "", errors.New("cannot find the absolute path of the parent directory (" + currentPath + "/../): " + err.Error())
	}

	if parentPath == currentPath {
		return "", errors.New("stuck at the same directory (" + parentPath + ")")
	}

	recursiveCount++
	return recursivelyFind(parentPath, recursiveCount)
}
