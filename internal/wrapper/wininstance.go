package wrapper

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// 	"syscall"
// 	"unsafe"
// )

// func (i *NgInstance) cacheLib() {
// 	var strBuff strings.Builder
// 	strBuff.WriteString("./cache/instances/")
// 	strBuff.WriteString(i.uid)
// 	println(i.uid)
// 	strBuff.WriteString(".dll")
// 	i.dllPath = strBuff.String()
// 	srcFile, err := os.Open("./ngspice.dll")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer srcFile.Close()
// 	dstFile, err := os.Create(i.dllPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer dstFile.Close()

// 	_, err = io.Copy(dstFile, srcFile)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (i *NgInstance) newSyscallNg() {
// 	i.cacheLib()
// 	h := syscall.MustLoadDLL(i.dllPath)
// 	i.isSyscall = true
// 	i.handler = h
// 	i.initP = h.MustFindProc(initName)
// 	i.initSyncP = h.MustFindProc(initSyncName)
// 	i.commandP = h.MustFindProc(commandName)
// 	i.runningP = h.MustFindProc(runningName)
// 	i.GetVecInfoP = h.MustFindProc(getVecInfoName)
// 	i.circP = h.MustFindProc(circName)
// 	i.curPlotP = h.MustFindProc(curPlotName)
// 	i.allPlotsP = h.MustFindProc(allPlotsName)
// 	i.allVecsP = h.MustFindProc(allVecsName)
// 	// i.setBkptP = h.MustFindProc(setBkptName)
// 	// i.cmInputPathP = h.MustFindProc(cmInputPathName)
// 	// i.initEvtP = h.MustFindProc(initEvtName)
// 	// i.evtNodeInfoP = h.MustFindProc(evtNodeInfoName)
// 	// i.allEvtNodesP = h.MustFindProc(allEvtNodesName)
// }

// func (i *NgInstance) wRelease() {
// 	i.handler.Release()
// 	os.Remove(i.dllPath)
// }

// func (i *NgInstance) wCall(p procHdl, args []unsafe.Pointer) uintptr {
// 	var ret uintptr
// 	var err error
// 	switch p.Name {
// 	case initName:
// 		ret, _, err = p.Call(uintptr(args[0]), uintptr(args[1]), uintptr(args[2]), uintptr(args[3]), uintptr(args[4]), uintptr(args[5]), uintptr(args[6]))
// 	case commandName:
// 		ret, _, err = p.Call(uintptr(args[0]))
// 	case runningName:
// 		ret, _, err = p.Call()
// 	case getVecInfoName:
// 		ret, _, err = p.Call(uintptr(args[0]))
// 	case circName:
// 		ret, _, err = p.Call(uintptr(args[0]))
// 	case curPlotName:
// 		ret, _, err = p.Call()
// 	case allPlotsName:
// 		ret, _, err = p.Call()
// 	}
// 	if err != nil {
// 		fmt.Printf("err.Error(): %v\n", err.Error())
// 	}
// 	return ret
// }
