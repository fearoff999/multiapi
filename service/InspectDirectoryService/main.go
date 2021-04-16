package InspectDirectoryService

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/fearoff999/multiapi/utils"
)

func GetDirectories(directoryPath string, exclusions []string) []string {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}

	res := []string{}
	for _, f := range files {
		if f.IsDir() && !utils.Contains(f.Name(), exclusions) {
			res = append(res, f.Name())
		}
	}

	return res
}

func isMatchExtension(fileName string, extensions []string) bool {
	expectedExtensions := strings.Join(extensions, "|")
	r := regexp.MustCompile(`\.(` + expectedExtensions + `)$`)
	return r.MatchString(fileName)
}

func GetFiles(directoryPath string, extensions []string) []string {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		panic(err)
	}

	res := []string{}
	for _, f := range files {
		if !f.IsDir() && isMatchExtension(f.Name(), extensions) {
			res = append(res, f.Name())
		}
	}

	return res
}
