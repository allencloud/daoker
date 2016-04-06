package docker

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Container struct {
	Id    string `json:"ID"`
	State State  `json:"State"`
}

// Containers returns an array of docker containers unmarshaled from config.json
// in docker root path.
func Containers() ([]Container, error) {
	containersPath := filepath.Join(utils.GetDockerRoot(), "containers")

	containerEntries, err := ioutil.ReadDir(containersPath)
	if err != nil {
		return nil, err
	}

	containers := []Container{}

	for _, entry := range containerEntries {
		entryName := entry.Name()
		if len(entryName) != len("ffb082df6289394f4d285ef2ea31051deed699f6b352cf4109fb7e97fd15237a") {
			continue
		}

		configFilename, err := getConfigFilename()
		if err != nil {
			return nil, err
		}

		containerJsonPath := filepath.Join(containersPath, entryName, configFilename)

		con, err := containerFromJson(containerJsonPath)
		if err != nil {
			continue
		}

		containers = append(containers, con)
	}
	return containers, nil
}

// Container returns a container by given IDOrName
func Container(IDOrName) (container, error) {

	return
}

// getConfigFilename gets container's config file name by docker version
// If docker version differs, json file's name differs, too.
// config.json in 1.10.0-, while config.v2.json in 1.10.0+
func getConfigFilename() (string, error) {
	var configFilename string

	match, err := util.CompareDockerVersion(utils.GetDockerVersion(), "1.10.0")
	if err != nil {
		return "", err
	}
	// If match is true, it means current docker version is newer or at least equal.
	if match {
		configFilename = "config.v2.json"
	} else {
		configFilename = "config.json"
	}
	return configFilename, nil
}

// containerFromJson unmarshals json file into a container instance
func containerFromJson(file string) (Container, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return Container{}, err
	}

	var con Container
	if err := json.Unmarshal(data, &con); err != nil {
		return Container{}, err
	}

	return con, nil
}
