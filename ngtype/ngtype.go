package ngtype

type VectorInfo struct {
	VName     string
	VType     int
	VFlags    int16
	VRealData []float64
	VCompData []complex128
	VLength   int
}

type VecInfo struct {
	Number  int
	VecName string
	IsReal  bool
	// PDVec      unsafe.Pointer
	// PDVecScale unsafe.Pointer
}

type VecInfoAll struct {
	Name     string
	Title    string
	Date     string
	Type     string
	VecCount int
	Vecs     []*VecInfo
}

type VecValues struct {
	Name      string
	CReal     float64
	CImag     float64
	IsScale   bool
	IsComplex bool
}

type VecValuesAll struct {
	VecCount int
	VecIndex int
	VecsA    []*VecValues
}

func (v *VecValuesAll) DeepCopy() *VecValuesAll {
	newvec := new(VecValuesAll)
	newvec.VecCount = v.VecCount
	newvec.VecIndex = v.VecIndex
	newvec.VecsA = make([]*VecValues, len(v.VecsA))
	for i, vec := range v.VecsA {
		newvec.VecsA[i] = vec.DeepCopy()
	}
	return newvec
}

func (v *VecValues) DeepCopy() *VecValues {
	newvec := new(VecValues)
	newvec.Name = v.Name
	newvec.CReal = v.CReal
	newvec.CImag = v.CImag
	newvec.IsScale = v.IsScale
	newvec.IsComplex = v.IsComplex
	return newvec
}

func (v *VecInfoAll) DeepCopy() *VecInfoAll {
	newvec := new(VecInfoAll)
	newvec.Name = v.Name
	newvec.Title = v.Title
	newvec.Date = v.Date
	newvec.Type = v.Type
	newvec.VecCount = v.VecCount
	newvec.Vecs = make([]*VecInfo, len(v.Vecs))
	for i, vec := range v.Vecs {
		newvec.Vecs[i] = vec.DeepCopy()
	}
	return newvec
}

func (v *VecInfo) DeepCopy() *VecInfo {
	newvec := new(VecInfo)
	newvec.Number = v.Number
	newvec.VecName = v.VecName
	newvec.IsReal = v.IsReal
	return newvec
}

func (v *VectorInfo) DeepCopy() *VectorInfo {
	newvec := new(VectorInfo)
	newvec.VName = v.VName
	newvec.VType = v.VType
	newvec.VFlags = v.VFlags
	newvec.VRealData = make([]float64, len(v.VRealData))
	copy(newvec.VRealData, v.VRealData)
	newvec.VCompData = make([]complex128, len(v.VCompData))
	copy(newvec.VCompData, v.VCompData)
	newvec.VLength = v.VLength
	return newvec
}
