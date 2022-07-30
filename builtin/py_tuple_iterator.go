package builtin

import (
	"errors"
	"fmt"
	"unsafe"
)

type PyTupleIterator struct {
	tuple *PyTuple
	index int
}

func (self *PyTupleIterator) Repr() string {
	return fmt.Sprintf("<tuple_iterator object at %p>", self)
}

func (self *PyTupleIterator) String() string {
	return self.Repr()
}

func (self *PyTupleIterator) Hash() (uint64, error) {
	return *(*uint64)(unsafe.Pointer(self)), nil
}

func (self *PyTupleIterator) Equal(other PyObject) *PyBool {
	if self == other {
		return PyTrue
	}
	return PyFalse
}

func (self *PyTupleIterator) Iter() PyIterator {
	return self
}

func (self *PyTupleIterator) Next() (PyObject, error) {
	if self.index == len(self.tuple.Value) {
		return nil, errors.New("iter error") // TODO: move to real error
	}
	value := self.tuple.Value[self.index]
	self.index++
	return value, nil
}
