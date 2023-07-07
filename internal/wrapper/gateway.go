package wrapper

/*
#include <stdbool.h>
#include <stdlib.h>

#include "sharedspice.h"
#include "gateway.h"
*/
import "C"
import (
	"unsafe"
)

// HACK: If ever someone arrives here, I'm sorry for the mess.
// Since I know I'll never dereference the pointer,
// I probably can use it more unsafely than it already is with uintptr_t, at least in Linux.
var (
	sendCharGtw        = C.sendChar_cgo
	sendStatGtw        = C.sendStat_cgo
	controlledExitGtw  = C.controlledExit_cgo
	sendDataGtw        = C.sendData_cgo
	sendInitDataGtw    = C.sendInitData_cgo
	bgThreadRunningGtw = C.bgThreadRunning_cgo
)

//export sendStat
func sendStat(s *C.char, i int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if ng.cbs.SendStat != nil {
		return ng.cbs.SendStat(C.GoString(s), i, ng)
	}
	return -1
}

//export sendChar
func sendChar(s *C.char, i int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if f := ng.cbs.SendChar; f != nil {
		return f(C.GoString(s), i, ng)
	}
	return -1
}

//export controlledExit
func controlledExit(i1 int, b1 bool, b2 bool, i2 int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if f := ng.cbs.ControlledExit; f != nil {
		return f(i1, b1, b2, i2, ng)
	}
	return -1
}

//export sendData
func sendData(pvva C.pvecvaluesall, i1 int, i2 int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if f := ng.cbs.SendData; f != nil {
		ng.curVecs.storeCurVecValsAll(pvva)
		return f(ng.curVecs.curVecValsAll, i1, i2, ng)
	}
	return -1
}

//export sendInitData
func sendInitData(pvia C.pvecinfoall, i int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if f := ng.cbs.SendInitData; f != nil {
		ng.curVecs.storeCurVecInfoAll(pvia)
		return f(ng.curVecs.curVecInfoAll, i, ng)
	}
	return -1
}

//export bgThreadRunning
func bgThreadRunning(b bool, i int, ptr unsafe.Pointer) int {
	ng := ngiFromPointer(ptr)
	if f := ng.cbs.BGThreadRunning; f != nil {
		return f(b, i, ng)
	}
	return -1
}
