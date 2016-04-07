package cli

import (
	//"os"
	"fmt"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// pidContainers shows container name if it contains the given pid
func pidContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Pid command takes exact one argument. See '%s pid --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(container.State.Pid)
}
