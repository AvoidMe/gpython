package builtin

type PyBool struct {
	Value bool
}

func (self PyBool) String() string {
	if !self.Value {
		return "False"
	}
	return "True"
}

func (self PyBool) Repr() string {
	return self.String()
}

func (self PyBool) IntValue() int64 {
	if self.Value {
		return 1
	}
	return 0
}

func (self PyBool) BinaryAdd(b PyObject) PyObject {
	if !self.Value {
		return b
	}
	switch bb := b.(type) {
	case PyInt:
		return PyInt{Value: bb.Value + self.IntValue()}
	case PyFloat:
		return PyFloat{Value: bb.Value + float64(self.IntValue())}
	case PyBool:
		return PyInt{Value: bb.IntValue() + self.IntValue()}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}
