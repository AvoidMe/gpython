package main

import (
	"io/ioutil"
	"log"
	"main/builtin"
	gpythonfunction "main/g_python_function"
	gpythonlist "main/g_python_list"
	gpythonstring "main/g_python_string"
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
			stack.Append(instruction.Args)
		case opcode.LOAD_NAME:
			name := instruction.Args.(gpythonstring.GpythonString)
			stack.Append(
				builtin.Builtin[name.Str],
			)
		case opcode.BUILD_LIST:
			stack.Append(
				gpythonlist.GpythonList{List: stack.PopN(instruction.Arg)},
			)
		case opcode.CALL_FUNCTION:
			args := gpythonlist.GpythonList{List: stack.PopN(instruction.Arg)}
			function := stack.Pop().(gpythonfunction.GPythonFunction)
			stack.Append(function.Callable(args, pyobject.None))
		case opcode.LIST_EXTEND:
			args := stack.Pop().(gpythonlist.GpythonList)
			list := stack.Pop().(gpythonlist.GpythonList)
			list.Extend(args.List)
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
