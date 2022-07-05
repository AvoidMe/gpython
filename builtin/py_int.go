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
	bInt, success := b.(PyInt)
	if success {
		return PyInt{Value: self.Value + bInt.Value}
	}
	bFloat, success := b.(PyFloat)
	if success {
		return PyFloat{Value: bFloat.Value + float64(self.Value)}
	}
	panic("Can't add number and non-number") // TODO: properly handle an error
}
