//go:build windows

package checkrunning

import (
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
)

func createMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func IsRunning(app_name string) (bool, error) {
	_, err := createMutex(app_name)
	if err != nil {
		if err.Error() == "Cannot create a file when that file already exists." {
			return true, nil
		}
		return false, err
	} else {
		return false, nil
	}
}
