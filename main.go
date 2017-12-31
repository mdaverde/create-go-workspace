package main

import (
	"os"

	"github.com/codegangsta/cli"
	"fmt"
)

var GlobalFlags []cli.Flag
var Commands []cli.Command

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Marlon L."
	app.Email = "milanlandaverde@gmail.com"
	app.Usage = "Generates the directory structure for a go workspace"
	app.Action = func (c *cli.Context) error {
		if c.NArg() == 1 && c.NumFlags() == 0{
			args := c.Args()
			return createWorkspace(args[0], &createWorkspaceOptions{
				Silent: false,
				DirEnv: true,
				MainGo: true,
			})
		}
		return nil
	}
	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = func (c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %s", err)
	}
}
