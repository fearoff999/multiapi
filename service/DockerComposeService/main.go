package DockerComposeService

import (
	"gopkg.in/yaml.v2"
)

type dockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]service `yaml:"services"`
}

type build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type service struct {
	Build       build             `yaml:"build"`
	Ports       []string          `yaml:"ports"`
	DependsOn   []string          `yaml:"depends_on"`
	Volumes     []string          `yaml:"volumes"`
	Environment map[string]string `yaml:"environment"`
	Command     string            `yaml:"command"`
}

func BuildDockerCompose(services map[string]service) dockerCompose {
	return dockerCompose{
		Version:  "3.9",
		Services: services,
	}
}

func MarshalDockerCompose(dockerComposeStruct dockerCompose) string {
	d, _ := yaml.Marshal(&dockerComposeStruct)
	return string(d)
}
