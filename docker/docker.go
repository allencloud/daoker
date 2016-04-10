package docker

import (
	"os"
	"path/filepath"
)

const DefaultDockerRoot string = "/var/lib/docker"
const DefaultDockerVersion string = "1.10.0"

// GetDockerRoot gets docker root path in Docker environment
func GetDockerRoot() string {
	dockerRoot := os.Getenv("DOCKER_ROOT")
	if dockerRoot == "" {
		dockerRoot = DefaultDockerRoot
	}
	return dockerRoot
}

// GetVolumeRoot gets docker volume path in Docker environment
func GetVolumeRoot() string {
	dockerRoot := os.Getenv("DOCKER_ROOT")
	if dockerRoot == "" {
		dockerRoot = filepath.Join(DefaultDockerRoot, "volumes")
	}
	return filepath.Join(dockerRoot, "volumes")
}

// GetDockerVersion gets docker version from env
func GetDockerVersion() string {
	dockerVersion := os.Getenv("DOCKER_VERSION")
	if dockerVersion == "" {
		return DefaultDockerVersion
	}
	// FIXME: validate dockerVersion to be "version.major.minor" format
	return dockerVersion
}
