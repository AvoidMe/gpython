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

func (self *PyBool) IntValue() *PyInt {
	if self.Value {
		return NewPyInt(1)
	}
	return NewPyInt(0)
}

func (self *PyBool) Hash() (int64, error) {
	return self.IntValue().Hash()
}

func (self *PyBool) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyFloat:
		if bb.Value == float64(self.IntValue().Int64()) {
			return PyTrue
		}
		return PyFalse
	case *PyInt:
		return self.IntValue().Equal(bb)
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
		return self.IntValue().BinaryAdd(bb)
	case *PyFloat:
		return self.IntValue().BinaryAdd(bb)
	case *PyBool:
		return self.IntValue().BinaryAdd(bb.IntValue())
	default:
		panic("Can't add number and non-number") // TODO: properly handle an error
	}
}

func (self *PyBool) BinarySubstract(b PyObject) PyObject {
	switch bb := b.(type) {
	case *PyInt:
		return self.IntValue().BinarySubstract(bb)
	case *PyFloat:
		return &PyFloat{Value: float64(self.IntValue().Int64()) - bb.Value}
	case *PyBool:
		return self.IntValue().BinarySubstract(bb.IntValue())
	default:
		panic("Can't substract number and non-number") // TODO: properly handle an error
	}
}
