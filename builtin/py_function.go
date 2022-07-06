package builtin

type PyFunction struct {
	Callable    func(PyObject, PyObject) PyObject
	StringValue string
	ReprValue   string
}

func (self PyFunction) String() string {
	return self.StringValue
}

func (self PyFunction) Repr() string {
	return self.ReprValue
}

func (self PyFunction) Equal(b PyObject) PyBool {
	switch bb := b.(type) {
	case PyFunction:
		if &bb == &self {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}
