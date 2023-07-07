package wrapper

import (
	"fmt"
	"os"
	"unsafe"

	/*
		#include <dlfcn.h>
		#include <stdlib.h>
		#include "instancegateway_linux.h"
	*/
	"C"
)

type symbolHdl struct {
	name string
	ptr  unsafe.Pointer
}

type procHdl = *symbolHdl
type libHdl = unsafe.Pointer

const (
	libName       = "libngspice.so"
	spinitEnv     = "SPICE_SCRIPTS"
	spinitPath    = "/share/ngspice/scripts/"
	cachePath     = "/cache/instances/"
	fileExt       = ".so"
	pathSeparator = '/'
)

func (i *NgInstance) new() {
	setSpinitPath(spinitPath)
	i.cacheLib()

	i.handler = i.loadLib()
	i.initP = loadSymbol(i.handler, initName)
	i.initSyncP = loadSymbol(i.handler, initSyncName)
	i.commandP = loadSymbol(i.handler, commandName)
	i.runningP = loadSymbol(i.handler, runningName)
	i.GetVecInfoP = loadSymbol(i.handler, getVecInfoName)
	i.circP = loadSymbol(i.handler, circName)
	i.curPlotP = loadSymbol(i.handler, curPlotName)
	i.allPlotsP = loadSymbol(i.handler, allPlotsName)
	i.allVecsP = loadSymbol(i.handler, allVecsName)
}

func (i *NgInstance) loadLib() libHdl {
	ln := C.CString(i.libPath)
	defer C.free(unsafe.Pointer(ln))

	h := C.dlopen(ln, C.RTLD_LAZY|C.RTLD_GLOBAL)
	if h == nil {
		panic(fmt.Errorf("dlopen failed: %s", C.GoString(C.dlerror())))
	}
	return h
}

func loadSymbol(lib libHdl, symbolName string) procHdl {
	symbol := C.CString(symbolName)
	defer C.free(unsafe.Pointer(symbol))

	sym := new(symbolHdl)
	ptr, err := C.dlsym(lib, symbol)
	if err != nil {
		panic(err)
	}
	sym.ptr = ptr
	sym.name = symbolName
	return sym
}

func (i *NgInstance) Release() {
	C.dlclose(i.handler)
	os.Remove(i.libPath)
}

func convertArgs(args []unsafe.Pointer) []C.uintptr_t {
	cArgs := make([]C.uintptr_t, len(args))
	for i, arg := range args {
		cArgs[i] = C.uintptr_t(uintptr(arg))
	}
	return cArgs
}

func (i *NgInstance) Call(p procHdl, args []unsafe.Pointer) uintptr {
	var ret uintptr
	cArgs := convertArgs(args)
	switch p.name {
	case initName:
		C.init_ng(p.ptr, cArgs[0], cArgs[1], cArgs[2], cArgs[3], cArgs[4], cArgs[5], cArgs[6])
	case commandName:
		C.command_ng(p.ptr, cArgs[0])
	case runningName:
		C.running_ng(p.ptr)
	case getVecInfoName:
		C.getVecInfo_ng(p.ptr, cArgs[0])
	case circName:
		C.circ_ng(p.ptr, cArgs[0])
	case curPlotName:
		C.curPlot_ng(p.ptr)
	case allPlotsName:
		C.allPlots_ng(p.ptr)
	}
	if ret < 0 {
		panic(fmt.Errorf("call to %s failed", p.name))
	}
	return ret
}
