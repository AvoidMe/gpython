package builtin

import (
	"testing"

	"github.com/AvoidMe/gpython/builtin"
)

func TestEmptyDict(t *testing.T) {
	d := &builtin.PyDict{}
	if d.String() != "{}" {
		t.Errorf("Empty dict unexpected value: %v", d.String())
	}
}

func TestDictSetItem(t *testing.T) {
	// Generic cases with every hashable python type as key
	cases := [][]builtin.PyObject{
		{builtin.NewPyString("a"), builtin.NewPyInt(1)},
		{builtin.NewPyString("b"), builtin.NewPyInt(2)},
		{builtin.NewPyString("c"), builtin.NewPyInt(3)},
		{builtin.NewPyInt(100), builtin.NewPyInt(4)},
		{builtin.NewPyInt(101), builtin.NewPyInt(5)},
		{builtin.NewPyInt(102), builtin.NewPyInt(6)},
		{&builtin.PyFloat{Value: 9997.5}, builtin.NewPyInt(7)},
		{&builtin.PyFloat{Value: 9998.123}, builtin.NewPyInt(8)},
		{&builtin.PyFloat{Value: 9999.456}, builtin.NewPyInt(9)},
		{builtin.PyNone, builtin.NewPyInt(10)},
		{builtin.PyTrue, builtin.NewPyInt(11)},
		{builtin.PyFalse, builtin.NewPyInt(12)},
	}
	_genericDictTest(cases, cases, t)
	// Override value cases
	input := [][]builtin.PyObject{
		{builtin.NewPyString("a"), builtin.NewPyInt(1)},
		{builtin.NewPyString("a"), builtin.NewPyInt(123)},
		{builtin.NewPyInt(100), builtin.NewPyString("hello")},
		{builtin.NewPyInt(100), builtin.NewPyString("world")},
		{&builtin.PyFloat{Value: 9997.5}, builtin.NewPyInt(7)},
		{&builtin.PyFloat{Value: 9997.5}, builtin.PyNone},
		{builtin.PyNone, builtin.NewPyInt(777)},
		{builtin.PyNone, builtin.NewPyInt(888)},
		{builtin.PyTrue, builtin.NewPyInt(999)},
		{builtin.PyTrue, builtin.NewPyInt(111)},
		{builtin.PyFalse, builtin.NewPyInt(222)},
		{builtin.PyFalse, builtin.NewPyInt(333)},
	}
	expected := [][]builtin.PyObject{
		{builtin.NewPyString("a"), builtin.NewPyInt(123)},
		{builtin.NewPyInt(100), builtin.NewPyString("world")},
		{&builtin.PyFloat{Value: 9997.5}, builtin.PyNone},
		{builtin.PyNone, builtin.NewPyInt(888)},
		{builtin.PyTrue, builtin.NewPyInt(111)},
		{builtin.PyFalse, builtin.NewPyInt(333)},
	}
	_genericDictTest(input, expected, t)

	// Simple tuple case
	input = [][]builtin.PyObject{
		{
			&builtin.PyTuple{
				Value: []builtin.PyObject{
					builtin.NewPyInt(1),
					builtin.NewPyInt(2),
					builtin.NewPyInt(3),
				},
			},
			builtin.NewPyString("Hello pytuple!"),
		},
	}
	expected = [][]builtin.PyObject{
		{
			&builtin.PyTuple{
				Value: []builtin.PyObject{
					builtin.NewPyInt(1),
					builtin.NewPyInt(2),
					builtin.NewPyInt(3),
				},
			},
			builtin.NewPyString("Hello pytuple!"),
		},
	}
	_genericDictTest(input, expected, t)

	// From this point we're testing some tricky cpython cases

	// float(15.0) should be the same key as int(15)
	input = [][]builtin.PyObject{
		{&builtin.PyFloat{Value: 15.0}, builtin.NewPyString("hello")},
	}
	expected = [][]builtin.PyObject{
		{builtin.NewPyInt(15), builtin.NewPyString("hello")},
		{&builtin.PyFloat{Value: 15.0}, builtin.NewPyString("hello")},
	}
	_genericDictTest(input, expected, t)

	// int(15) should be the same key as float(15.0)
	input = [][]builtin.PyObject{
		{builtin.NewPyInt(15), builtin.NewPyString("hello")},
	}
	expected = [][]builtin.PyObject{
		{builtin.NewPyInt(15), builtin.NewPyString("hello")},
		{&builtin.PyFloat{Value: 15.0}, builtin.NewPyString("hello")},
	}
	_genericDictTest(input, expected, t)

	// True should be the same key as int(1) and float(1.0)
	input = [][]builtin.PyObject{
		{builtin.PyTrue, builtin.NewPyString("hello true")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyTrue, builtin.NewPyString("hello true")},
		{builtin.NewPyInt(1), builtin.NewPyString("hello true")},
		{&builtin.PyFloat{Value: 1}, builtin.NewPyString("hello true")},
	}
	_genericDictTest(input, expected, t)

	// False should be the same key as int(0) and float(0.0)
	input = [][]builtin.PyObject{
		{builtin.PyFalse, builtin.NewPyString("hello false")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyFalse, builtin.NewPyString("hello false")},
		{builtin.NewPyInt(0), builtin.NewPyString("hello false")},
		{&builtin.PyFloat{Value: 0}, builtin.NewPyString("hello false")},
	}
	_genericDictTest(input, expected, t)

	// int(1) should be the same key as True and float(1)
	input = [][]builtin.PyObject{
		{builtin.NewPyInt(1), builtin.NewPyString("hello int(1)")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyTrue, builtin.NewPyString("hello int(1)")},
		{builtin.NewPyInt(1), builtin.NewPyString("hello int(1)")},
		{&builtin.PyFloat{Value: 1}, builtin.NewPyString("hello int(1)")},
	}
	_genericDictTest(input, expected, t)

	// float(1) should be the same key as True and int(1)
	input = [][]builtin.PyObject{
		{&builtin.PyFloat{Value: 1}, builtin.NewPyString("hello float(1)")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyTrue, builtin.NewPyString("hello float(1)")},
		{builtin.NewPyInt(1), builtin.NewPyString("hello float(1)")},
		{&builtin.PyFloat{Value: 1}, builtin.NewPyString("hello float(1)")},
	}
	_genericDictTest(input, expected, t)

	// int(0) should be the same key as False and float(0)
	input = [][]builtin.PyObject{
		{builtin.NewPyInt(0), builtin.NewPyString("hello int(0)")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyFalse, builtin.NewPyString("hello int(0)")},
		{builtin.NewPyInt(0), builtin.NewPyString("hello int(0)")},
		{&builtin.PyFloat{Value: 0}, builtin.NewPyString("hello int(0)")},
	}
	_genericDictTest(input, expected, t)

	// float(0) should be the same key as False and int(0)
	input = [][]builtin.PyObject{
		{&builtin.PyFloat{Value: 0}, builtin.NewPyString("hello float(0)")},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyFalse, builtin.NewPyString("hello float(0)")},
		{builtin.NewPyInt(0), builtin.NewPyString("hello float(0)")},
		{&builtin.PyFloat{Value: 0}, builtin.NewPyString("hello float(0)")},
	}
	_genericDictTest(input, expected, t)
	// TODO: add error tests
	// TODO: add functions as dict keys/values
}

func _genericDictTest(input [][]builtin.PyObject, expectedList [][]builtin.PyObject, t *testing.T) {
	result := &builtin.PyDict{}
	for i := 0; i < len(input); i++ {
		key, value := input[i][0], input[i][1]
		result.SetItem(key, value)
	}
	for i := 0; i < len(expectedList); i++ {
		key, expected := expectedList[i][0], expectedList[i][1]
		answer, _ := result.GetItem(key) // TODO: add error checking
		if answer.Equal(expected) != builtin.PyTrue {
			t.Errorf("After insert dict has unexpected key/value: %v: %v, %v", key, expected, answer)
		}
	}

}
