package builtin

import "errors"

type PyList struct {
	Value []PyObject
}

func (self *PyList) Pop() PyObject {
	value := self.Value[len(self.Value)-1]
	self.Value = self.Value[:len(self.Value)-1]
	return value
}

func (self *PyList) SetItem(index PyObject, value PyObject) error {
	itemIndex := index.(*PyInt)         // TODO: add error handling
	self.Value[itemIndex.Value] = value // TODO: add index checking
	return nil
}

func (self *PyList) GetItem(index PyObject) (PyObject, error) {
	itemIndex := index.(*PyInt)             // TODO: add error handling
	return self.Value[itemIndex.Value], nil // TODO: add index checking
}

func (self *PyList) PopN(n int) []PyObject {
	items := []PyObject{}
	items = append(items, self.Value[len(self.Value)-n:]...)
	self.Value = self.Value[:len(self.Value)-n]
	return items
}

func (self *PyList) Append(value PyObject) {
	self.Value = append(self.Value, value)
}

func (self *PyList) Extend(values []PyObject) {
	self.Value = append(self.Value, values...)
}

func (self *PyList) String() string {
	result := "["
	for index, arg := range self.Value {
		result += arg.Repr()
		if index != len(self.Value)-1 {
			result += ", "
		}
	}
	return result + "]"
}

func (self *PyList) Repr() string {
	return self.String()
}

func (self *PyList) Hash() (uint64, error) {
	return 0, errors.New("unhashable type: 'list'") // TODO: move to TypeError
}

func (self *PyList) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyList:
		if len(bb.Value) != len(self.Value) {
			return PyFalse
		}
		for index := range self.Value {
			if self.Value[index].Equal(bb.Value[index]) == PyFalse {
				return PyFalse
			}
		}
		return PyTrue
	default:
		return PyFalse
	}
}

func (self *PyList) BinaryAdd(b PyObject) PyObject {
	bList, success := b.(*PyList)
	if !success {
		panic("Can't add list and non-list") // TODO: properly handle an error
	}
	result := []PyObject{}
	result = append(result, self.Value...)
	result = append(result, bList.Value...)
	return &PyList{Value: result}
}
