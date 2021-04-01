package EnvService

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

func replaceExtension(s string) string {
	r := regexp.MustCompile(`\.[^\.]+$`)
	return r.ReplaceAllLiteralString(s, "")
}

func buildFilePathsString(filePaths []string) string {
	elements := []string{}
	tpl := "{ name: \"{{.Name}}\", url: \"/api/{{.FileName}}\" }"
	t := template.Must(template.New("").Parse(tpl))
	buf := &bytes.Buffer{}

	for _, val := range filePaths {
		t.Execute(buf, struct {
			Name     string
			FileName string
		}{
			replaceExtension(val),
			val,
		})
		elements = append(elements, buf.String())
		buf.Reset()
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

func GetEnvVariableString(serviceName string, filePaths []string) string {
	filePathsString := buildFilePathsString(filePaths)
	tpl := "{{.ServiceName}}_URLS={{.FilePathsString}}"
	t := template.Must(template.New("").Parse(tpl))
	buf := &bytes.Buffer{}
	t.Execute(buf, struct {
		ServiceName     string
		FilePathsString string
	}{
		strings.ToUpper(serviceName),
		filePathsString,
	})
	return buf.String()
}
