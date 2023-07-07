package wrapper

/*
#cgo CFLAGS: -I./inc

#include <stdbool.h>
#include <stdlib.h>

#include "sharedspice.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Patrick-Varela/gongs/ngcallback"
	"github.com/Patrick-Varela/gongs/ngtype"
)

func Init(c *ngcallback.Callbacks) *NgInstance {
	i := newNgInstance(c)
	i.init()
	return i
}

func (ng *NgInstance) init() int {
	ng.call(ng.initP, sendCharGtw, sendStatGtw, controlledExitGtw, sendDataGtw, sendInitDataGtw, bgThreadRunningGtw, unsafe.Pointer(ng))
	return 0
}

func (ng *NgInstance) GetUID() string {
	return ng.uid
}

func (ng *NgInstance) Quit() int {
	c := unsafe.Pointer(C.CString("quit"))
	defer C.free(c)
	ng.call(ng.commandP, c)
	ng.release()
	return 0
}

func (ng *NgInstance) Command(s string) int {
	c := unsafe.Pointer(C.CString(s))
	defer C.free(unsafe.Pointer(c))
	o := int(ng.call(ng.commandP, c))
	return o
}

func (ng *NgInstance) Clear() int {
	return int(ng.call(ng.commandP, nil))
}

func (ng *NgInstance) Running() bool {
	return ng.call(ng.runningP) != 0 //I have no idea if this will work. I'm assuming that 0 is false and anything else is true
}

func (ng *NgInstance) GetVecInfo(name string) *ngtype.VectorInfo {
	var ret *ngtype.VectorInfo
	str := unsafe.Pointer(C.CString(name))
	defer C.free(str)
	if pvec := C.pvector_info(unsafe.Pointer(ng.call(ng.GetVecInfoP, str))); pvec != nil {
		ret = ng.curVecs.storeVectorInfo(pvec)
	}
	return ret
}

func (ng *NgInstance) Circ(circa []string) int {
	n := len(circa)
	var cStrings [256]*C.char
	for i := 0; i < n; i++ {
		cStrings[i] = C.CString(circa[i])
		defer C.free(unsafe.Pointer(cStrings[i]))
	}
	cStrings[n] = nil
	pCString := unsafe.Pointer(&cStrings[0])
	o := int(ng.call(ng.circP, pCString))
	return o
}

func (ng *NgInstance) CurPlot() string {
	var ret string
	if pchar := (*C.char)(unsafe.Pointer(ng.call(ng.curPlotP))); pchar != nil {
		ret = C.GoString(pchar)
	}
	return ret
}

func (ng *NgInstance) AllPlots() []string {
	var ret []string
	if ppchar := (**C.char)(unsafe.Pointer(ng.call(ng.allPlotsP))); ppchar != nil {
		ret = ppcharToStringSlice(ppchar)
	}
	return ret
}

func (ng *NgInstance) AllVecs(s string) []string {
	var ret []string
	c := unsafe.Pointer(C.CString(s))
	defer C.free(c)
	if ppchar := (**C.char)(unsafe.Pointer(ng.call(ng.allVecsP, c))); ppchar != nil {
		ret = ppcharToStringSlice(ppchar)
	}
	return ret
}

func ClearCache() {
	clearCache()
}
