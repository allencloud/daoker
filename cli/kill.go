package cli

import (
	"fmt"
	"os"
	"syscall"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// killContainer stops a container forcefully by executing `kill -9`
func killContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Kill command takes exact one argument. See '%s kill --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.GetContainer(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	pid := container.State.Pid

	if pid <= 0 {
		fmt.Printf("Container (%s) is not running\n", container.ID)
		os.Exit(0)
	}

	// FIXME:check whether the container exists

	// kill -9 pid
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		fmt.Printf("Failed to kill %d (%v)", pid, err)
		os.Exit(1)
	}

	fmt.Printf("Container %s has been killed successfully", container.ID)
}
