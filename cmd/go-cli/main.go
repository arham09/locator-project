package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"

	"go-cli/internal/handler"
	"go-cli/internal/model"
)

func main() {
	var language string

	app := &cli.App{
		Name:    "Sinbad Utility",
		Usage:   "To help Sinbad sails the seven seas much faster",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "migrate the model db",
				Action: func(c *cli.Context) error {
					model.InitialMigration()
					fmt.Println("Model has been migrated")
					return nil
				},
			},
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "options for task Project",
				Subcommands: []*cli.Command{
					{
						Name:  "open",
						Usage: "open a Listed Project",
						Action: func(c *cli.Context) error {
							handler.GetHandler(c.Args().First())
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a new Project",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())

							name := c.Args().First()

							handler.AddHandler(name)

							fmt.Println("new Project:", c.Args().First())
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "See Project list",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())
							handler.ReadHandler()
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing Project",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())
							fmt.Println("removed Project:", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
