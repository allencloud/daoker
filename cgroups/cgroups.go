package cgroups

import (
	"os"
)

const DefaultCgroupsRoot string = "/sys/fs/cgroup"

// GetCgroupsRoot gets cgroups root path in Docker environment
func GetCgroupsRoot() string {
	cgroupsRoot := os.Getenv("CGROUPS_ROOT")
	if cgroupsRoot == "" {
		cgroupsRoot = DefaultCgroupsRoot
	}
	return cgroupsRoot
}
