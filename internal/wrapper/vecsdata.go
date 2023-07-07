package wrapper

/*
#include <stdbool.h>

#include "sharedspice.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Patrick-Varela/gongs/ngtype"
)

/* Dvec flags. */
const (
	vf_real      = (1 << 0) /* The data is real. */
	vf_complex   = (1 << 1) /* The data is complex. */
	vf_accum     = (1 << 2) /* writedata should save this vector. */
	vf_plot      = (1 << 3) /* writedata should incrementally plot it. */
	vf_print     = (1 << 4) /* writedata should print this vector. */
	vf_mingiven  = (1 << 5) /* The v_minsignal value is valid. */
	vf_maxgiven  = (1 << 6) /* The v_maxsignal value is valid. */
	vf_permanent = (1 << 7) /* Don't garbage collect this vector. */
)

/* Plot types. */
const (
	plot_lin   = 1
	plot_comb  = 2
	plot_point = 3
)

type curVecsData struct {
	curVecValsAll *ngtype.VecValuesAll
	curVecInfoAll *ngtype.VecInfoAll
	curVecInfo    *ngtype.VecInfo
	curVectorInfo *ngtype.VectorInfo
}

func (vecData *curVecsData) storeCurVecValsAll(pvva C.pvecvaluesall) {
	veccount := int(pvva.veccount)

	if vecData.curVecValsAll == nil {
		vecData.curVecValsAll = new(ngtype.VecValuesAll)
	}
	if vecData.curVecValsAll.VecCount != veccount {
		vecData.curVecValsAll.VecsA = make([]*ngtype.VecValues, veccount)
	}

	vecData.curVecValsAll.VecCount = veccount
	vecData.curVecValsAll.VecIndex = int(pvva.vecindex)
	for i, v := range unsafe.Slice(pvva.vecsa, vecData.curVecValsAll.VecCount) {
		storeVecVal(v, &vecData.curVecValsAll.VecsA[i])
	}
}

func storeVecVal(pvecv C.pvecvalues, vecsA **ngtype.VecValues) {
	if pvecv == nil {
		return
	}
	if *vecsA == nil {
		*vecsA = new(ngtype.VecValues)
	}
	(*vecsA).Name = C.GoString(pvecv.name)
	(*vecsA).CReal = float64(pvecv.creal)
	(*vecsA).CImag = float64(pvecv.cimag)
	(*vecsA).IsScale = bool(pvecv.is_scale)
	(*vecsA).IsComplex = bool(pvecv.is_complex)
}

func (vecData *curVecsData) storeCurVecInfoAll(pvia C.pvecinfoall) {
	veccount := int(pvia.veccount)

	if vecData.curVecInfoAll == nil {
		vecData.curVecInfoAll = new(ngtype.VecInfoAll)
	}
	if vecData.curVecInfoAll.VecCount != veccount {
		vecData.curVecInfoAll.Vecs = make([]*ngtype.VecInfo, veccount)
	}

	vecData.curVecInfoAll.Date = C.GoString(pvia.date)
	vecData.curVecInfoAll.Name = C.GoString(pvia.name)
	vecData.curVecInfoAll.Title = C.GoString(pvia.title)
	vecData.curVecInfoAll.VecCount = int(pvia.veccount)
	for i, v := range unsafe.Slice(pvia.vecs, vecData.curVecInfoAll.VecCount) {
		storeVecInfo(v, &vecData.curVecInfoAll.Vecs[i])
	}
}

func storeVecInfo(pveci C.pvecinfo, vecs **ngtype.VecInfo) {
	if pveci == nil {
		return
	}
	if *vecs == nil {
		*vecs = new(ngtype.VecInfo)
	}
	(*vecs).IsReal = bool(pveci.is_real)
	(*vecs).Number = int(pveci.number)
	// (*vecs).PDVec = pveci.pdvec
	// (*vecs).PDVecScale = pveci.pdvecscale
	(*vecs).VecName = C.GoString(pveci.vecname)
}

func (vecData *curVecsData) storeVectorInfo(pvectori C.pvector_info) *ngtype.VectorInfo {
	if pvectori != nil {
		if vecData.curVectorInfo == nil {
			vecData.curVectorInfo = new(ngtype.VectorInfo)
		}
		vecData.curVectorInfo.VName = C.GoString(pvectori.v_name)
		vecData.curVectorInfo.VType = int(pvectori.v_type)
		vecData.curVectorInfo.VFlags = int16(pvectori.v_flags)
		vecData.curVectorInfo.VLength = int(pvectori.v_length)
		if vecData.curVectorInfo.VFlags&vf_real != 0 {
			vecData.curVectorInfo.VRealData = unsafe.Slice((*float64)(pvectori.v_realdata), vecData.curVectorInfo.VLength)
		}
		if vecData.curVectorInfo.VFlags&vf_complex != 0 {
			vecData.curVectorInfo.VCompData = make([]complex128, vecData.curVectorInfo.VLength)
			for i, v := range unsafe.Slice(pvectori.v_compdata, vecData.curVectorInfo.VLength) {
				vecData.curVectorInfo.VCompData[i] = complex(v.cx_real, v.cx_imag)
			}
		}
	}
	return vecData.curVectorInfo
}
