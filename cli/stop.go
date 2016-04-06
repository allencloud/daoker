package cli

import (
	"fmt"
	"os"
	"syscall"

	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// stopContainer stops a container by executing `kill`
func killContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Stop command takes exact one argument. See '%s stop --help'.", c.App.Name)
	}

	IDOrName := c.Args()[0]
	container, err := docker.Container(IDOrName)
	if err != nil {
		log.Fatal(err.Error())
	}

	pid := container.State.Pid

	if pid <= 0 {
		fmt.Printf("Container (%s) is not running\n", container.ID)
		os.Exit(0)
	}

	// FIXME:check whether the container exists

	// kill pid
	// FIXME: signal is not SIGKILL
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		fmt.Printf("Failed to stop %d (%v)", pid, err)
		os.Exit(1)
	}

	os.Exit(0)

}
