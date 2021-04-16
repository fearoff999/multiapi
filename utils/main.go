package utils

import (
	"bytes"
	"text/template"
)

func ReplaceTpl(tpl string, replacements interface{}) string {
	t := template.Must(template.New("").Parse(tpl))
	buf := &bytes.Buffer{}
	t.Execute(buf, replacements)
	return buf.String()
}

func Contains(needle string, haystack []string) bool {
	for _, r := range haystack {
		if needle == r {
			return true
		}
	}
	return false
}
