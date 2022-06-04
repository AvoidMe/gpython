package main

import (
	"io/ioutil"
	"log"
	"main/builtin"
	"main/opcode"
	"main/pyobject"
	pythontogo "main/python_to_go"
)

func popFromStack(stack *[]pyobject.PyObject) pyobject.PyObject {
	s := *stack
	value := s[len(s)-1]
	*stack = s[:len(s)-1]
	return value
}

func evalLoop(frame []opcode.Opcode) {
	stack := []pyobject.PyObject{}
	for _, value := range frame {
		log.Printf("Stack: %v\n", stack)
		log.Printf("Evaluating opcode: %v\n", value)
		switch value.Command {
		case opcode.LOAD_CONST:
			stack = append(stack, value.Args[0])
		case opcode.LOAD_NAME:
			stack = append(stack, pyobject.PyObject{Function: builtin.Builtin[value.Args[0].Value]})
		case opcode.CALL_FUNCTION:
			argument := popFromStack(&stack)
			function := popFromStack(&stack)
			called_func := function.Function.(func(pyobject.PyObject))
			called_func(argument)
		case opcode.RETURN_VALUE:
			log.Println("RETURN_VALUE is not implemented")
		default:
			log.Printf("Undefined opcode: %v\n", value)
		}
	}
}

func main() {
	log.SetOutput(ioutil.Discard)
	frame := pythontogo.LoadJson()
	evalLoop(frame)
}
