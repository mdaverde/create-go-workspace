package main

import (
	"os"

	"fmt"
	"github.com/codegangsta/cli"
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
	var readMe bool

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "silent, s",
			Destination: &silent,
			Usage:       "suppress output (default: false)",
		},
		cli.BoolTFlag{
			Name:        "dir-env",
			Destination: &dirEnv,
			Usage:       "generate .envrc (default: true)",
		},
		cli.BoolTFlag{
			Name:        "main-go",
			Destination: &mainGo,
			Usage:       "generate main.go (default: true)",
		},
		cli.BoolTFlag{
			Name:        "read-me",
			Destination: &readMe,
			Usage:       "generate README.md (default: true)",
		},
	}

	app.Action = func(c *cli.Context) error {
		if numArgs := c.NArg(); numArgs > 0 {
			args := c.Args()
			return createWorkspace(args[numArgs-1], &createWorkspaceOptions{
				Silent: silent,
				DirEnv: dirEnv,
				MainGo: mainGo,
				ReadMe: readMe,
			})
		} else if numArgs < 1 {
			return errors.New("Provide a project name (i.e. github.com/mdaverde/great-idea)")
		}
		return nil
	}

	app.Commands = Commands
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %s", err)
	}
}
