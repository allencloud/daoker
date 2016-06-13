package cli

import (
	"fmt"
	"strconv"
	"os"

	"../cgroups"
	"../docker"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// pidExcdContainer returns containers whose process number exceeds a specified number
func pidExcContainer(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalf("Pidexc command takes exact one argument. See '%s pidexcd --help'.", c.App.Name)
	}

	excdNumStr := c.Args()[0]

	excdNum, err := strconv.Atoi(excdNumStr)
	if err != nil {
		log.Fatalf("Pidexc command's argument should be positive, invalid input (%s)", excdNumStr)
		os.Exit(1)
	}

	containers, err := docker.Containers()
	if err != nil {
		log.Fatal(err)
	}

	var exist = false
	for _, container := range containers {

		ID := container.ID

		num, err := cgroups.ContainerPidNum(ID)
		if err != nil {
			log.Fatal(err.Error())
		}

		if num > excdNum {
			fmt.Printf("Container(%s)'s process number is %d, which exceeds %d\n", ID, num, excdNum)
			exist = true
		}
	}
	if !exist {
		fmt.Printf("There are no containers whose process number exceeds %d\n", excdNum)
	}

}
