package pythontogo

import (
	"encoding/json"
	gpythonlist "main/g_python_list"
	gpythonnumber "main/g_python_number"
	gpythonstring "main/g_python_string"
	"main/opcode"
	"main/pyobject"
	"os"
	"strings"
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
				op.Args = gpythonstring.GpythonString{Str: v}
			case float64:
				// Golang unmarshal treats every JSON number as float64:
				// https://pkg.go.dev/encoding/json#Unmarshal
				// Here we're trying to check if source value was int or float
				if strings.Contains(value.Argrepr, ".") {
					op.Args = gpythonnumber.GpythonFloat{Float: v}
				} else {
					op.Args = gpythonnumber.GpythonInt{Int: int64(v)}
				}
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
