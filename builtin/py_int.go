package builtin

import (
	"fmt"
	"math/big"
)

//func int64ToByte(f int64) []byte {
//	var buf bytes.Buffer
//	binary.Write(&buf, binary.BigEndian, f)
//	return buf.Bytes()
//}

type PyInt struct {
	Value *big.Int
}

func (self *PyInt) Int64() int64 {
	return self.Value.Int64()
}

func (self *PyInt) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyInt) Repr() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self *PyInt) Hash() (int64, error) {
	// CPython returns different hash for objects, which are more than 2**30
	return self.Int64(), nil
	//h := maphash.Hash{}
	//h.SetSeed(*GetPyHashSeed())
	//h.Write(int64ToByte(self.Value))
	//return h.Sum64(), nil
}

func (self *PyInt) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyFloat:
		// TODO: this is probably wrong
		if bb.Value == float64(self.Int64()) {
			return PyTrue
		}
		return PyFalse
	case *PyInt:
		if bb.Value.Cmp(self.Value) == 0 {
			return PyTrue
		}
		return PyFalse
	case *PyBool:
		if bb.IntValue().Value.Cmp(self.Value) == 0 {
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
		sum := big.NewInt(0)
		return &PyInt{Value: sum.Add(self.Value, bb.Value)}
	case *PyFloat:
		// TODO: this is probably wrong
		return &PyFloat{Value: bb.Value + float64(self.Int64())}
	case *PyBool:
		sum := big.NewInt(bb.IntValue().Int64())
		return &PyInt{Value: sum.Add(self.Value, sum)}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}

func (self *PyInt) BinarySubstract(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		sum := big.NewInt(0)
		return &PyInt{Value: sum.Sub(self.Value, bb.Value)}
	case *PyFloat:
		// TODO: this is probably wrong
		return &PyFloat{Value: float64(self.Int64()) - bb.Value}
	case *PyBool:
		sum := big.NewInt(0)
		sum.Add(sum, self.Value)
		return &PyInt{Value: sum.Sub(sum, bb.IntValue().Value)}
	default:
		panic("Can't substract number and non-number") // TODO: properly handle an error
	}
}

func NewPyInt(value int64) *PyInt {
	return &PyInt{Value: big.NewInt(value)}
}
