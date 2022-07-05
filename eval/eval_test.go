package eval

import (
	"reflect"
	"testing"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
)

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
			Args:   builtin.PyInt{Value: expected},
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyInt)
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
			builtin.PyString{Value: "First"},
			builtin.PyString{Value: "Second"},
		},
	}
	secondList := builtin.PyList{
		Value: []builtin.PyObject{
			builtin.PyString{Value: "Third"},
			builtin.PyString{Value: "Fourth"},
		},
	}
	expected := []builtin.PyString{
		{Value: "First"},
		{Value: "Second"},
		{Value: "Third"},
		{Value: "Fourth"},
	}
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   firstList,
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   secondList,
		},
		{
			Opcode: opcode.LIST_EXTEND,
			Arg:    0,
			Args:   builtin.None,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyList)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if len(answer.Value) != len(expected) {
		t.Errorf("Returns wrong list: %v", answer)
	}
	for index, value := range answer.Value {
		strValue, ok := value.(builtin.PyString)
		if !ok {
			t.Errorf("Got: %v, expected: %v", answer, expected)
		}
		if strValue.Value != expected[index].Value {
			t.Errorf("Got: %v, expected: %v", answer, expected)
		}
	}
}

func TestBuildList(t *testing.T) {
	expected := []builtin.PyString{
		{Value: "First"},
		{Value: "Second"},
		{Value: "Third"},
		{Value: "Fourth"},
	}
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: "First"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: "Second"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: "Third"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: "Fourth"},
		},
		{
			Opcode: opcode.BUILD_LIST,
			Arg:    4,
			Args:   builtin.None,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyList)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if len(answer.Value) != len(expected) {
		t.Errorf("Returns wrong list: %v", answer)
	}
	for index, value := range answer.Value {
		strValue, ok := value.(builtin.PyString)
		if !ok {
			t.Errorf("Got: %v, expected: %v", answer, expected)
		}
		if strValue.Value != expected[index].Value {
			t.Errorf("Got: %v, expected: %v", answer, expected)
		}
	}
}

func TestStoreLocalName(t *testing.T) {
	expected := 42.05
	input := []opcode.Instruction{
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyFloat{Value: 42.05},
		},
		{
			Opcode: opcode.STORE_NAME,
			Arg:    0,
			Args:   builtin.PyString{Value: "x"},
		},
		{
			Opcode: opcode.LOAD_NAME,
			Arg:    0,
			Args:   builtin.PyString{Value: "x"},
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyFloat)
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
			Args:   builtin.PyFloat{Value: 42.05},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyInt{Value: 42},
		},
		{
			Opcode: opcode.POP_TOP,
			Arg:    0,
			Args:   builtin.None,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyFloat)
	if !ok {
		t.Errorf("Returns incorrect type: %v", result)
	}
	if answer.Value != expected {
		t.Errorf("Returns incorect value: %v", answer)
	}
}

func TestBinaryAdd(t *testing.T) {
	input := []opcode.Instruction{
		// str + str
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: "Hello"},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyString{Value: ", world!"},
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.None,
		},
		// int + int
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyInt{Value: 10},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyInt{Value: 32},
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.None,
		},
		// int + float
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyInt{Value: 10},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyFloat{Value: 32.54321},
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.None,
		},
		// float + int
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyFloat{Value: 10.54321},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyInt{Value: 32},
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.None,
		},
		// float + float
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyFloat{Value: 10.54321},
		},
		{
			Opcode: opcode.LOAD_CONST,
			Arg:    0,
			Args:   builtin.PyFloat{Value: 32.12345},
		},
		{
			Opcode: opcode.BINARY_ADD,
			Arg:    0,
			Args:   builtin.None,
		},
		// Return values
		{
			Opcode: opcode.BUILD_LIST,
			Arg:    5,
			Args:   builtin.None,
		},
		{
			Opcode: opcode.RETURN_VALUE,
			Arg:    0,
			Args:   builtin.None,
		},
	}
	expected := []builtin.PyObject{
		builtin.PyString{Value: "Hello, world!"},
		builtin.PyInt{Value: 42},
		builtin.PyFloat{Value: 42.54321},
		builtin.PyFloat{Value: 42.54321},
		builtin.PyFloat{Value: 42.66666},
	}
	result := EvalInstructions(input)
	answer, ok := result.(builtin.PyList)
	if !ok {
		t.Errorf("Returns wrong type: %v", answer)
	}
	if len(answer.Value) != len(expected) {
		t.Errorf("Returns wrong list: %v", answer)
	}
	for index, value := range answer.Value {
		if reflect.TypeOf(value) != reflect.TypeOf(expected[index]) {
			t.Errorf("Returns incorect value: %v", answer)
		}
		switch a := value.(type) {
		case builtin.PyString:
			b := expected[index].(builtin.PyString)
			if a.Value != b.Value {
				t.Errorf("Returns incorect value: %v", answer)
			}
		case builtin.PyInt:
			b := expected[index].(builtin.PyInt)
			if a.Value != b.Value {
				t.Errorf("Returns incorect value: %v", answer)
			}
		case builtin.PyFloat:
			b := expected[index].(builtin.PyFloat)
			if a.Value != b.Value {
				t.Errorf("Returns incorect value: %v", answer)
			}
		default:
			t.Errorf("Expected has unexpected type: %v", expected)
		}
	}
}
