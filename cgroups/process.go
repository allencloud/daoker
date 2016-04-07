package cgroups

import ()

// MainPid searches memory.oom_control of this container
// oom_kill_disable 0
// under_oom 0
// if under_oom equals to 1, return true.
func MainPid(ID string) (int, error) {
	return 0, nil
}
