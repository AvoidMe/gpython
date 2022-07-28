package builtin

import "unsafe"

type NoneType struct{}

func (*NoneType) Repr() string {
	return "None"
}

func (*NoneType) String() string {
	return "None"
}

func (self *NoneType) Equal(b PyObject) *PyBool {
	switch b.(type) {
	case *NoneType:
		return PyTrue
	default:
		return PyFalse
	}
}

func (self *NoneType) Hash() (uint64, error) {
	return *(*uint64)(unsafe.Pointer(self)), nil
}
