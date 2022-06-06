package pythontogo

import (
	"encoding/json"
	gpythonlist "main/g_python_list"
	"main/opcode"
	"main/pyobject"
	"os"
	"reflect"
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
			switch reflect.TypeOf(value.Argval).Kind() {
			case reflect.String:
				str := value.Argval.(string)
				op.Args = []pyobject.PyObject{
					{Value: str},
				}
			case reflect.Slice:
				s := reflect.ValueOf(value.Argval)
				list := gpythonlist.GpythonList{List: []pyobject.PyObject{}}
				for i := 0; i < s.Len(); i++ {
					str := s.Index(i).Interface().(string)
					list.Append(pyobject.PyObject{Value: str})
				}
				op.Args = []pyobject.PyObject{
					{Value: list},
				}
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
