package DockerComposeService

import (
	"gopkg.in/yaml.v2"
)

type dockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type Service struct {
	Build       build             `yaml:"build,omitempty"`
	Image       string            `yaml:"image,omitempty"`
	Ports       []string          `yaml:"ports,omitempty"`
	DependsOn   []string          `yaml:"depends_on,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Command     string            `yaml:"command,omitempty"`
	Restart     string            `yaml:"restart,omitempty"`
}

func BuildDockerCompose(services map[string]Service) dockerCompose {
	return dockerCompose{
		Version:  "3.3",
		Services: services,
	}
}

func MarshalDockerCompose(dockerComposeStruct dockerCompose) string {
	d, _ := yaml.Marshal(&dockerComposeStruct)
	return string(d)
}
