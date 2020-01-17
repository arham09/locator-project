package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
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
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					fmt.Println("added task:", c.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					fmt.Println("completed task:", c.Args().First())
					return nil
				},
			},
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "options for task Project",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new Project",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())
							fmt.Println("new Project:", c.Args().First())
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "See Project list",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())
							fmt.Println("new Project:", c.Args().First())
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
