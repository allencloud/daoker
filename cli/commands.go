package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:   "ps",
			Usage:  "List containers",
			Action: listContainers,
		},
		{
			Name:   "kill",
			Usage:  "Stop a container by force",
			Action: killContainer,
		},
		{
			Name:   "stop",
			Usage:  "Stop a container",
			Action: stopContainer,
		},
		{
			Name:   "volume",
			Usage:  "Show all details of a container's volumes",
			Action: volumeContainer,
		},
		{
			Name:   "pid",
			Usage:  "Print container name if it contains the given pid",
			Action: pidContainer,
		},
		{
			Name:   "oom",
			Usage:  "Return true if a container is under oom",
			Action: oomContainer,
		},
		{
			Name:   "pidnum",
			Usage:  "Print process numbers in your specified container",
			Action: pidNumContainer,
		},
		{
			Name:   "logs",
			Usage:  "Add specific details into a container's log",
			Flags: []cli.Flag{
				flLogAppend,flLogSize,
			},
			Action: logContainer,
		},
	}
)
