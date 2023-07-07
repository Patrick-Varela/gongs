// A go ngspice wrapper
//
// gongs is a wrapper around the ngspice shared library.
// It provides a interface that you can work with native go types,
// and have multiple ngspice instances running at the same time.
//
// To use gongs, you need to have the ngspice shared library DLL or SO file in the same directory as your executable.
// You can download the ngspice shared library from the ngspice website.
package gongs

import (
	"github.com/Patrick-Varela/gongs/internal/wrapper"
	"github.com/Patrick-Varela/gongs/ngcallback"
	"github.com/Patrick-Varela/gongs/ngtype"
)

// NgSpice is the main struct used to interact with the ngspice library
// it's a wrapper around the NgInstance struct.
type NgSpice struct{ inst *wrapper.NgInstance }
type NgCallbacks = ngcallback.Callbacks
type NgIReturn = ngcallback.NgI

// Init initializes the ngspice library and returns a pointer to the NgSpice instance.
//
// The NgCallbacks struct is used to pass the callback functions to the ngspice library and cannot be nil
// but can have every field set to nil.
//
// The NgCallbacks struct is copied and can be reused after the call to Init.
// The NgSpice instance should be closed with the Quit method when it is no longer needed.
func Init(callbacks *NgCallbacks) *NgSpice {
	return &NgSpice{inst: wrapper.Init(callbacks)}
}

// GetUID returns the UID of the ngspice instance.
func (ngspice *NgSpice) GetUID() string {
	return ngspice.inst.GetUID()
}

// Command sends a command to the ngspice instance.
//
// The command string is sent to the ngspice library and
// The return value is 0 if the command was successful and 1 if the command failed.
func (ngspice *NgSpice) Command(command string) int {
	return ngspice.inst.Command(command)
}

// Clear clears the ngspice instance's internal control structures.
//
// The return value is 0 if the command was successful and 1 if the command failed.
func (ngspice *NgSpice) Clear() int {
	return ngspice.inst.Clear()
}

// Quit closes the ngspice instance
func (ngspice *NgSpice) Quit() int {
	return ngspice.inst.Quit()
}

// Running returns true if the ngspice instance is running.
func (ngspice *NgSpice) Running() bool {
	return ngspice.inst.Running()
}

// GetVar returns the value of a variable in the ngspice instance.
//
// The VecName parameter is the name of the Vector (may be in the form ’vectorname’ or <plotname>.vectorname) to get the info for.
// The return value is a [ngtype.VectorInfo] pointer.
func (ngspice *NgSpice) GetVecInfo(vecName string) *ngtype.VectorInfo {
	return ngspice.inst.GetVecInfo(vecName)
}

// Circ sends a circuit to the ngspice instance.
//
// Each string in the circ slice is a line of the circuit.
// The return value is 0 if the circuit load was successful and 1 if the circuit load failed.
func (ngspice *NgSpice) Circ(circ []string) int {
	return ngspice.inst.Circ(circ)
}

// CurPlot returns the name of the current plot.
func (ngspice *NgSpice) CurPlot() string {
	return ngspice.inst.CurPlot()
}

// AllPlots returns an array of all plot names.
func (ngspice *NgSpice) AllPlots() []string {
	return ngspice.inst.AllPlots()
}

/*
return to the caller an array of vector names in the plot named by plotname
*/
func (ngspice *NgSpice) AllVecs(plotname string) []string {
	return ngspice.inst.AllVecs(plotname)
}

/*
ClearCache clears the cache of the ngspice instances.
*/
func ClearCache() {
	wrapper.ClearCache()
}
