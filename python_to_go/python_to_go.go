package pythontogo

import (
	"encoding/json"
	"os"
	"strings"

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

func LoadJson() []opcode.Instruction {
	var bytecode []PythonBytecode
	var output []opcode.Instruction
	dat, _ := os.ReadFile("output.json")
	json.Unmarshal(dat, &bytecode)
	for _, value := range bytecode {
		op := opcode.Instruction{Opcode: value.Opcode, Arg: value.Arg}
		if value.Argval != nil {
			switch v := value.Argval.(type) {
			case string:
				op.Args = builtin.PyString{Value: v}
			case float64:
				// Golang unmarshal treats every JSON number as float64:
				// https://pkg.go.dev/encoding/json#Unmarshal
				// Here we're trying to check if source value was int or float
				if strings.Contains(value.Argrepr, ".") {
					op.Args = builtin.PyFloat{Value: v}
				} else {
					op.Args = builtin.PyInt{Value: int64(v)}
				}
			case []interface{}:
				list := builtin.PyList{Value: []builtin.PyObject{}}
				for i := 0; i < len(v); i++ {
					list.Append(builtin.PyString{Value: v[i].(string)})
				}
				op.Args = list
			}
		} else {
			op.Args = builtin.None
		}
		output = append(output, op)
	}
	return output
}
