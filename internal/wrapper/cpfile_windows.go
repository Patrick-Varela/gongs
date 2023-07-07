package wrapper

// This script was based on https://stackoverflow.com/a/51350523
// It was slightly modified to fix deprecation warnings
// It is used to copy files on windows with a better performance than the os package
import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	copyFileProc = syscall.MustLoadDLL("kernel32.dll").MustFindProc("CopyFileW")
)

func CopyFile(src, dst string, overwrite bool) error {
	srcW, _ := syscall.UTF16PtrFromString(src)
	dstW, _ := syscall.UTF16PtrFromString(dst)

	var failIfExists uintptr
	if overwrite {
		failIfExists = 0
	} else {
		failIfExists = 1
	}

	out, _, err := copyFileProc.Call(
		uintptr(unsafe.Pointer(srcW)),
		uintptr(unsafe.Pointer(dstW)),
		failIfExists)
	if out == 0 {
		panic(fmt.Errorf("copyFile failed: %v", err))
	}

	return nil
}
