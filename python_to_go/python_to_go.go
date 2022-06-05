package pythontogo

import (
	"encoding/json"
	"main/opcode"
	"main/pyobject"
	"os"
)

type PythonBytecode struct {
	Opcode         int
	Opname         string
	Arg            int
	Argval         *string
	Offset         int
	Is_jump_target bool
}

func LoadJson() []opcode.Opcode {
	var bytecode []PythonBytecode
	var output []opcode.Opcode
	dat, _ := os.ReadFile("output.json")
	json.Unmarshal(dat, &bytecode)
	for _, value := range bytecode {
		op := opcode.Opcode{Command: value.Opcode, Arg: value.Arg}
		if value.Argval != nil {
			op.Args = []pyobject.PyObject{
				{Value: *value.Argval},
			}
		} else {
			op.Args = []pyobject.PyObject{
				pyobject.None,
			}
		}
		output = append(output, op)
	}
	return output
}
