package pythontogo

import (
	"encoding/json"
	"os"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
)

type PythonBytecode struct {
	Opcode         int
	Opname         string
	Arg            int
	Argval         interface{}
	Argrepr        string
	Offset         int
	Is_jump_target bool
}

func jsonTypeToPy(value interface{}) builtin.PyObject {
	switch v := value.(type) {
	case bool:
		return builtin.PyBool{Value: v}
	case string:
		return builtin.PyString{Value: v}
	case float64:
		// Golang unmarshal treats every JSON number as float64:
		// https://pkg.go.dev/encoding/json#Unmarshal
		// Here we're trying to check if source value was int or float
		if float64(int64(v)) == v {
			return builtin.PyInt{Value: int64(v)}
		} else {
			return builtin.PyFloat{Value: v}
		}
	case []interface{}:
		list := builtin.PyList{Value: []builtin.PyObject{}}
		for i := 0; i < len(v); i++ {
			list.Append(jsonTypeToPy(v[i]))
		}
		return list
	default:
		panic("Json has undefined type")
	}
}

func LoadJson() []opcode.Instruction {
	var bytecode []PythonBytecode
	var output []opcode.Instruction
	dat, _ := os.ReadFile("output.json")
	json.Unmarshal(dat, &bytecode)
	for _, value := range bytecode {
		op := opcode.Instruction{Opcode: value.Opcode, Arg: value.Arg}
		if value.Argval != nil {
			op.Args = jsonTypeToPy(value.Argval)
		} else {
			op.Args = builtin.PyNone
		}
		output = append(output, op)
	}
	return output
}
