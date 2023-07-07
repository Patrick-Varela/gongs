package wrapper

import "C"
import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func ppcharToStringSlice(ppchar **C.char) []string {
	var goStrings []string
	for i := 0; ; i++ {
		cString := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ppchar)) + uintptr(i)*unsafe.Sizeof(ppchar)))
		if cString == nil {
			break
		}
		goStrings = append(goStrings, C.GoString(cString))
	}
	return goStrings
}

func clearCache() {
	dir := getCacheDir()
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var strBuff strings.Builder
	for _, file := range files {
		strBuff.WriteString(dir)
		strBuff.WriteString(file.Name())
		err := os.Remove(strBuff.String())
		if err != nil {
			fmt.Println(err)
		}
		strBuff.Reset()
	}
}
