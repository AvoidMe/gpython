package eval

import (
	"reflect"
	"testing"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
)

type binaryTestCase struct {
	Left     builtin.PyObject
	Right    builtin.PyObject
	Expected builtin.PyObject
}

func TestEmptyInput(t *testing.T) {
	input := []opcode.Instruction{}
	EvalInstructions(input)
}

func TestReturnValue(t *testing.T) {
	var expected int64 = 42
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyInt{Value: expected},
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(*builtin.PyInt)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if answer.Value != 42 {
		t.Errorf("Returns incorrect value: %v", answer)
	}
}

func TestListExtend(t *testing.T) {
	firstList := builtin.PyList{
		Value: []builtin.PyObject{
			&builtin.PyString{Value: "First"},
			&builtin.PyString{Value: "Second"},
		},
	}
	secondList := builtin.PyList{
		Value: []builtin.PyObject{
			&builtin.PyString{Value: "Third"},
			&builtin.PyString{Value: "Fourth"},
		},
	}
	expected := builtin.PyList{
		Value: []builtin.PyObject{
			&builtin.PyString{Value: "First"},
			&builtin.PyString{Value: "Second"},
			&builtin.PyString{Value: "Third"},
			&builtin.PyString{Value: "Fourth"},
		},
	}
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &firstList,
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &secondList,
		},
		{
			Opcode: opcode.LIST_EXTEND,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(input)
	if expected.Equal(result) == builtin.PyFalse {
		t.Errorf("Got: %v, expected: %v", result, expected)
	}
}

func TestBuildList(t *testing.T) {
	expected := builtin.PyList{
		Value: []builtin.PyObject{
			&builtin.PyString{Value: "First"},
			&builtin.PyString{Value: "Second"},
			&builtin.PyString{Value: "Third"},
			&builtin.PyString{Value: "Fourth"},
		},
	}
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyString{Value: "First"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyString{Value: "Second"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyString{Value: "Third"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyString{Value: "Fourth"},
		},
		{
			Opcode: opcode.BUILD_LIST,
			Arg:    4,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(input)
	if expected.Equal(result) == builtin.PyFalse {
		t.Errorf("Got: %v, expected: %v", result, expected)
	}
}

func TestStoreLocalName(t *testing.T) {
	expected := 42.05
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyFloat{Value: 42.05},
		},
		{
			Opcode: opcode.STORE_NAME,
			Arg:    0,
			Args:   &builtin.PyString{Value: "x"},
		},
		{
			Opcode: opcode.LOAD_NAME,
			Arg:    0,
			Args:   &builtin.PyString{Value: "x"},
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(*builtin.PyFloat)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if answer.Value != expected {
		t.Errorf("Returns incorect value: %v", answer)
	}
}

func TestPopTop(t *testing.T) {
	expected := 42.05
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyFloat{Value: 42.05},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   &builtin.PyInt{Value: 42},
		},
		{
			Opcode: opcode.POP_TOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(*builtin.PyFloat)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if answer.Value != expected {
		t.Errorf("Returns incorect value: %v", answer)
	}
}

func TestBinaryAdd(t *testing.T) {
	cases := []binaryTestCase{
		{
			&builtin.PyString{Value: "Hello"},
			&builtin.PyString{Value: ", world!"},
			&builtin.PyString{Value: "Hello, world!"},
		},
		{
			&builtin.PyInt{Value: 10},
			&builtin.PyInt{Value: 32},
			&builtin.PyInt{Value: 42},
		},
		{
			&builtin.PyInt{Value: 10},
			&builtin.PyFloat{Value: 32.54321},
			&builtin.PyFloat{Value: 42.54321},
		},
		{
			&builtin.PyFloat{Value: 10.54321},
			&builtin.PyInt{Value: 32},
			&builtin.PyFloat{Value: 42.54321},
		},
		{
			&builtin.PyFloat{Value: 10.54321},
			&builtin.PyFloat{Value: 32.12345},
			&builtin.PyFloat{Value: 42.66666},
		},
		{
			builtin.PyTrue,
			builtin.PyFalse,
			&builtin.PyInt{Value: 1},
		},
		{
			builtin.PyFalse,
			builtin.PyFalse,
			&builtin.PyInt{Value: 0},
		},
		{
			builtin.PyTrue,
			builtin.PyTrue,
			&builtin.PyInt{Value: 2},
		},
		{
			builtin.PyTrue,
			&builtin.PyInt{Value: 41},
			&builtin.PyInt{Value: 42},
		},
		{
			builtin.PyFalse,
			&builtin.PyInt{Value: 41},
			&builtin.PyInt{Value: 41},
		},
		{
			builtin.PyFalse,
			&builtin.PyFloat{Value: 41.5},
			&builtin.PyFloat{Value: 41.5},
		},
		{
			&builtin.PyFloat{Value: 41.5},
			builtin.PyFalse,
			&builtin.PyFloat{Value: 41.5},
		},
		{
			builtin.PyTrue,
			&builtin.PyFloat{Value: 41.5},
			&builtin.PyFloat{Value: 42.5},
		},
		{
			&builtin.PyFloat{Value: 41.5},
			builtin.PyTrue,
			&builtin.PyFloat{Value: 42.5},
		},
		{
			&builtin.PyInt{Value: 41},
			builtin.PyFalse,
			&builtin.PyInt{Value: 41},
		},
		{
			&builtin.PyInt{Value: 41},
			builtin.PyTrue,
			&builtin.PyInt{Value: 42},
		},
	}
	for _, testCase := range cases {
		_testBinaryAdd(t, testCase.Left, testCase.Right, testCase.Expected)
	}
}

func _testBinaryAdd(t *testing.T, a builtin.PyObject, b builtin.PyObject, expected builtin.PyObject) {
	instructions := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   a,
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   b,
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.PyEq,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(instructions)
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)
	if expectedType != resultType {
		t.Errorf("Returns incorect type, a: %v, b: %v, expected: %v, answer: %v", a, b, expectedType, resultType)
	}
	if result.Equal(expected) == builtin.PyFalse {
		t.Errorf("Returns incorect value, a: %v, b: %v, expected: %v, answer: %v", a, b, expected, result)
	}
}

func TestCompareOpEqual(t *testing.T) {
	cases := []binaryTestCase{
		// None
		{builtin.PyNone, builtin.PyNone, builtin.PyTrue},
		{builtin.PyNone, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// bool == bool
		{builtin.PyTrue, builtin.PyTrue, builtin.PyTrue},
		{builtin.PyTrue, builtin.PyFalse, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyTrue, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyFalse, builtin.PyTrue},
		// bool == unexpected_type
		{builtin.PyTrue, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// bool == int
		{builtin.PyTrue, &builtin.PyInt{Value: 1}, builtin.PyTrue},
		{builtin.PyTrue, &builtin.PyInt{Value: 0}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyInt{Value: 1}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyInt{Value: 0}, builtin.PyTrue},
		// bool == float
		{builtin.PyTrue, &builtin.PyFloat{Value: 1.0}, builtin.PyTrue},
		{builtin.PyTrue, &builtin.PyFloat{Value: 0.0}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyFloat{Value: 1.0}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyFloat{Value: 0.0}, builtin.PyTrue},
		// int == bool
		{&builtin.PyInt{Value: 1}, builtin.PyTrue, builtin.PyTrue},
		{&builtin.PyInt{Value: 0}, builtin.PyTrue, builtin.PyFalse},
		{&builtin.PyInt{Value: 1}, builtin.PyFalse, builtin.PyFalse},
		{&builtin.PyInt{Value: 0}, builtin.PyFalse, builtin.PyTrue},
		// int == int
		{&builtin.PyInt{Value: 42}, &builtin.PyInt{Value: 42}, builtin.PyTrue},
		{&builtin.PyInt{Value: 42}, &builtin.PyInt{Value: 444}, builtin.PyFalse},
		// int == unexpected_type
		{&builtin.PyInt{Value: 42}, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// int == float
		{&builtin.PyInt{Value: 42}, &builtin.PyFloat{Value: 42.0}, builtin.PyTrue},
		{&builtin.PyInt{Value: 42}, &builtin.PyFloat{Value: 42.5}, builtin.PyFalse},
		// float == int
		{&builtin.PyFloat{Value: 42.0}, &builtin.PyInt{Value: 42}, builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyInt{Value: 42}, builtin.PyFalse},
		// float == float
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.5}, builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.555}, builtin.PyFalse},
		// float == bool
		{&builtin.PyFloat{Value: 1.0}, builtin.PyTrue, builtin.PyTrue},
		{&builtin.PyFloat{Value: 0.0}, builtin.PyTrue, builtin.PyFalse},
		{&builtin.PyFloat{Value: 1.0}, builtin.PyFalse, builtin.PyFalse},
		{&builtin.PyFloat{Value: 0.0}, builtin.PyFalse, builtin.PyTrue},
		// float == unextected_type
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyString{Value: "testy"}, builtin.PyFalse},

		// function == function
		{builtin.Builtin["print"], builtin.Builtin["print"], builtin.PyTrue},
		// TODO: function == function -> PyFals

		// function == unexpected_type
		{builtin.Builtin["print"], &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// list == unexpected_type
		{
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			&builtin.PyString{Value: "testy"},
			builtin.PyFalse,
		},
		// list == list
		{
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy_test_test"}},
			},
			builtin.PyFalse,
		},
		{
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			builtin.PyTrue,
		},
		// list with nested lists
		{
			&builtin.PyList{
				Value: []builtin.PyObject{
					&builtin.PyString{Value: "testy"},
					&builtin.PyList{
						Value: []builtin.PyObject{
							&builtin.PyString{Value: "testy"},
						},
					},
				},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			builtin.PyFalse,
		},
		{
			&builtin.PyList{
				Value: []builtin.PyObject{
					&builtin.PyString{Value: "testy"},
					&builtin.PyList{
						Value: []builtin.PyObject{
							&builtin.PyString{Value: "testy"},
						},
					},
				},
			},
			&builtin.PyList{
				Value: []builtin.PyObject{
					&builtin.PyString{Value: "testy"},
					&builtin.PyList{
						Value: []builtin.PyObject{
							&builtin.PyString{Value: "testy"},
						},
					},
				},
			},
			builtin.PyTrue,
		},
	}
	for _, testCase := range cases {
		_testCompareOpEqual(t, testCase.Left, testCase.Right, testCase.Expected.(*builtin.PyBool))
	}
}

func _testCompareOpEqual(t *testing.T, a builtin.PyObject, b builtin.PyObject, expected *builtin.PyBool) {
	instructions := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   a,
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   b,
		},
		{
			Opcode: opcode.COMPARE_OP,
			Arg:    0,
			Args:   builtin.PyEq,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(instructions)
	answer, ok := result.(*builtin.PyBool)
	if !ok {
		t.Errorf("Returns wrong type: %v", answer)
	}
	if answer.Value != expected.Value {
		t.Errorf("Returns incorect value, a: %v, b: %v, expected: %v, answer: %v", a, b, expected, answer)
	}
}

func TestIsOp(t *testing.T) {
	listItem := &builtin.PyList{Value: []builtin.PyObject{
		&builtin.PyString{Value: "testy"},
	}}
	intItem := &builtin.PyInt{Value: 42}
	floatItem := &builtin.PyFloat{Value: 42.5}
	cases := []binaryTestCase{
		// None
		{builtin.PyNone, builtin.PyNone, builtin.PyTrue},
		{builtin.PyNone, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// bool is bool
		{builtin.PyTrue, builtin.PyTrue, builtin.PyTrue},
		{builtin.PyTrue, builtin.PyFalse, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyTrue, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyFalse, builtin.PyTrue},
		// bool is unexpected_type
		{builtin.PyTrue, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// int is int
		{intItem, intItem, builtin.PyTrue},
		{&builtin.PyInt{Value: 9999999}, intItem, builtin.PyFalse},
		// int is unexpected_type
		{&builtin.PyInt{Value: 42}, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// float is float
		{floatItem, floatItem, builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.5}, builtin.PyFalse},
		// float is unextected_type
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// function is function
		{builtin.Builtin["print"], builtin.Builtin["print"], builtin.PyTrue},
		// TODO: function is function -> PyFalse
		// function is unexpected_type
		{builtin.Builtin["print"], &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// list is unexpected_type
		{listItem, &builtin.PyString{Value: "testy"}, builtin.PyFalse},
		// list is list
		{listItem, listItem, builtin.PyTrue},
		{
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				&builtin.PyString{Value: "testy"}},
			},
			builtin.PyFalse,
		},
	}
	for _, testCase := range cases {
		_testIsOp(t, testCase.Left, testCase.Right, testCase.Expected.(*builtin.PyBool))
	}
}

func _testIsOp(t *testing.T, a builtin.PyObject, b builtin.PyObject, expected builtin.PyObject) {
	instructions := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   a,
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   b,
		},
		{
			Opcode: opcode.IS_OP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	result := EvalInstructions(instructions)
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)
	if expectedType != resultType {
		t.Errorf("Returns incorect type, a: %v, b: %v, expected: %v, answer: %v", a, b, expectedType, resultType)
	}
	if result.Equal(expected) == builtin.PyFalse {
		t.Errorf("Returns incorect value, a: %v, b: %v, expected: %v, answer: %v", a, b, expected, result)
	}
}
