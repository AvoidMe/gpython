package builtin

type PyObject interface {
	Repr() string
	String() string
	Hash() (int64, error)   // __hash__
	Equal(PyObject) *PyBool // __eq__
}

type PyBinaryAdd interface {
	PyObject
	BinaryAdd(PyObject) PyObject // __add__
}

type PyBinarySubstract interface {
	PyObject
	BinarySubstract(PyObject) PyObject // __sub__
}

type PySetItem interface {
	PyObject
	SetItem(PyObject, PyObject) error // __setitem__
}

type PyGetItem interface {
	PyObject
	GetItem(PyObject) (PyObject, error) // __getitem__
}

type PyIterable interface {
	PyObject
	Iter() PyIterator // __iter__
}

type PyIterator interface {
	PyIterable
	Next() (PyObject, error) // __next__
}
