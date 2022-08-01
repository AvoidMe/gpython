/*
	Right now golang doesn't have support map[] keying with complex types
	In ideal world I would prefer map[PyObject]PyObject as dictionary
	backend, but PyInt, PyTuple and PySet contains arrays in their implementation.

	This implementation isn't optimized from perfomance/memory perspective and can be
	done in much more efficient manner.
*/
package builtin

import (
	"errors"
)

type PyDictKeyValue struct {
	Key   PyObject
	Value PyObject
}

type PyDict struct {
	Value map[int64][]*PyDictKeyValue
}

func (self *PyDict) String() string {
	result := "{"
	indexI := 0
	for _, values := range self.Value {
		for indexJ, kv := range values {
			result += kv.Key.Repr() + ": " + kv.Value.Repr()
			if indexJ != len(values)-1 {
				result += ", "
			}
		}
		if indexI != len(self.Value)-1 {
			result += ", "
		}
		indexI++
	}
	return result + "}"
}

func (self *PyDict) Repr() string {
	return self.String()
}

func (self *PyDict) Hash() (int64, error) {
	return 0, errors.New("unhashable type: 'dict'") // TODO: move to TypeError
}

func (self *PyDict) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyDict:
		if len(bb.Value) != len(self.Value) {
			return PyFalse
		}
		for key, values := range self.Value {
			for _, kv := range values {
				result := PyFalse
				for _, bbKv := range bb.Value[key] {
					if bbKv.Key.Equal(kv.Key) == PyTrue && bbKv.Value.Equal(kv.Value) == PyTrue {
						result = PyTrue
						break
					}
				}
				if result == PyFalse {
					return PyFalse
				}
			}
		}
		return PyTrue
	default:
		return PyFalse
	}
}

func (self *PyDict) SetItem(key PyObject, value PyObject) error {
	if self.Value == nil {
		self.Value = map[int64][]*PyDictKeyValue{}
	}
	hash, err := key.Hash()
	if err != nil {
		return err
	}
	_, ok := self.Value[hash]
	if !ok {
		self.Value[hash] = make([]*PyDictKeyValue, 0)
	}
	items := self.Value[hash]
	for _, item := range items {
		if item.Key.Equal(key) == PyTrue {
			item.Value = value
			return nil
		}
	}
	self.Value[hash] = append(items, &PyDictKeyValue{Key: key, Value: value})
	return nil
}

func (self *PyDict) GetItem(key PyObject) (PyObject, error) {
	if self.Value == nil {
		// TODO: return error
	}
	hash, err := key.Hash()
	if err != nil {
		return nil, err
	}
	items, ok := self.Value[hash]
	if !ok {
		// TOOD: return error
	}
	for _, item := range items {
		if item.Key.Equal(key) == PyTrue {
			return item.Value, nil
		}
	}
	return nil, nil // TODO: return KeyError or something
}
