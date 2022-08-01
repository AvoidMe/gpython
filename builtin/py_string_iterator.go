package builtin

import (
	"errors"
	"fmt"
	"unsafe"
)

type PyStringIterator struct {
	str   *PyString
	index int
}

func (self *PyStringIterator) Repr() string {
	return fmt.Sprintf("<string_iterator object at %p>", self)
}

func (self *PyStringIterator) String() string {
	return self.Repr()
}

func (self *PyStringIterator) Hash() (int64, error) {
	return *(*int64)(unsafe.Pointer(self)), nil
}

func (self *PyStringIterator) Equal(other PyObject) *PyBool {
	if self == other {
		return PyTrue
	}
	return PyFalse
}

func (self *PyStringIterator) Iter() PyIterator {
	return self
}

func (self *PyStringIterator) Next() (PyObject, error) {
	if self.index == len(self.str.Value) {
		return nil, errors.New("iter error") // TODO: move to real error
	}
	value := self.str.Value[self.index]
	self.index++
	return &PyString{Value: []rune{value}}, nil
}
