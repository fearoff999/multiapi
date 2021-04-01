package NginxLocationService

import (
	"errors"

	"github.com/fearoff999/multiapi/utils"
)

func GenerateLocationOutput(serviceName string, port string) string {
	if serviceName == "" || port == "" {
		panic(errors.New("empty input arguments"))
	}

	const tpl = `
location /{{.ServiceName}} {
	proxy_pass http://{{.ServiceName}}-swagger:{{.Port}};
}`

	return utils.ReplaceTpl(tpl, struct {
		ServiceName string
		Port        string
	}{
		serviceName,
		port,
	})
}
