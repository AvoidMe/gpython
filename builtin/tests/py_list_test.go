package builtin

import (
	"testing"

	"github.com/AvoidMe/gpython/builtin"
)

// TODO: add negative test cases
// TODO: add functions as list value
func TestEmptyList(t *testing.T) {
	l := &builtin.PyList{}
	if l.String() != "[]" {
		t.Errorf("Empty list unexpected value: %v", l.String())
	}
}

type listCase struct {
	Index int64
	Value builtin.PyObject
}

func TestListSetItem(t *testing.T) {
	input := []builtin.PyObject{
		builtin.NewPyInt(0),
		builtin.NewPyInt(1),
		builtin.NewPyInt(2),
		builtin.NewPyInt(3),
		builtin.NewPyInt(4),
		builtin.NewPyInt(5),
		builtin.PyNone,
		builtin.PyNone,
		builtin.PyNone,
		builtin.PyNone,
		&builtin.PyFloat{Value: 11.111},
		&builtin.PyFloat{Value: 22.222},
		&builtin.PyFloat{Value: 33.333},
		&builtin.PyFloat{Value: 44.444},
		builtin.PyTrue,
		builtin.PyFalse,
		builtin.PyTrue,
		builtin.PyFalse,
	}
	setCases := []listCase{
		{1, builtin.NewPyString("first")},
		{5, builtin.NewPyString("last")},
		{6, builtin.PyTrue},
		{11, builtin.NewPyString("Miss me?")},
	}
	expected := []builtin.PyObject{
		builtin.NewPyInt(0),
		builtin.NewPyString("first"),
		builtin.NewPyInt(2),
		builtin.NewPyInt(3),
		builtin.NewPyInt(4),
		builtin.NewPyString("last"),
		builtin.PyTrue,
		builtin.PyNone,
		builtin.PyNone,
		builtin.PyNone,
		&builtin.PyFloat{Value: 11.111},
		builtin.NewPyString("Miss me?"),
		&builtin.PyFloat{Value: 33.333},
		&builtin.PyFloat{Value: 44.444},
		builtin.PyTrue,
		builtin.PyFalse,
		builtin.PyTrue,
		builtin.PyFalse,
	}
	_genericListTest(input, setCases, expected, t)
}

func _genericListTest(input []builtin.PyObject, setCases []listCase, expectedList []builtin.PyObject, t *testing.T) {
	result := &builtin.PyList{}
	for i := 0; i < len(input); i++ {
		result.Append(input[i])
	}
	for _, setCase := range setCases {
		result.SetItem(builtin.NewPyInt(setCase.Index), setCase.Value)
	}
	for i := 0; i < len(expectedList); i++ {
		answer, _ := result.GetItem(builtin.NewPyInt(int64(i))) // TODO: add error checking
		expected := expectedList[i]
		if answer.Equal(expected) != builtin.PyTrue {
			t.Errorf("After insert list has unexpected value: %v: %v, %v", i, expected, answer)
		}
	}

}
