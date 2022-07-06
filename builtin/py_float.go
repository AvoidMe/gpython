package builtin

import (
	"fmt"
)

type PyFloat struct {
	Value float64
}

func (self PyFloat) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self PyFloat) Repr() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self PyFloat) Equal(b PyObject) PyBool {
	switch bb := b.(type) {
	case PyFloat:
		if bb.Value == self.Value {
			return PyTrue
		}
		return PyFalse
	case PyInt:
		if float64(bb.Value) == self.Value {
			return PyTrue
		}
		return PyFalse
	case PyBool:
		if float64(bb.IntValue()) == self.Value {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self PyFloat) BinaryAdd(b PyObject) PyObject {
	switch bb := b.(type) {
	case PyFloat:
		return PyFloat{Value: bb.Value + self.Value}
	case PyInt:
		return PyFloat{Value: self.Value + float64(bb.Value)}
	case PyBool:
		return PyFloat{Value: float64(bb.IntValue()) + self.Value}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}
