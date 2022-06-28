package gpythonstring

type GpythonString struct {
	Str string
}

func (v GpythonString) String() string {
	return v.Str
}
