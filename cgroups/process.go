package cgroups

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"errors"
	"path/filepath"
	"strings"
)

// ErrPidInNoContainer is exported
var ErrPidInNoContainer = errors.New("No container has your input pid.")

// ContainsPid returns container ID if given pid in it,
// and returns an error if not found or got an error
func ContainsPid(pid int)(string, error){
	cgroupRoot := GetCgroupsRoot()

	// FIXME:
	// if a container or docker use --cgroup-parent parameter,
	// then the following does not work unfortunately
	dockerSubsystemPath := filepath.Join(cgroupRoot, "memory", "docker")

	// traverse all dirs which are all container_id
	dirs, _ := ioutil.ReadDir(dockerSubsystemPath)

	for _, dir := range dirs{
		if dir.IsDir() == false{
			continue
		}

		found, err := findPid(dockerSubsystemPath, dir.Name(),pid)
		if err != nil {
			continue
		}

		// container dir.Name() contains pid
		if found == true {
			return dir.Name(), nil
		}else{
			continue
		}
	}
	return "", fmt.Errorf("No container contains pid %d", pid)
}

// ContainerPidNum returns pid number given ID container contains,
// and returns an error if not found or got an error
func ContainerPidNum(containerID string)(int, error){
	cgroupRoot := GetCgroupsRoot()

	procsPath := filepath.Join(cgroupRoot, "memory", "docker", containerID, "cgroups.procs")

	pids, err := getPidsInContainer(procsPath)
	if err != nil {
		return -1, fmt.Errorf("Failed to get pid array (%v)", err)
	}

	return len(pids), nil
}


// findPid returns true, if containerID has this pid
func findPid(path, containerID string, pid int)(bool, error){
	procsPath := filepath.Join(path, containerID, "cgroups.procs")

	pids, err := getPidsInContainer(procsPath)
	if err != nil {
		return false, fmt.Errorf("Failed to get pid array (%v)", err)
	}

	// traverse
	for _, value := range pids {
		if value == strconv.Itoa(pid){
			return true, nil
		}
	}

	return false, fmt.Errorf("Container %s has no pid %d", containerID, pid)
}

// getPidsContainer gets an array of all pids in a container 
func getPidsInContainer(path string)([]string, error){
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Read the data in file cgroup.procs
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, fmt.Errorf("No process in container.")
	}

	dataStr := string(data)
	dataStr = dataStr[:len(dataStr)-1]
	pids := strings.Split(dataStr, "\n")

	return pids, nil
}