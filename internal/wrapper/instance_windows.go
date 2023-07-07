package wrapper

import (
	"os"
	"syscall"
	"unsafe"
)

type procHdl = *syscall.Proc
type libHdl = *syscall.DLL

const (
	libName       = "ngspice.dll"
	spinitEnv     = "SPICE_SCRIPTS"
	spinitPath    = "\\share\\ngspice\\scripts\\"
	cachePath     = "\\cache\\instances\\"
	fileExt       = ".dll"
	pathSeparator = '\\'
)

func (i *NgInstance) new() {
	setSpinitPath(spinitPath)
	i.cacheLib()
	h := syscall.MustLoadDLL(i.libPath)
	i.handler = h
	i.initP = h.MustFindProc(initName)
	i.initSyncP = h.MustFindProc(initSyncName)
	i.commandP = h.MustFindProc(commandName)
	i.runningP = h.MustFindProc(runningName)
	i.GetVecInfoP = h.MustFindProc(getVecInfoName)
	i.circP = h.MustFindProc(circName)
	i.curPlotP = h.MustFindProc(curPlotName)
	i.allPlotsP = h.MustFindProc(allPlotsName)
	i.allVecsP = h.MustFindProc(allVecsName)
	// i.setBkptP = h.MustFindProc(setBkptName)
	// i.cmInputPathP = h.MustFindProc(cmInputPathName)
	// i.initEvtP = h.MustFindProc(initEvtName)
	// i.evtNodeInfoP = h.MustFindProc(evtNodeInfoName)
	// i.allEvtNodesP = h.MustFindProc(allEvtNodesName)
}

func (i *NgInstance) Release() {
	i.handler.Release()
	os.Remove(i.libPath)
}

func (i *NgInstance) Call(p procHdl, args []unsafe.Pointer) uintptr {
	var ret uintptr
	var err error
	switch p.Name {
	case initName:
		ret, _, err = p.Call(uintptr(args[0]), uintptr(args[1]), uintptr(args[2]), uintptr(args[3]), uintptr(args[4]), uintptr(args[5]), uintptr(args[6]))
	case commandName:
		ret, _, err = p.Call(uintptr(args[0]))
	case runningName:
		ret, _, err = p.Call()
	case getVecInfoName:
		ret, _, err = p.Call(uintptr(args[0]))
	case circName:
		ret, _, err = p.Call(uintptr(args[0]))
	case curPlotName:
		ret, _, err = p.Call()
	case allPlotsName:
		ret, _, err = p.Call()
	}
	if err != nil {
		// fmt.Printf("%s err.Error(): %v\n", p.Name, err)
	}
	return ret
}
