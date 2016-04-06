package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:   "ps",
			Usage:  "List containers",
			Action: listContainers,
		},
	}
)
