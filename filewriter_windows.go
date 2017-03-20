// +build windows

package filewriter

import (
	"os"
	"syscall"
	"unsafe"
)

func makeInheritSa() *syscall.SecurityAttributes {
	var sa syscall.SecurityAttributes
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 1
	return &sa
}

func createFile(name string, perm os.FileMode) (*os.File, error) {
	path, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	securityAttributes := makeInheritSa()

	fd, err := syscall.CreateFile(path,
		syscall.FILE_APPEND_DATA,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_DELETE,
		securityAttributes,
		syscall.OPEN_ALWAYS,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0)
	if err != nil {
		return nil, err
	}

	file := os.NewFile(uintptr(fd), name)

	return file, nil
}
