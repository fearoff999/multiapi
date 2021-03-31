package service

import (
	"bytes"
	"errors"
	"os"
	"text/template"
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

func assertDir(path string) {
	if exists := dirExists(path); !exists {
		err := os.Mkdir(path, 0775)
		if err != nil {
			panic(err)
		}
	}
}

func write(serviceName string, output string) {
	if serviceName == "" || output == "" {
		panic(errors.New("Empty input arguments"))
	}

	path := "../nginx_confs/"
	assertDir(path)
	os.Remove(path + serviceName + ".conf")
	file, _ := os.Create(path + serviceName + ".conf")
	defer file.Close()
	file.WriteString(output)
}

func generateLocationOutput(serviceName string, port string) string {
	if serviceName == "" || port == "" {
		panic(errors.New("Empty input arguments"))
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
