// +build !windows

package filewriter

import (
	"os"
)

func createFile(name string, perm os.FileMode) (*os.File, error) {
	flag := os.O_WRONLY | os.O_CREATE | os.O_APPEND

	return os.OpenFile(name, flag, perm)
}
