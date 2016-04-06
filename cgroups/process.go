package cgroups

import (
	"os"
	"path/filepath"

	"../utils"
)

// MainPid searches memory.oom_control of this container
// oom_kill_disable 0
// under_oom 0
// if under_oom equals to 1, return true.
func MainPid(ID string) (int, error) {
	oomFilepath := filepath.Join(GetCgroupsRoot(), "memory", "docker", ID, "memory.oom_control")

}
