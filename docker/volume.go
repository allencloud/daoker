package docker

import (
	"fmt"
	"path/filepath"

	"../utils"
	"github.com/docker/go-units"
)

// MountPoint is the intersection point between a volume and a container. It
// specifies which volume is to be used and where inside a container it should
// be mounted.
type MountPoint struct {
	Source      string // Container host directory
	Destination string // Inside the container
	RW          bool   // True if writable
	Name        string // Name set by user
	Driver      string // Volume driver to use

	//Volume      Volume `json:"-"`

	// Note Mode is not used on Windows
	Mode string `json:"Relabel"` // Originally field was `Relabel`"

	// Note Propagation is not used on Windows
	Propagation string // Mount propagation string
	Named       bool   // specifies if the mountpoint was specified by name

	// Specifies if data should be copied from the container before the first mount
	// Use a pointer here so we can tell if the user set this value explicitly
	// This allows us to error out when the user explicitly enabled copy but we can't copy due to the volume being populated
	CopyData bool `json:"-"`
}

// ContainerVolumeUsage gets the disk usage and inode usage of a volume
// only internal volumes are considered like "VOLUME /var/lib/mysql" in Dockerfile
// we will skip "-v a:b" volumes
func ContainerVolumeUsage(mp MountPoint) (string, int64, error) {
	mountPath, err := getHostMountPath(mp.Name)
	if err != nil {
		return "", 0, err
	}

	diskUsageSize := utils.GetDirDiskSpace(mountPath)

	// use go-units to translate it into human readable size such as "100MB"
	diskUsageSizeStr := units.HumanSize(diskUsageSize)

	inodeUsageSize := utils.GetDirInodes(mountPath)

	return diskUsageSizeStr, inodeUsageSize, nil
}

// getHostMountPath gets absolute host path of volume
func getHostMountPath(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("not an internal volume")
	}

	volumePath := GetVolumeRoot()

	// FIXME:
	// for docker 1.10.0 path is like below
	// while in old version of docker engine, not like this
	hostMountPath := filepath.Join(volumePath, "name", "_data")

	return hostMountPath, nil
}
