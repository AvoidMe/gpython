package builtin

type PyObject interface {
	Repr() string
	String() string
	// Equal() bool // __eq__ method
}

// Interface, which indicates what object has __add__ method
type PyBinaryAdd interface {
	PyObject
	BinaryAdd(PyObject) PyObject
}
