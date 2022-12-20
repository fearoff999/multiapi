package RunController

import (
	"fmt"
	"log"
	"os/exec"
)

func Run() {
	_, err := exec.Command("/bin/bash", "-c", "docker-compose --env-file .env up -d").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration started successfuly")
}

func Stop() {
	_, err := exec.Command("/bin/bash", "-c", "docker-compose down --remove-orphans").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Configuration stoped successfuly")
}
