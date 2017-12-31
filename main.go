package main

import (
	"os"

	"github.com/codegangsta/cli"
	"fmt"
	"github.com/pkg/errors"
)

var Commands []cli.Command

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Marlon L."
	app.Email = "milanlandaverde@gmail.com"
	app.Usage = "Generates the directory structure for a go workspace"

	var silent bool
	var dirEnv bool
	var mainGo bool

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "silent",
			Destination: &silent,
		},
		cli.BoolTFlag{
			Name: "dir-env",
			Destination: &dirEnv,
		},
		cli.BoolTFlag{
			Name: "main-go",
			Destination: &mainGo,
		},
	}

	app.Action = func (c *cli.Context) error {
		if numArgs := c.NArg(); numArgs > 0 {
			args := c.Args()
			return createWorkspace(args[numArgs - 1], &createWorkspaceOptions{
				Silent: silent,
				DirEnv: dirEnv,
				MainGo: mainGo,
			})
		} else if numArgs < 1 {
			return errors.New("Provide a project name (i.e. github.com/mdaverde/great-idea)")
		}
		return nil
	}

	app.Commands = Commands
	app.CommandNotFound = func (c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %s", err)
	}
}
