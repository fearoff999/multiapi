package InitializeController

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fearoff999/multiapi/service/DockerComposeService"
	"github.com/fearoff999/multiapi/service/EnvService"
	"github.com/fearoff999/multiapi/service/FileService"
	"github.com/fearoff999/multiapi/service/InspectDirectoryService"
	"github.com/fearoff999/multiapi/service/NginxLocationService"
	"github.com/fearoff999/multiapi/utils"
)

func scanDirs() ([]string, map[string][]string) {
	currentDirectory, _ := os.Getwd()

	files := map[string][]string{}
	directories := []string{}

	scannedDirectories := InspectDirectoryService.GetDirectories(currentDirectory)
	for _, dirName := range scannedDirectories {
		if dirName == "nginx_config" {
			continue
		}

		currentFiles := InspectDirectoryService.GetFiles(dirName, []string{"yaml", "yml"})
		if len(currentFiles) == 0 {
			continue
		}

		directories = append(directories, dirName)
		files[dirName] = []string{}

		for _, fileName := range currentFiles {
			files[dirName] = append(files[dirName], "./"+dirName+"/"+fileName)
		}
	}

	return directories, files
}

func writeNginxConfig(name string, port string) {
	info, err := os.Stat("./basic_auth/." + name)
	basicAuthString := ""
	if err == nil && !info.IsDir() {
		basicAuthTpl := `
	auth_basic "Restricted API";
	auth_basic_user_file /etc/nginx/basic_auth/.{{.Name}};
`
		basicAuthString = utils.ReplaceTpl(basicAuthTpl, struct {
			Name string
		}{
			Name: name,
		})
	}
	out := NginxLocationService.GenerateLocationOutput(name, port, basicAuthString)
	FileService.Write("./nginx_config/swagger/", name+".conf", out)
}

func writeDefaultNginxConfig(names []string) {
	out := `
server {
	listen	80;
	listen	[::]:80;
	server_name localhost;

	location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

	include /etc/nginx/conf.d/swagger/*.conf;
}
`
	FileService.Write("./nginx_config/", "default.conf", out)
}

func addSwaggerService(services map[string]DockerComposeService.Service, name string, port string) map[string]DockerComposeService.Service {
	services[name+"-swagger"] = DockerComposeService.Service{
		Image: "swaggerapi/swagger-ui",
		Environment: map[string]string{
			"URLS": "${" + strings.ToUpper(name) + "_URLS}",
		},
		Volumes: []string{
			"./" + name + ":/usr/share/nginx/html/api",
		},
	}

	return services
}

func addNginxService(services map[string]DockerComposeService.Service) map[string]DockerComposeService.Service {
	servicesNames := []string{}

	for key := range services {
		servicesNames = append(servicesNames, key)
	}

	services["nginx"] = DockerComposeService.Service{
		Image: "nginx:latest",
		Volumes: []string{
			"./nginx_config/:/etc/nginx/conf.d/",
			"./basic_auth/:/etc/nginx/basic_auth/",
		},
		DependsOn: servicesNames,
		Ports: []string{
			"80:80",
		},
	}

	return services
}

func Initialize() {
	dirs, files := scanDirs()
	services := map[string]DockerComposeService.Service{}
	defaultPort := "8080"
	envString := ""
	for _, dir := range dirs {
		// port := fmt.Sprint(defaultPort + i)
		services = addSwaggerService(services, dir, defaultPort)
		envString += EnvService.GetEnvVariableString(dir, files[dir]) + "\n"
		writeNginxConfig(dir, defaultPort)
	}
	writeDefaultNginxConfig(dirs)
	services = addNginxService(services)
	FileService.Write("./", ".env", envString)
	dc := DockerComposeService.BuildDockerCompose(services)
	FileService.Write("./", "docker-compose.yaml", DockerComposeService.MarshalDockerCompose(dc))
	fmt.Println("Configuration initialized successfuly")
}

func CleanUp() {
	os.RemoveAll("nginx_config")
	os.Remove("docker-compose.yaml")
	os.Remove(".env")
	fmt.Println("Configuration cleanuped successfuly")
}

func removeBasicAuthFile(dirName string) {
	os.Remove("./basic_auth/." + dirName)
}

func Unprotect(dirName string) {
	removeBasicAuthFile(dirName)
}

func writeBasicAuthFile(dirName string, user string, password string) {
	FileService.AssertDir("./basic_auth/")
	_, err := exec.Command("htpasswd", "-mbc", "./basic_auth/."+dirName, user, password).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func Protect(dirName string, user string, password string) {
	writeBasicAuthFile(dirName, user, password)
}
