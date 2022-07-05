package builtin

type NoneType struct{}

func (NoneType) Repr() string {
	return "None"
}

func (NoneType) String() string {
	return "None"
}

var None = NoneType{}
