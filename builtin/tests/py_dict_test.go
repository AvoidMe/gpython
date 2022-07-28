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
		{&builtin.PyString{Value: "a"}, &builtin.PyInt{Value: 1}},
		{&builtin.PyString{Value: "b"}, &builtin.PyInt{Value: 2}},
		{&builtin.PyString{Value: "c"}, &builtin.PyInt{Value: 3}},
		{&builtin.PyInt{Value: 100}, &builtin.PyInt{Value: 4}},
		{&builtin.PyInt{Value: 101}, &builtin.PyInt{Value: 5}},
		{&builtin.PyInt{Value: 102}, &builtin.PyInt{Value: 6}},
		{&builtin.PyFloat{Value: 9997.5}, &builtin.PyInt{Value: 7}},
		{&builtin.PyFloat{Value: 9998.123}, &builtin.PyInt{Value: 8}},
		{&builtin.PyFloat{Value: 9999.456}, &builtin.PyInt{Value: 9}},
		{builtin.PyNone, &builtin.PyInt{Value: 10}},
		{builtin.PyTrue, &builtin.PyInt{Value: 11}},
		{builtin.PyFalse, &builtin.PyInt{Value: 12}},
	}
	_genericDictTest(cases, cases, t)
	// Override value cases
	input := [][]builtin.PyObject{
		{&builtin.PyString{Value: "a"}, &builtin.PyInt{Value: 1}},
		{&builtin.PyString{Value: "a"}, &builtin.PyInt{Value: 123}},
		{&builtin.PyInt{Value: 100}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyInt{Value: 100}, &builtin.PyString{Value: "world"}},
		{&builtin.PyFloat{Value: 9997.5}, &builtin.PyInt{Value: 7}},
		{&builtin.PyFloat{Value: 9997.5}, builtin.PyNone},
		{builtin.PyNone, &builtin.PyInt{Value: 777}},
		{builtin.PyNone, &builtin.PyInt{Value: 888}},
		{builtin.PyTrue, &builtin.PyInt{Value: 999}},
		{builtin.PyTrue, &builtin.PyInt{Value: 111}},
		{builtin.PyFalse, &builtin.PyInt{Value: 222}},
		{builtin.PyFalse, &builtin.PyInt{Value: 333}},
	}
	expected := [][]builtin.PyObject{
		{&builtin.PyString{Value: "a"}, &builtin.PyInt{Value: 123}},
		{&builtin.PyInt{Value: 100}, &builtin.PyString{Value: "world"}},
		{&builtin.PyFloat{Value: 9997.5}, builtin.PyNone},
		{builtin.PyNone, &builtin.PyInt{Value: 888}},
		{builtin.PyTrue, &builtin.PyInt{Value: 111}},
		{builtin.PyFalse, &builtin.PyInt{Value: 333}},
	}
	_genericDictTest(input, expected, t)

	// From this point we're testing some tricky cpython cases

	// float(15.0) should be the same key as int(15)
	input = [][]builtin.PyObject{
		{&builtin.PyFloat{Value: 15.0}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.123}, &builtin.PyString{Value: "world"}},
	}
	expected = [][]builtin.PyObject{
		{&builtin.PyInt{Value: 15}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.0}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.123}, &builtin.PyString{Value: "world"}},
	}
	_genericDictTest(input, expected, t)

	// int(15) should be the same key as float(15.0)
	input = [][]builtin.PyObject{
		{&builtin.PyInt{Value: 15}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.123}, &builtin.PyString{Value: "world"}},
	}
	expected = [][]builtin.PyObject{
		{&builtin.PyInt{Value: 15}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.0}, &builtin.PyString{Value: "hello"}},
		{&builtin.PyFloat{Value: 15.123}, &builtin.PyString{Value: "world"}},
	}
	_genericDictTest(input, expected, t)

	// True should be the same key as int(1) and float(1.0)
	input = [][]builtin.PyObject{
		{builtin.PyTrue, &builtin.PyString{Value: "hello true"}},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyTrue, &builtin.PyString{Value: "hello true"}},
		{&builtin.PyInt{Value: 1}, &builtin.PyString{Value: "hello true"}},
		{&builtin.PyFloat{Value: 1}, &builtin.PyString{Value: "hello true"}},
	}
	_genericDictTest(input, expected, t)

	// False should be the same key as int(0) and float(0.0)
	input = [][]builtin.PyObject{
		{builtin.PyFalse, &builtin.PyString{Value: "hello false"}},
	}
	expected = [][]builtin.PyObject{
		{builtin.PyFalse, &builtin.PyString{Value: "hello false"}},
		{&builtin.PyInt{Value: 0}, &builtin.PyString{Value: "hello false"}},
		{&builtin.PyFloat{Value: 0}, &builtin.PyString{Value: "hello false"}},
	}
	_genericDictTest(input, expected, t)

	// TODO: add error tests
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
