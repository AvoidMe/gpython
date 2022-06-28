package pyobject

type PyObject interface {
	Repr() string
	String() string
}

type NoneType struct{}

func (NoneType) Repr() string {
	return "None"
}

func (NoneType) String() string {
	return "None"
}

var None = NoneType{}
