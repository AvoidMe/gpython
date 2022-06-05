package main

import (
	"io/ioutil"
	"log"
	"main/builtin"
	"main/opcode"
	"main/pyobject"
	pythontogo "main/python_to_go"
)

func pop(stack *[]pyobject.PyObject) pyobject.PyObject {
	s := *stack
	value := s[len(s)-1]
	*stack = s[:len(s)-1]
	return value
}

func popN(stack *[]pyobject.PyObject, n int) []pyobject.PyObject {
	s := *stack
	values := s[len(s)-n:]
	*stack = s[:len(s)-n]
	return values
}

func evalLoop(frame []opcode.Opcode) {
	stack := []pyobject.PyObject{}
	for _, instruction := range frame {
		log.Printf("Stack: %v\n", stack)
		log.Printf("Evaluating instruction: %v\n", instruction)
		switch instruction.Command {
		case opcode.POP_TOP:
			pop(&stack)
		case opcode.LOAD_CONST:
			stack = append(stack, instruction.Args[0])
		case opcode.LOAD_NAME:
			stack = append(stack, pyobject.PyObject{Function: builtin.Builtin[instruction.Args[0].Value]})
		case opcode.CALL_FUNCTION:
			args := pyobject.PyObject{Tuple: popN(&stack, instruction.Arg)}
			function := pop(&stack)
			called_func := function.Function.(func(*pyobject.PyObject, *pyobject.PyObject) pyobject.PyObject)
			result := called_func(&args, &pyobject.None)
			stack = append(stack, result)
		case opcode.RETURN_VALUE:
			log.Println("RETURN_VALUE is not implemented")
		default:
			log.Printf("Undefined opcode: %v\n", instruction)
		}
	}
}

func main() {
	log.SetOutput(ioutil.Discard)
	frame := pythontogo.LoadJson()
	evalLoop(frame)
}
