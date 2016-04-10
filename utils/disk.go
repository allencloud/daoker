package utils

import (
	"os"
	"path/filepath"
)

// GetDirDiskSpace returns disk space a directory including all files takes
func GetDirDiskSpace(path string) float64 {
	size := float64(0)
	filepath.Walk(path, func(_ string, file os.FileInfo, _ error) error {
		size += float64(file.Size())
		return nil
	})
	return size
}

// GetDirInodes returns inodes consumed by a directory
// FIXME: it is not accurate
func GetDirInodes(path string) int64 {
	inodeNum := int64(0)
	filepath.Walk(path, func(_ string, _ os.FileInfo, _ error) error {
		inodeNum++
		return nil
	})
	return inodeNum
}
