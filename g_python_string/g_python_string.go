package gpythonstring

import (
	"main/pyobject"
)

type GpythonString struct {
	Str string
}

func (v GpythonString) String() string {
	return v.Str
}

func (v GpythonString) Repr() string {
	return "'" + v.Str + "'"
}

func (v GpythonString) BinaryAdd(b pyobject.PyObject) pyobject.PyObject {
	bStr, success := b.(GpythonString)
	if !success {
		panic("Can't add string and non-string") // TODO: properly handle an error
	}
	return GpythonString{Str: v.Str + bStr.Str}
}
