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

	fmt.Printf("Container %s has %d volumes", container.ID, len(container.MountPoints))

	i := 1
	for _, mount := range container.MountPoints {
		fmt.Printf(" Volume %d :", i)
		i++

		fmt.Printf("  └ Host src Path: %s\n")
		fmt.Printf("  └ Container Dst Path: %s\n")

		diskUsage, inodeUsage, err := docker.ContainerVolumeUsage(mount)
		if err != nil {
			log.Errorf("Failed to get usage of volume %s: %v", mount.Name, err)
		}

		fmt.Printf("  └ Volume disk Usage: %s\n", diskUsage)
		fmt.Printf("  └ Volume inode Usage: %d\n\n", inodeUsage)
	}
}
