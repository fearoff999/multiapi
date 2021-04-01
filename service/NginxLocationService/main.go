package NginxLocationService

import (
	"bytes"
	"errors"
	"text/template"
)

func GenerateLocationOutput(serviceName string, port string) string {
	if serviceName == "" || port == "" {
		panic(errors.New("empty input arguments"))
	}

	const tpl = `
location /{{.ServiceName}} {
	proxy_pass http://{{.ServiceName}}-swagger:{{.Port}};
}`
	t := template.Must(template.New("").Parse(tpl))
	buf := &bytes.Buffer{}
	t.Execute(buf, struct {
		ServiceName string
		Port        string
	}{
		serviceName,
		port,
	})
	return buf.String()
}
