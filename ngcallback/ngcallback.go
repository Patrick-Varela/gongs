package ngcallback

import (
	"github.com/Patrick-Varela/gongs/ngtype"
)

type NgI interface {
	GetUID() string
	Command(command string) int
	Clear() int
	Quit() int
	Running() bool
	GetVecInfo(vecName string) *ngtype.VectorInfo
	Circ(circ []string) int
	CurPlot() string
	AllPlots() []string
	AllVecs(plotname string) []string
}

// SendChar is a function for reading printf, fprintf.
//
// The first argument is the string to be printed.
// The second argument is the identification number of the calling ngspice shared library.
// The third argument is the return pointer received from the caller NgSpice instance.
type SendChar func(string, int, NgI) int

// SendStat is a function for the status string and percent value.
//
// The first argument is the simulation status and value (in percent).
// The second argument is the identification number of the calling ngspice shared library.
// The third argument is the return pointer received from the caller NgSpice instance.
type SendStat func(string, int, NgI) int

// ControlledExit is a function for setting a 'quit' signal in caller.
//
// The first argument is the exit status.
// The second argument is a boolean indicating whether the shared library should be immediately unloaded (true)
// or just set a flag, and unload when the function has returned (false).
// The third argument is a boolean indicating whether the exit is due to a 'quit' command (true)
// or an ngspice.dll error (false).
// The fourth argument is the identification number of the calling ngspice shared library.
// The fifth argument is the return pointer received from the caller NgSpice instance.
type ControlledExit func(int, bool, bool, int, NgI) int

// SendData is a function  for sending an array of structs containing data values of all vectors in the current plot (simulation output)
//
// The first argument is a pointer to a struct containing the data values of all vectors in the current plot.
// The second argument is the number of vectors in the current plot.
// The third argument is the identification number of the calling ngspice shared library.
// The fourth argument is the return pointer received from the caller NgSpice instance.
type SendData func(*ngtype.VecValuesAll, int, int, NgI) int

// SendInitData is a function for sending an array of structs containing info on all vectors in the current plot (immediately before simulation starts)
//
// The first argument is a pointer to a struct containing the info on all vectors in the current plot.
// The second argument is the identification number of the calling ngspice shared library.
// The third argument is the return pointer received from the caller NgSpice instance.
type SendInitData func(*ngtype.VecInfoAll, int, NgI) int

// BGThreadRunning is a function for sending a boolean signal (true if thread is running)
//
// The first argument is a boolean indicating whether the background thread is running (true) or not (false).
// The second argument is the identification number of the calling ngspice shared library.
// The third argument is the return pointer received from the caller NgSpice instance.
type BGThreadRunning func(bool, int, NgI) int

// type SendEvtData func(int, float64, float64, string, unsafe.Pointer, int, int, int, unsafe.Pointer) int
// type SendInitEvtData func(int, int, string, string, int, unsafe.Pointer) int

// Callbacks is a struct containing the callback functions that will be passed to the ngspice library.
type Callbacks struct {
	SendChar        SendChar
	SendStat        SendStat
	ControlledExit  ControlledExit
	SendData        SendData
	SendInitData    SendInitData
	BGThreadRunning BGThreadRunning
	// SendEvtData     SendEvtData
	// SendInitEvtData SendInitEvtData
}

// Copy returns a copy of the Callbacks struct.
func (c *Callbacks) Copy() *Callbacks {
	cb := Callbacks{}
	cb.SendChar = c.SendChar
	cb.SendStat = c.SendStat
	cb.ControlledExit = c.ControlledExit
	cb.SendData = c.SendData
	cb.SendInitData = c.SendInitData
	cb.BGThreadRunning = c.BGThreadRunning
	// cb.SendEvtData = c.SendEvtData
	// cb.SendInitEvtData = c.SendInitEvtData
	return &cb
}
