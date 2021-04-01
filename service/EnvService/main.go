package EnvService

import (
	"regexp"
	"strings"

	"github.com/fearoff999/multiapi/utils"
)

func replaceExtension(s string) string {
	r := regexp.MustCompile(`\.[^\.]+$`)
	return r.ReplaceAllLiteralString(s, "")
}

func buildFilePathsString(filePaths []string) string {
	elements := []string{}
	tpl := "{ name: \"{{.Name}}\", url: \"/api/{{.FileName}}\" }"

	for _, val := range filePaths {
		res := utils.ReplaceTpl(tpl, struct {
			Name     string
			FileName string
		}{
			replaceExtension(val),
			val,
		})
		elements = append(elements, res)
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

func GetEnvVariableString(serviceName string, filePaths []string) string {
	filePathsString := buildFilePathsString(filePaths)
	tpl := "{{.ServiceName}}_URLS={{.FilePathsString}}"
	return utils.ReplaceTpl(tpl, struct {
		ServiceName     string
		FilePathsString string
	}{
		strings.ToUpper(serviceName),
		filePathsString,
	})
}
