package builtin

import (
	"fmt"
)

type PyInt struct {
	Value int64
}

func (self PyInt) String() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self PyInt) Repr() string {
	return fmt.Sprintf("%v", self.Value)
}

func (self PyInt) BinaryAdd(b PyObject) PyObject {
	switch bb := b.(type) {
	case PyInt:
		return PyInt{Value: self.Value + bb.Value}
	case PyFloat:
		return PyFloat{Value: bb.Value + float64(self.Value)}
	case PyBool:
		return PyInt{Value: bb.IntValue() + self.Value}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}
