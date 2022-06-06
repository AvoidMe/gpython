package gpythonlist

import (
	"fmt"
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

func (v *GpythonList) String() string {
	result := "["
	for index, arg := range v.List {
		innerList, success := arg.Value.(GpythonList)
		if success {
			result += innerList.String()
		} else {
			result += fmt.Sprintf("\"%s\"", arg.Value.(string))
		}
		if index != len(v.List)-1 {
			result += ", "
		}
	}
	return result + "]"
}
