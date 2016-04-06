package cli

import (
	"os"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// pidContainers lists all container in the filesystem
func pidContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Pid command takes exact one argument. See '%s pid --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.Container(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

}
