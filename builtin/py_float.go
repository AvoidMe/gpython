package builtin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/maphash"
)

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, f)
	return buf.Bytes()
}

type PyFloat struct {
	Value float64
}

func (self *PyFloat) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyFloat) Repr() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyFloat) Hash() (int64, error) {
	if float64(int64(self.Value)) == self.Value {
		return int64(self.Value), nil
	}
	h := maphash.Hash{}
	h.SetSeed(*GetPyHashSeed())
	h.Write(float64ToByte(self.Value))
	return int64(h.Sum64()), nil
}

func (self *PyFloat) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyFloat:
		if bb.Value == self.Value {
			return PyTrue
		}
		return PyFalse
	case *PyInt:
		if float64(bb.Int64()) == self.Value {
			return PyTrue
		}
		return PyFalse
	case *PyBool:
		if float64(bb.IntValue().Int64()) == self.Value {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self *PyFloat) BinaryAdd(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyFloat:
		return &PyFloat{Value: bb.Value + self.Value}
	case *PyInt:
		return &PyFloat{Value: self.Value + float64(bb.Int64())}
	case *PyBool:
		return &PyFloat{Value: float64(bb.IntValue().Int64()) + self.Value}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}

func (self *PyFloat) BinarySubstract(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyFloat:
		return &PyFloat{Value: self.Value - bb.Value}
	case *PyInt:
		// TODO: this is probably wrong
		return &PyFloat{Value: self.Value - float64(bb.Int64())}
	case *PyBool:
		// TODO: this is probably wrong
		return &PyFloat{Value: self.Value - float64(bb.IntValue().Int64())}
	default:
		panic("Can't substract number and non-number") // TODO: properly handle an error
	}
}
