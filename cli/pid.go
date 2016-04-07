package cli

import (
	"os"
	"fmt"
	"strconv"

	"../cgroups"
	//"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// pidContainer returns container name if it contains the given pid
// otherwise, returns "No container contains pid"
func pidContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Pid command takes exact one argument. See '%s pid --help'.", c.App.Name)
	}

	pidNumStr := c.Args()[0]

	pidNum, err := strconv.Atoi(pidNumStr)
	if err != nil {
		log.Fatalf("Pid command's argument should be positive, invalid input (%s)", pidNumStr)
		os.Exit(1)
	}

	containerID, err := cgroups.ContainsPid(pidNum)
	if err != nil {
		if err == cgroups.ErrPidInNoContainer {
			fmt.Print(err.Error())
			os.Exit(0)
		}else{
			log.Fatalf("Failed to find pid in cgroups filesystem (%v)", err)
			os.Exit(1)
		}
	}

	// get a correct container ID
	fmt.Printf("Pid(%d) is in container(%s)\n", pidNum, containerID)
}
