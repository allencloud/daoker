package cli

import (
	"fmt"

	"../cgroups"
	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// oomContainer returns true if a container is under oom.
func oomContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Oom command takes exact one argument. See '%s oom --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	ID := container.ID

	isOom, err := cgroups.CheckContainerOOM(ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	if isOom == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
