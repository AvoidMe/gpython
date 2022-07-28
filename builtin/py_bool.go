package builtin

type PyBool struct {
	Value bool
}

func (self *PyBool) String() string {
	if !self.Value {
		return "False"
	}
	return "True"
}

func (self *PyBool) Repr() string {
	return self.String()
}

func (self *PyBool) IntValue() int64 {
	if self.Value {
		return 1
	}
	return 0
}

func (self *PyBool) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyFloat:
		if bb.Value == float64(self.IntValue()) {
			return PyTrue
		}
		return PyFalse
	case *PyInt:
		if bb.Value == self.IntValue() {
			return PyTrue
		}
		return PyFalse
	case *PyBool:
		if bb.Value == self.Value {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self *PyBool) BinaryAdd(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		return &PyInt{Value: bb.Value + self.IntValue()}
	case *PyFloat:
		return &PyFloat{Value: bb.Value + float64(self.IntValue())}
	case *PyBool:
		return &PyInt{Value: bb.IntValue() + self.IntValue()}
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}

func (self *PyBool) BinarySubstract(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		return &PyInt{Value: self.IntValue() - bb.Value}
	case *PyFloat:
		return &PyFloat{Value: float64(self.IntValue()) - bb.Value}
	case *PyBool:
		return &PyInt{Value: self.IntValue() - bb.IntValue()}
	default:
		panic("Can't substract number and non-number") // TODO: properly handle an error
	}
}
