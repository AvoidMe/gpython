package builtin

type PyTuple struct {
	Value []PyObject
}

func (self *PyTuple) GetItem(index PyObject) (PyObject, error) {
	itemIndex := index.(*PyInt)                     // TODO: add error handling
	return self.Value[itemIndex.Value.Int64()], nil // TODO: add index checking
}

func (self *PyTuple) String() string {
	result := "("
	for index, arg := range self.Value {
		result += arg.Repr()
		if index != len(self.Value)-1 {
			result += ", "
		}
	}
	return result + ")"
}

func (self *PyTuple) Repr() string {
	return self.String()
}

func (self *PyTuple) Hash() (int64, error) {
	// TODO: this is bad hash implementation
	totalHash := int64(0)
	for _, item := range self.Value {
		hash, err := item.Hash()
		if err != nil {
			return 0, err
		}
		totalHash += hash
	}
	return totalHash, nil // TODO: move to TypeError
}

func (self *PyTuple) Equal(b PyObject) *PyBool {
	switch bb := b.(type) {
	case *PyTuple:
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

func (self *PyTuple) BinaryAdd(b PyObject) PyObject {
	bTuple, success := b.(*PyTuple)
	if !success {
		panic("Can't add list and non-list") // TODO: properly handle an error
	}
	result := []PyObject{}
	result = append(result, self.Value...)
	result = append(result, bTuple.Value...)
	return &PyTuple{Value: result}
}

func (self *PyTuple) Iter() PyIterator {
	return &PyTupleIterator{self, 0}
}
