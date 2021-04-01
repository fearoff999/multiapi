package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fearoff999/multiapi/service/DockerComposeService"
	"github.com/urfave/cli/v2"
)

func main() {
	DockerComposeService.Write()
	return
	app := (&cli.App{
		Name:  "multiapi",
		Usage: "Swagger-UI nginx+docker-compose wrapper",
		Action: func(c *cli.Context) error {
			fmt.Println("Not implemented yet")
			return nil
		},
	})

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
