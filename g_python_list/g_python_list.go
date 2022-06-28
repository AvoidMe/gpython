package gpythonlist

import (
	"main/pyobject"
)

type GpythonList struct {
	List []pyobject.PyObject
}

func (v *GpythonList) Pop() pyobject.PyObject {
	value := v.List[len(v.List)-1]
	v.List = v.List[:len(v.List)-1]
	return value
}

func (v *GpythonList) PopN(n int) []pyobject.PyObject {
	items := []pyobject.PyObject{}
	items = append(items, v.List[len(v.List)-n:]...)
	v.List = v.List[:len(v.List)-n]
	return items
}

func (v *GpythonList) Append(value pyobject.PyObject) {
	v.List = append(v.List, value)
}

func (v *GpythonList) Extend(values []pyobject.PyObject) {
	v.List = append(v.List, values...)
}

func (v GpythonList) String() string {
	result := "["
	for index, arg := range v.List {
		result += arg.Repr()
		if index != len(v.List)-1 {
			result += ", "
		}
	}
	return result + "]"
}

func (v GpythonList) Repr() string {
	return v.String()
}
