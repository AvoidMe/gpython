package builtin

type PyObject interface {
	Repr() string
	String() string
	Hash() (int64, error)
	Equal(PyObject) *PyBool // __eq__ method
}

type PyBinaryAdd interface {
	PyObject
	BinaryAdd(PyObject) PyObject // __add__ method
}

type PyBinarySubstract interface {
	PyObject
	BinarySubstract(PyObject) PyObject // __sub__ method
}

type PySetItem interface {
	PyObject
	SetItem(PyObject, PyObject) error
}

type PyGetItem interface {
	PyObject
	GetItem(PyObject) (PyObject, error)
}

type PyIterable interface {
	PyObject
	Iter() PyIterator // __iter__ method
}

type PyIterator interface {
	PyIterable
	Next() (PyObject, error) // __next__ method
}
