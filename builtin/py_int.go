package builtin

import (
	"fmt"
)

//func int64ToByte(f int64) []byte {
//	var buf bytes.Buffer
//	binary.Write(&buf, binary.BigEndian, f)
//	return buf.Bytes()
//}

type PyInt struct {
	Value int64
}

func (self *PyInt) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyInt) Repr() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyInt) Hash() (uint64, error) {
	// CPython returns different hash for objects, which are more than 2**30
	return uint64(self.Value), nil
	//h := maphash.Hash{}
	//h.SetSeed(*GetPyHashSeed())
	//h.Write(int64ToByte(self.Value))
	//return h.Sum64(), nil
}

func (self *PyInt) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyFloat:
		if bb.Value == float64(self.Value) {
			return PyTrue
		}
		return PyFalse
	case *PyInt:
		if bb.Value == self.Value {
			return PyTrue
		}
		return PyFalse
	case *PyBool:
		if bb.IntValue() == self.Value {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self *PyInt) BinaryAdd(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		return &PyInt{Value: self.Value + bb.Value}
	case *PyFloat:
		return &PyFloat{Value: bb.Value + float64(self.Value)}
	case *PyBool:
		return &PyInt{Value: bb.IntValue() + self.Value}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}

func (self *PyInt) BinarySubstract(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		return &PyInt{Value: self.Value - bb.Value}
	case *PyFloat:
		return &PyFloat{Value: float64(self.Value) - bb.Value}
	case *PyBool:
		return &PyInt{Value: self.Value - bb.IntValue()}
	default:
		panic("Can't substract number and non-number") // TODO: properly handle an error
	}
}
