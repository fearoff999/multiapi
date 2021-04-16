package InitializeController

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fearoff999/multiapi/service/DockerComposeService"
	"github.com/fearoff999/multiapi/service/EnvService"
	"github.com/fearoff999/multiapi/service/FileService"
	"github.com/fearoff999/multiapi/service/FrontendService"
	"github.com/fearoff999/multiapi/service/InspectDirectoryService"
	"github.com/fearoff999/multiapi/service/NginxLocationService"
	"github.com/fearoff999/multiapi/utils"
)

func scanDirs() ([]string, map[string][]string) {
	currentDirectory, _ := os.Getwd()

	files := map[string][]string{}
	directories := []string{}

	exclusions := []string{"nginx_config", "basic_auth"}
	scannedDirectories := InspectDirectoryService.GetDirectories(currentDirectory, exclusions)
	for _, dirName := range scannedDirectories {
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

func writeNginxConfig(name string, port string, protected bool) {
	basicAuthString := "auth_basic off;"
	if protected {
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

	auth_basic "Restricted API";
	auth_basic_user_file /etc/nginx/basic_auth/.root;

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
		Restart: "always",
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
			"./html/:/usr/share/nginx/html/",
			"./nginx_config/:/etc/nginx/conf.d/",
			"./basic_auth/:/etc/nginx/basic_auth/",
		},
		DependsOn: servicesNames,
		Ports: []string{
			"80:80",
		},
		Restart: "always",
	}

	return services
}

func protectRoot() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input root user: ")
	user, _ := reader.ReadString('\n')
	fmt.Print("Input root pass: ")
	password, _ := reader.ReadString('\n')
	writeBasicAuthFile("root", strings.ReplaceAll(user, "\n", ""), strings.ReplaceAll(password, "\n", ""))
}

func Initialize() {
	_, err := os.Stat("./basic_auth/.root")
	if err != nil {
		protectRoot()
	}

	rootHtpasswdContent, err := ioutil.ReadFile("./basic_auth/.root")
	if err != nil {
		panic(err)
	}

	dirs, files := scanDirs()
	services := map[string]DockerComposeService.Service{}
	htmlConfig := map[string]bool{}
	defaultPort := "8080"
	envString := ""
	for _, dir := range dirs {
		services = addSwaggerService(services, dir, defaultPort)
		envString += EnvService.GetEnvVariableString(dir, files[dir]) + "\n"
		isProtected := false
		htpasswdPath := "./" + dir + "/.htpasswd"
		info, err := os.Stat(htpasswdPath)
		if err == nil && !info.IsDir() {
			isProtected = true
			bacontent, _ := ioutil.ReadFile(htpasswdPath)
			bastring := string(bacontent)
			removeBasicAuthFile(dir)
			file, _ := os.Create("./basic_auth/." + dir)
			defer file.Close()
			file.WriteString(bastring + string(rootHtpasswdContent))
		}
		writeNginxConfig(dir, defaultPort, isProtected)
		htmlConfig[dir] = isProtected
	}
	writeDefaultNginxConfig(dirs)
	htmlOut := FrontendService.GenerateHtml(htmlConfig)
	FileService.Write("./html/", "index.html", htmlOut)
	services = addNginxService(services)
	FileService.Write("./", ".env", envString)
	dc := DockerComposeService.BuildDockerCompose(services)
	FileService.Write("./", "docker-compose.yaml", DockerComposeService.MarshalDockerCompose(dc))
	fmt.Println("Configuration initialized successfuly")
}

func CleanUp() {
	os.RemoveAll("nginx_config")
	os.RemoveAll("html")
	os.Remove("docker-compose.yaml")
	os.Remove(".env")
	fmt.Println("Configuration cleanuped successfuly")
}

func removeBasicAuthFile(dirName string) {
	os.Remove("./basic_auth/." + dirName)
}

func removeBasicAuthFileForService(dirName string) {
	os.Remove("./" + dirName + "/.htpasswd")
}

func Unprotect(dirName string) {
	removeBasicAuthFileForService(dirName)
	removeBasicAuthFile(dirName)
}

func writeBasicAuthFile(dirName string, user string, password string) {
	FileService.AssertDir("./basic_auth/")
	exec.Command("htpasswd", "-mbc", "./basic_auth/."+dirName, user, password).Output()
}

func writeBasicAuthFileForService(dirName string, user string, password string) {
	FileService.AssertDir("./" + dirName + "/")
	exec.Command("htpasswd", "-mbc", "./"+dirName+"/.htpasswd", user, password).Output()
}

func Protect(dirName string, user string, password string) {
	writeBasicAuthFileForService(dirName, user, password)
}
