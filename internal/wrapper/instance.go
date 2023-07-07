package wrapper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/Patrick-Varela/gongs/ngcallback"
)

const (
	//counts as number of zeroes
	uidHalfLen = 10000000

	initName       = "ngSpice_Init"
	initSyncName   = "ngSpice_Init_Sync"
	commandName    = "ngSpice_Command"
	runningName    = "ngSpice_running"
	getVecInfoName = "ngGet_Vec_Info"
	circName       = "ngSpice_Circ"
	curPlotName    = "ngSpice_CurPlot"
	allPlotsName   = "ngSpice_AllPlots"
	allVecsName    = "ngSpice_AllVecs"
	// setBkptName     = "ngSpice_SetBkpt"
	// cmInputPathName = "ngCM_Input_Path"
	// initEvtName     = "ngSpice_Init_Evt"
	// evtNodeInfoName = "ngGet_Evt_NodeInfo"
	// allEvtNodesName = "ngSpice_AllEvtNodes"
)

var (
	cachePrepared = false
	spinitPathSet = false
	srcLibPath    string
	strBuilder    strings.Builder
)

type NgInstance struct {
	curVecs     *curVecsData
	uid         string
	libPath     string
	cbs         *ngcallback.Callbacks
	handler     libHdl
	initP       procHdl
	initSyncP   procHdl
	commandP    procHdl
	runningP    procHdl
	GetVecInfoP procHdl
	circP       procHdl
	curPlotP    procHdl
	allPlotsP   procHdl
	allVecsP    procHdl
	// setBkptP     procHdl
	// cmInputPathP procHdl
	// initEvtP     procHdl
	// evtNodeInfoP procHdl
	// allEvtNodesP procHdl
}

func newNgInstance(cbs *ngcallback.Callbacks) *NgInstance {
	inst := new(NgInstance)
	inst.curVecs = new(curVecsData)
	inst.newUID()
	inst.cbs = cbs.Copy()
	inst.new()
	return inst
}

func (i *NgInstance) release() {
	i.Release()
}

func (i *NgInstance) call(symbol procHdl, args ...unsafe.Pointer) uintptr {
	return i.Call(symbol, args)
}

func ngiFromPointer(p unsafe.Pointer) *NgInstance {
	return (*NgInstance)(p)
}

func (i *NgInstance) newUID() {
	p := int64(uintptr(unsafe.Pointer(i))) % uidHalfLen
	t := (time.Now().UnixMicro() % uidHalfLen) * uidHalfLen
	i.uid = strconv.FormatInt(p+t, 32)
}

func getWD() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func setSpinitPath(path string) {
	if spinitPathSet {
		return
	}
	strBuilder.Reset()
	strBuilder.WriteString(getWD())
	strBuilder.WriteString(path)
	err := os.Setenv(spinitEnv, strBuilder.String())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getCacheDir() string {
	strBuilder.Reset()
	defer strBuilder.Reset()
	strBuilder.WriteString(getWD())
	strBuilder.WriteString(cachePath)
	return strBuilder.String()
}

func prepareForCache() {
	if cachePrepared {
		return
	}
	strBuilder.Reset()
	strBuilder.WriteString(getCacheDir())
	err := os.MkdirAll(strBuilder.String(), os.ModePerm)
	if err != nil {
		panic(err)
	}
	cachePrepared = true
}

func getSrcLibPath() string {
	if srcLibPath == "" {
		strBuilder.Reset()
		strBuilder.WriteString(getWD())
		strBuilder.WriteByte(pathSeparator)
		strBuilder.WriteString(libName)
		srcLibPath = strBuilder.String()
		_, err := os.Stat(srcLibPath)
		if err != nil {
			panic(fmt.Errorf("ngpsice library not found in %s:%s", getWD(), err.Error()))
		}
		return srcLibPath
	}
	return srcLibPath
}

func (i *NgInstance) copyLibToPath() {
	err := CopyFile(getSrcLibPath(), i.libPath, true)
	if err != nil {
		panic(err)
	}
}

func (i *NgInstance) setLibPath() {
	strBuilder.Reset()
	strBuilder.WriteString(getCacheDir())
	strBuilder.WriteString(i.uid)
	strBuilder.WriteString(fileExt)
	i.libPath = strBuilder.String()
}

func (i *NgInstance) cacheLib() {
	prepareForCache()
	i.setLibPath()
	i.copyLibToPath()
}
