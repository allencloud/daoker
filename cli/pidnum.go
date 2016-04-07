package cli

import (
	//"os"
	"fmt"

	"../cgroups"
	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// pidNumContainer returns pid number the given ID container contains
func pidNumContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Pid command takes exact one argument. See '%s pid --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]

	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	ID := container.ID

	num, err := cgroups.ContainerPidNum(ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%d processed in container(%s)\n", num, ID)
}
