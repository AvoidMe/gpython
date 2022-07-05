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

func (self PyFloat) BinaryAdd(b PyObject) PyObject {
	bInt, success := b.(PyInt)
	if success {
		return PyFloat{Value: self.Value + float64(bInt.Value)}
	}
	bFloat, success := b.(PyFloat)
	if success {
		return PyFloat{Value: bFloat.Value + self.Value}
	}
	panic("Can't add number and non-number") // TODO: properly handle an error
}
