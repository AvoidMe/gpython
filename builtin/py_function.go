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
