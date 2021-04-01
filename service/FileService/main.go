package FileService

import (
	"errors"
	"os"
)

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}

func AssertDir(path string) {
	if exists := dirExists(path); !exists {
		err := os.Mkdir(path, 0775)
		if err != nil {
			panic(err)
		}
	}
}

func Write(dirName string, fileName string, output string) {
	if dirName == "" || fileName == "" || output == "" {
		panic(errors.New("empty input arguments"))
	}

	AssertDir(dirName)
	os.Remove(dirName + fileName)
	file, _ := os.Create(dirName + fileName)
	defer file.Close()
	file.WriteString(output)
}
