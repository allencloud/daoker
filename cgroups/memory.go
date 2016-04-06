package cgroups

import (
	"os"
	"path/filepath"

	"../utils"
)

// CheckContainerOOM searches memory.oom_control of this container
// oom_kill_disable 0
// under_oom 0
// if under_oom equals to 1, return true.
func CheckContainerOOM(ID string) (bool, error) {
	oomFilepath := filepath.Join(GetCgroupsRoot(), "memory", "docker", ID, "memory.oom_control")

	ch, err := readUnderOomChar(oomFilepath)
	if err != nil {
		return false, err
	}

	switch char {
	case '0':
		return false, nil
	case '1':
		return true, nil
	default:
		return false, fmt.Error("Unexpected char got when checking %s", oomFilepath)
	}
}

// In file memory.oom_control, here is the details.
// We need to fetch the last 0/1 char of under_oom
// oom_kill_disable 1
// under_oom 0
func readUnderOomChar(oomFilepath string) (byte, error) {
	// check if file exists
	_, err := os.Stat(filename)
	if err != nil && os.IsExist(err) {
		return '-', fmt.Error("File %s does not exist. Maybe your container is not running.", oomFilepath)
	}

	// read details from file
	underOomDetails, err := ioutil.ReadFile(oomFilepath)
	if err != nil {
		return '-', fmt.Errorf("Failed to read file %s", oomFilepath)
	}

	length := len(underOomDetails)
	char := underOomDetails[length-2]
	return char, nil
}
