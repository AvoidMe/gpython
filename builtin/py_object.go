package builtin

type PyObject interface {
	Repr() string
	String() string
	Equal(PyObject) *PyBool // __eq__ method
}

type PyBinaryAdd interface {
	PyObject
	BinaryAdd(PyObject) PyObject // __add__ method
}
