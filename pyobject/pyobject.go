package pyobject

type PyObject struct {
	Value    string
	Tuple    []PyObject
	Function interface{}
}

var None = PyObject{}
