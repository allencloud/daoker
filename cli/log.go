package cli

import (
	"fmt"
	"strings"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// logContainer adds specific details into a container's log
func logContainer(c *cli.Context) {
	if len(c.Args()) != 2 {
		log.Fatalf("addlog command takes more than 1 argument. See '%s addlog --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	ID := container.ID
	logPath := container.LogPath

	arguments := c.Args()[1:]
	content := strings.Join(arguments, " ")

	if err := docker.AddContainerLog(logPath, content); err != nil {
		log.Fatalf("Failed to add log into container %s: %v", ID, err)
	}

	fmt.Printf("Log added to container %s.\n"+
		"You can check by by command 'docker logs %s'", ID, ID)
}
