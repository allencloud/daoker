package cli

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"../docker"
	"github.com/codegangsta/cli"
)

// listContainers lists all container in the filesystem
func listContainers(c *cli.Context) {
	containers, err := docker.Containers()
	if err != nil {
		log.Fatal(err)
	}

	for index, con := range containers {
		fmt.Println(index)
		fmt.Printf("Name:%s \n", con.Name)
		fmt.Printf("ID:  %s \n", con.ID)
		fmt.Printf("Status:%s \n", con.State.StateString())
		fmt.Printf("Image:%s \n", con.Config.Image)
		fmt.Printf("Command:%s \n", con.Config.Cmd)
		fmt.Println("-------------------------------------\n\n")
	}
}
