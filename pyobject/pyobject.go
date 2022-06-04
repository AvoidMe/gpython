package pyobject

type PyObject struct {
	Value    string
	Function interface{}
}

var None = PyObject{}
