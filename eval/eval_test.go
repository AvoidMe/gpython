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

func TestNOP(t *testing.T) {
	input := []opcode.Instruction{
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
		{
			Opcode: opcode.NOP,
			Arg:    0,
			Args:   builtin.PyNone,
		},
	}
	EvalInstructions(input)
}

func TestReturnValue(t *testing.T) {
	var expected int64 = 42
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.NewPyInt(expected),
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
	if answer.Equal(builtin.NewPyInt(42)) != builtin.PyTrue {
		t.Errorf("Returns incorrect value: %v", answer)
	}
}

func TestListExtend(t *testing.T) {
	firstList := builtin.PyList{
		Value: []builtin.PyObject{
			builtin.NewPyString("First"),
			builtin.NewPyString("Second"),
		},
	}
	secondList := builtin.PyList{
		Value: []builtin.PyObject{
			builtin.NewPyString("Third"),
			builtin.NewPyString("Fourth"),
		},
	}
	expected := builtin.PyList{
		Value: []builtin.PyObject{
			builtin.NewPyString("First"),
			builtin.NewPyString("Second"),
			builtin.NewPyString("Third"),
			builtin.NewPyString("Fourth"),
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
			builtin.NewPyString("First"),
			builtin.NewPyString("Second"),
			builtin.NewPyString("Third"),
			builtin.NewPyString("Fourth"),
		},
	}
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.NewPyString("First"),
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.NewPyString("Second"),
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.NewPyString("Third"),
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.NewPyString("Fourth"),
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
			Args:   builtin.NewPyString("x"),
		},
		{
			Opcode: opcode.LOAD_NAME,
			Arg:    0,
			Args:   builtin.NewPyString("x"),
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
			Args:   builtin.NewPyInt(42),
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
			builtin.NewPyString("Hello"),
			builtin.NewPyString(", world!"),
			builtin.NewPyString("Hello, world!"),
		},
		{
			builtin.NewPyInt(10),
			builtin.NewPyInt(32),
			builtin.NewPyInt(42),
		},
		{
			builtin.NewPyInt(10),
			&builtin.PyFloat{Value: 32.54321},
			&builtin.PyFloat{Value: 42.54321},
		},
		{
			&builtin.PyFloat{Value: 10.54321},
			builtin.NewPyInt(32),
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
			builtin.NewPyInt(1),
		},
		{
			builtin.PyFalse,
			builtin.PyFalse,
			builtin.NewPyInt(0),
		},
		{
			builtin.PyTrue,
			builtin.PyTrue,
			builtin.NewPyInt(2),
		},
		{
			builtin.PyTrue,
			builtin.NewPyInt(41),
			builtin.NewPyInt(42),
		},
		{
			builtin.PyFalse,
			builtin.NewPyInt(41),
			builtin.NewPyInt(41),
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
			builtin.NewPyInt(41),
			builtin.PyFalse,
			builtin.NewPyInt(41),
		},
		{
			builtin.NewPyInt(41),
			builtin.PyTrue,
			builtin.NewPyInt(42),
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

func TestBinarySubstract(t *testing.T) {
	cases := []binaryTestCase{
		{
			builtin.NewPyInt(52),
			builtin.NewPyInt(10),
			builtin.NewPyInt(42),
		},
		{
			&builtin.PyFloat{Value: 52.54321},
			builtin.NewPyInt(10),
			&builtin.PyFloat{Value: 42.54321},
		},
		{
			&builtin.PyFloat{Value: 52.54321},
			builtin.NewPyInt(10),
			&builtin.PyFloat{Value: 42.54321},
		},
		{
			&builtin.PyFloat{Value: 52.66666},
			&builtin.PyFloat{Value: 10.54321},
			&builtin.PyFloat{Value: 42.12345},
		},
		{
			builtin.PyTrue,
			builtin.PyFalse,
			builtin.NewPyInt(1),
		},
		{
			builtin.PyFalse,
			builtin.PyFalse,
			builtin.NewPyInt(0),
		},
		{
			builtin.PyTrue,
			builtin.PyTrue,
			builtin.NewPyInt(0),
		},
		{
			builtin.PyTrue,
			builtin.NewPyInt(41),
			builtin.NewPyInt(-40),
		},
		{
			builtin.PyFalse,
			builtin.NewPyInt(41),
			builtin.NewPyInt(-41),
		},
		{
			builtin.PyFalse,
			&builtin.PyFloat{Value: 41.5},
			&builtin.PyFloat{Value: -41.5},
		},
		{
			&builtin.PyFloat{Value: 41.5},
			builtin.PyFalse,
			&builtin.PyFloat{Value: 41.5},
		},
		{
			builtin.PyTrue,
			&builtin.PyFloat{Value: 41.5},
			&builtin.PyFloat{Value: -40.5},
		},
		{
			&builtin.PyFloat{Value: 41.5},
			builtin.PyTrue,
			&builtin.PyFloat{Value: 40.5},
		},
		{
			builtin.NewPyInt(41),
			builtin.PyFalse,
			builtin.NewPyInt(41),
		},
		{
			builtin.NewPyInt(41),
			builtin.PyTrue,
			builtin.NewPyInt(40),
		},
	}
	for _, testCase := range cases {
		_testBinarySubstract(t, testCase.Left, testCase.Right, testCase.Expected)
	}
}

func _testBinarySubstract(t *testing.T, a builtin.PyObject, b builtin.PyObject, expected builtin.PyObject) {
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
			Opcode: opcode.BINARY_SUBSTRACT,
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
		{builtin.PyNone, builtin.NewPyString("testy"), builtin.PyFalse},
		// bool == bool
		{builtin.PyTrue, builtin.PyTrue, builtin.PyTrue},
		{builtin.PyTrue, builtin.PyFalse, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyTrue, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyFalse, builtin.PyTrue},
		// bool == unexpected_type
		{builtin.PyTrue, builtin.NewPyString("testy"), builtin.PyFalse},
		// bool == int
		{builtin.PyTrue, builtin.NewPyInt(1), builtin.PyTrue},
		{builtin.PyTrue, builtin.NewPyInt(0), builtin.PyFalse},
		{builtin.PyFalse, builtin.NewPyInt(1), builtin.PyFalse},
		{builtin.PyFalse, builtin.NewPyInt(0), builtin.PyTrue},
		// bool == float
		{builtin.PyTrue, &builtin.PyFloat{Value: 1.0}, builtin.PyTrue},
		{builtin.PyTrue, &builtin.PyFloat{Value: 0.0}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyFloat{Value: 1.0}, builtin.PyFalse},
		{builtin.PyFalse, &builtin.PyFloat{Value: 0.0}, builtin.PyTrue},
		// int == bool
		{builtin.NewPyInt(1), builtin.PyTrue, builtin.PyTrue},
		{builtin.NewPyInt(0), builtin.PyTrue, builtin.PyFalse},
		{builtin.NewPyInt(1), builtin.PyFalse, builtin.PyFalse},
		{builtin.NewPyInt(0), builtin.PyFalse, builtin.PyTrue},
		// int == int
		{builtin.NewPyInt(42), builtin.NewPyInt(42), builtin.PyTrue},
		{builtin.NewPyInt(42), builtin.NewPyInt(444), builtin.PyFalse},
		// int == unexpected_type
		{builtin.NewPyInt(42), builtin.NewPyString("testy"), builtin.PyFalse},
		// int == float
		{builtin.NewPyInt(42), &builtin.PyFloat{Value: 42.0}, builtin.PyTrue},
		{builtin.NewPyInt(42), &builtin.PyFloat{Value: 42.5}, builtin.PyFalse},
		// float == int
		{&builtin.PyFloat{Value: 42.0}, builtin.NewPyInt(42), builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, builtin.NewPyInt(42), builtin.PyFalse},
		// float == float
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.5}, builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.555}, builtin.PyFalse},
		// float == bool
		{&builtin.PyFloat{Value: 1.0}, builtin.PyTrue, builtin.PyTrue},
		{&builtin.PyFloat{Value: 0.0}, builtin.PyTrue, builtin.PyFalse},
		{&builtin.PyFloat{Value: 1.0}, builtin.PyFalse, builtin.PyFalse},
		{&builtin.PyFloat{Value: 0.0}, builtin.PyFalse, builtin.PyTrue},
		// float == unextected_type
		{&builtin.PyFloat{Value: 42.5}, builtin.NewPyString("testy"), builtin.PyFalse},

		// function == function
		{builtin.Builtin["print"], builtin.Builtin["print"], builtin.PyTrue},
		// TODO: function == function -> PyFals

		// function == unexpected_type
		{builtin.Builtin["print"], builtin.NewPyString("testy"), builtin.PyFalse},
		// list == unexpected_type
		{
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			builtin.NewPyString("testy"),
			builtin.PyFalse,
		},
		// list == list
		{
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy_test_test")},
			},
			builtin.PyFalse,
		},
		{
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			builtin.PyTrue,
		},
		// list with nested lists
		{
			&builtin.PyList{
				Value: []builtin.PyObject{
					builtin.NewPyString("testy"),
					&builtin.PyList{
						Value: []builtin.PyObject{
							builtin.NewPyString("testy"),
						},
					},
				},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			builtin.PyFalse,
		},
		{
			&builtin.PyList{
				Value: []builtin.PyObject{
					builtin.NewPyString("testy"),
					&builtin.PyList{
						Value: []builtin.PyObject{
							builtin.NewPyString("testy"),
						},
					},
				},
			},
			&builtin.PyList{
				Value: []builtin.PyObject{
					builtin.NewPyString("testy"),
					&builtin.PyList{
						Value: []builtin.PyObject{
							builtin.NewPyString("testy"),
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
		builtin.NewPyString("testy"),
	}}
	intItem := builtin.NewPyInt(42)
	floatItem := &builtin.PyFloat{Value: 42.5}
	cases := []binaryTestCase{
		// None
		{builtin.PyNone, builtin.PyNone, builtin.PyTrue},
		{builtin.PyNone, builtin.NewPyString("testy"), builtin.PyFalse},
		// bool is bool
		{builtin.PyTrue, builtin.PyTrue, builtin.PyTrue},
		{builtin.PyTrue, builtin.PyFalse, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyTrue, builtin.PyFalse},
		{builtin.PyFalse, builtin.PyFalse, builtin.PyTrue},
		// bool is unexpected_type
		{builtin.PyTrue, builtin.NewPyString("testy"), builtin.PyFalse},
		// int is int
		{intItem, intItem, builtin.PyTrue},
		{builtin.NewPyInt(9999999), intItem, builtin.PyFalse},
		// int is unexpected_type
		{builtin.NewPyInt(42), builtin.NewPyString("testy"), builtin.PyFalse},
		// float is float
		{floatItem, floatItem, builtin.PyTrue},
		{&builtin.PyFloat{Value: 42.5}, &builtin.PyFloat{Value: 42.5}, builtin.PyFalse},
		// float is unextected_type
		{&builtin.PyFloat{Value: 42.5}, builtin.NewPyString("testy"), builtin.PyFalse},
		// function is function
		{builtin.Builtin["print"], builtin.Builtin["print"], builtin.PyTrue},
		// TODO: function is function -> PyFalse
		// function is unexpected_type
		{builtin.Builtin["print"], builtin.NewPyString("testy"), builtin.PyFalse},
		// list is unexpected_type
		{listItem, builtin.NewPyString("testy"), builtin.PyFalse},
		// list is list
		{listItem, listItem, builtin.PyTrue},
		{
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
			},
			&builtin.PyList{Value: []builtin.PyObject{
				builtin.NewPyString("testy")},
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

// TODO:
//	func testGetItemDict
//  func testSetItemDict
//  func testMakeDict
//  func testMakeDictKeyMap
