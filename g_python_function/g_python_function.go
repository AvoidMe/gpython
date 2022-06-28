package gpythonfunction

import (
	"main/pyobject"
)

type GPythonFunction struct {
	Callable    func(pyobject.PyObject, pyobject.PyObject) pyobject.PyObject
	StringValue string
	ReprValue   string
}

func (f GPythonFunction) String() string {
	return f.StringValue
}

func (f GPythonFunction) Repr() string {
	return f.ReprValue
}
