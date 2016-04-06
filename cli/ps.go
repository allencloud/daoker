package cli

import (
	//"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/allencloud/daoker/docker"
	"github.com/codegangsta/cli"
)

// listContainers lists all container in the filesystem
func listContainers(c *cli.Context) {
	containers, err := docker.Containers()
	if err != nil {
		log.Fatal(err)
	}
}
