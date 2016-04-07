package cli

import (
	"fmt"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// volumeContainer shows all details of a container's volume
func volumeContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Stop command takes exact one argument. See '%s stop --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(container.Config.Volumes)

}
