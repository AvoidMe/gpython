package pythontogo

import (
	"encoding/json"
	gpythonlist "main/g_python_list"
	gpythonstring "main/g_python_string"
	"main/opcode"
	"main/pyobject"
	"os"
)

type PythonBytecode struct {
	Opcode         int
	Opname         string
	Arg            int
	Argval         interface{}
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
				op.Args = gpythonstring.GpythonString{Str: v}
			case []interface{}:
				list := gpythonlist.GpythonList{List: []pyobject.PyObject{}}
				for i := 0; i < len(v); i++ {
					list.Append(gpythonstring.GpythonString{Str: v[i].(string)})
				}
				op.Args = list
			}
		} else {
			op.Args = pyobject.None
		}
		output = append(output, op)
	}
	return output
}
