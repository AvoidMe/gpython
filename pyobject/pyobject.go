package pyobject

type PyObject struct {
	Value interface{}
}

var None = PyObject{}
