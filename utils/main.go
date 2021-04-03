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
