package pyobject

type PyObject interface {
	Repr() string
	String() string
}

// Interface, which indicates what object has __add__ method
type PyBinaryAdd interface {
	PyObject
	BinaryAdd(PyObject) PyObject
}

type NoneType struct{}

func (NoneType) Repr() string {
	return "None"
}

func (NoneType) String() string {
	return "None"
}

var None = NoneType{}
