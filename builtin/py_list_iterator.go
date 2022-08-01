package builtin

import (
	"errors"
	"fmt"
	"unsafe"
)

type PyListIterator struct {
	list  *PyList
	index int
}

func (self *PyListIterator) Repr() string {
	return fmt.Sprintf("<list_iterator object at %p>", self)
}

func (self *PyListIterator) String() string {
	return self.Repr()
}

func (self *PyListIterator) Hash() (int64, error) {
	return *(*int64)(unsafe.Pointer(self)), nil
}

func (self *PyListIterator) Equal(other PyObject) *PyBool {
	if self == other {
		return PyTrue
	}
	return PyFalse
}

func (self *PyListIterator) Iter() PyIterator {
	return self
}

func (self *PyListIterator) Next() (PyObject, error) {
	// TODO: Add error if list size changed during iteration
	if self.index == len(self.list.Value) {
		return nil, errors.New("iter error") // TODO: move to real error
	}
	value := self.list.Value[self.index]
	self.index++
	return value, nil
}
