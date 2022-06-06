// TODO: I decided to remove all of the pointers,
// since I obviously fucked up somewhere with them,
// I need to have a better understanding of it before using it aggressively
package main

import (
	"io/ioutil"
	"log"
	"main/builtin"
	gpythonlist "main/g_python_list"
	"main/opcode"
	"main/pyobject"
	pythontogo "main/python_to_go"
)

func evalFrame(frame []opcode.Instruction) pyobject.PyObject {
	var returnValue pyobject.PyObject
	stack := gpythonlist.GpythonList{}
	for _, instruction := range frame {
		log.Printf("Stack: %v\n", stack)
		log.Printf("Evaluating instruction: %v\n", instruction)
		switch instruction.Opcode {
		case opcode.POP_TOP:
			stack.Pop()
		case opcode.LOAD_CONST:
			stack.Append(instruction.Args[0])
		case opcode.LOAD_NAME:
			stack.Append(pyobject.PyObject{Value: builtin.Builtin[instruction.Args[0].Value.(string)]})
		case opcode.BUILD_LIST:
			stack.Append(
				pyobject.PyObject{
					Value: gpythonlist.GpythonList{List: stack.PopN(instruction.Arg)},
				},
			)
		case opcode.CALL_FUNCTION:
			listArgs := stack.PopN(instruction.Arg)
			args := pyobject.PyObject{Value: gpythonlist.GpythonList{List: listArgs}}
			function := stack.Pop()
			called_func := function.Value.(func(pyobject.PyObject, pyobject.PyObject) pyobject.PyObject)
			result := called_func(args, pyobject.None)
			stack.Append(result)
		case opcode.LIST_EXTEND:
			value := stack.Pop()
			list := stack.Pop()
			argToExtend := value.Value.(gpythonlist.GpythonList)
			realList := list.Value.(gpythonlist.GpythonList)
			realList.Extend(argToExtend.List)
		case opcode.RETURN_VALUE:
			returnValue = stack.Pop()
		default:
			log.Printf("Undefined opcode: %v\n", instruction)
		}
	}
	return returnValue
}

func main() {
	log.SetOutput(ioutil.Discard)
	frame := pythontogo.LoadJson()
	evalFrame(frame)
}
