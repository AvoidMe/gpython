package builtin

type PyString struct {
	Value string
}

func (self *PyString) String() string {
	return self.Value
}

func (self *PyString) Repr() string {
	return "'" + self.Value + "'"
}

func (self *PyString) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyString:
		if bb.Value == self.Value {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self *PyString) BinaryAdd(b PyObject) PyObject {
	bStr, success := b.(*PyString)
	if !success {
		panic("Can't add string and non-string") // TODO: properly handle an error
	}
	return &PyString{Value: self.Value + bStr.Value}
}
