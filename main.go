package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
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
