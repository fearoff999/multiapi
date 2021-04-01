package DockerComposeService

import "testing"

func getDockerCompose() dockerCompose {
	return BuildDockerCompose(map[string]service{
		"nginx": {
			Build: build{
				Context:    ".",
				Dockerfile: "./Dockerfile.test",
			},
			DependsOn: []string{"swagger1", "swagger2"},
			Volumes:   []string{"/var/log:/var/log"},
			Environment: map[string]string{
				"URLS": "[{url: \"blabla\", file: \"blabla.yaml\"}]",
			},
		},
	})
}

func TestBuildDockerCompose(t *testing.T) {
	dc := getDockerCompose()
	if dc.Version != "3.9" || dc.Services["nginx"].Build.Dockerfile != "./Dockerfile.test" {
		t.Error("Docker-compose structure is not valid")
	}
}

func TestMarshalDockerCompose(t *testing.T) {
	res := MarshalDockerCompose(getDockerCompose())
	expected := `version: "3.9"
services:
  nginx:
    build:
      context: .
      dockerfile: ./Dockerfile.test
    ports: []
    depends_on:
    - swagger1
    - swagger2
    volumes:
    - /var/log:/var/log
    environment:
      URLS: '[{url: "blabla", file: "blabla.yaml"}]'
    command: ""
`
	if res != expected {
		t.Errorf("Docker-compose yaml is not valid, got\n%v want\n%v", res, expected)
	}
}
