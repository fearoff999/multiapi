package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fearoff999/multiapi/controller/InitializeController"
	"github.com/fearoff999/multiapi/controller/RunController"
	"github.com/urfave/cli/v2"
)

func main() {
	app := (&cli.App{
		Name:  "multiapi",
		Usage: "Swagger-UI nginx+docker-compose wrapper",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize all required files",
				Action: func(c *cli.Context) error {
					InitializeController.Initialize()
					return nil
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run an initialized setup with docker-compose",
				Action: func(c *cli.Context) error {
					RunController.Run()
					return nil
				},
			},
			{
				Name:    "stop",
				Aliases: []string{"s"},
				Usage:   "Stop configuration with docker-compose",
				Action: func(c *cli.Context) error {
					RunController.Stop()
					return nil
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "Clean all configuration files",
				Action: func(c *cli.Context) error {
					InitializeController.CleanUp()
					return nil
				},
			},
			{
				Name:    "maga",
				Aliases: []string{"m"},
				Usage:   "Make Api Great Again (Stop configuration, CleanUp existing configs, Initialize new one, Run configuration)",
				Action: func(c *cli.Context) error {
					RunController.Stop()
					InitializeController.CleanUp()
					InitializeController.Initialize()
					RunController.Run()
					return nil
				},
			},
			{
				Name:    "protect",
				Aliases: []string{"p"},
				Usage:   "protect #project#",
				Action: func(c *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)

					fmt.Println("")
					fmt.Print("Input project: ")
					dirName, _ := reader.ReadString('\n')
					fmt.Print("Input user: ")
					user, _ := reader.ReadString('\n')
					fmt.Print("Input pass: ")
					password, _ := reader.ReadString('\n')

					InitializeController.Protect(
						strings.ReplaceAll(dirName, "\n", ""),
						strings.ReplaceAll(user, "\n", ""),
						strings.ReplaceAll(password, "\n", ""))

					return nil
				},
			},
			{
				Name:    "unprotect",
				Aliases: []string{"u"},
				Usage:   "uprotect",
				Action: func(c *cli.Context) error {
					reader := bufio.NewReader(os.Stdin)

					fmt.Println("")
					fmt.Print("Input project: ")
					dirName, _ := reader.ReadString('\n')

					InitializeController.Unprotect(strings.ReplaceAll(dirName, "\n", ""))

					return nil
				},
			},
		},
	})

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
