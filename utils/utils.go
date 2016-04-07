package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareDockerVersion compares current docker version with specified docker version
// curVersion: 1.10.0
// comVersion: 1.2.0
// return true
func CompareDockerVersion(curVersion, comVersion string) (bool, error) {
	current := strings.SplitN(curVersion, ".", 3)
	if len(current) != 3 {
		return false, fmt.Errorf("Complete docker version has 3 numbers (%s)", curVersion)
	}

	compare := strings.SplitN(comVersion, ".", 3)
	if len(compare) != 3 {
		return false, fmt.Errorf("Complete docker version has 3 numbers (%s)", comVersion)
	}

	for i := 0; i < 3; i++ {
		a, err := strconv.Atoi(current[i])
		if err != nil {
			return false, fmt.Errorf("version should be an integer (%s)", current[i])
		}

		b, err := strconv.Atoi(compare[i])
		if err != nil {
			return false, fmt.Errorf("version should be an integer (%s)", compare[i])
		}

		if a > b {
			return true, nil
		} else if a == b {
			continue
		} else {
			return false, nil
		}
	}

	// version, major, minor are all the same
	return true, nil
}
