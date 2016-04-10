package cli

import (
	"github.com/codegangsta/cli"
)

var (
	flLogAppend = cli.BoolFlag{
		Name:  "append",
		Usage: "Append specific logs to container.",
	}

	flLogSize = cli.BoolFlag{
		Name:  "size",
		Usage: "Show the size of container's log file.",
	}
)
