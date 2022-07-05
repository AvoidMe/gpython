package gpythonint

import (
	"fmt"
	"main/pyobject"
)

// int
type GpythonInt struct {
	Int int64
}

func (v GpythonInt) String() string {
	return fmt.Sprintf("%v", v.Int)
}

func (v GpythonInt) Repr() string {
	return fmt.Sprintf("%v", v.Int)
}

func (v GpythonInt) BinaryAdd(b pyobject.PyObject) pyobject.PyObject {
	bInt, success := b.(GpythonInt)
	if success {
		return GpythonInt{Int: v.Int + bInt.Int}
	}
	bFloat, success := b.(GpythonFloat)
	if success {
		return GpythonFloat{Float: bFloat.Float + float64(v.Int)}
	}
	panic("Can't add number and non-number") // TODO: properly handle an error
}

// float
type GpythonFloat struct {
	Float float64
}

func (v GpythonFloat) String() string {
	return fmt.Sprintf("%v", v.Float)
}

func (v GpythonFloat) Repr() string {
	return fmt.Sprintf("%v", v.Float)
}

func (v GpythonFloat) BinaryAdd(b pyobject.PyObject) pyobject.PyObject {
	bInt, success := b.(GpythonInt)
	if success {
		return GpythonFloat{Float: v.Float + float64(bInt.Int)}
	}
	bFloat, success := b.(GpythonFloat)
	if success {
		return GpythonFloat{Float: bFloat.Float + v.Float}
	}
	panic("Can't add number and non-number") // TODO: properly handle an error
}
