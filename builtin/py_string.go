package builtin

import "hash/maphash"

type PyString struct {
	Value []rune
}

func (self *PyString) String() string {
	return string(self.Value)
}

func (self *PyString) Repr() string {
	return "'" + string(self.Value) + "'"
}

func (self *PyString) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyString:
		if bb.String() == self.String() {
			return PyTrue
		}
		return PyFalse
	default:
		return PyFalse
	}
}

func (self *PyString) Hash() (uint64, error) {
	h := maphash.Hash{}
	h.SetSeed(*GetPyHashSeed())
	h.WriteString(self.String())
	return h.Sum64(), nil
}

func (self *PyString) BinaryAdd(b PyObject) PyObject {
	bStr, success := b.(*PyString)
	if !success {
		panic("Can't add string and non-string") // TODO: properly handle an error
	}
	newValue := []rune{}
	newValue = append(newValue, self.Value...)
	newValue = append(newValue, bStr.Value...)
	return &PyString{Value: newValue}
}

func NewPyString(s string) *PyString {
	return &PyString{Value: []rune(s)}
}
