package main

import (
	"io/ioutil"
	"log"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
	pythontogo "github.com/AvoidMe/gpython/python_to_go"
)

func evalFrame(frame []opcode.Instruction) builtin.PyObject {
	var returnValue builtin.PyObject
	stack := builtin.PyList{}
	locals := map[string]builtin.PyObject{}
	for _, instruction := range frame {
		log.Printf("Stack: %v\n", stack)
		log.Printf("Evaluating instruction: %v\n", instruction)
		switch instruction.Opcode {
		case opcode.POP_TOP:
			stack.Pop()
		case opcode.BINARY_ADD:
			// Main idea for custom objects is:
			//	handle error while casting like this:
			//	 a, err := stack.Pop().(builtin.PybinaryAdd)
			//	if error occurs, than assume __dict__ has __add__ method,
			//  else: print same error as cpython does
			a := stack.Pop().(builtin.PyBinaryAdd) // TODO: error handling
			b := stack.Pop().(builtin.PyBinaryAdd) // TODO: error handling
			result := b.BinaryAdd(a)               // TODO: error handling
			stack.Append(result)
		case opcode.STORE_NAME:
			value := stack.Pop()
			locals[instruction.Args.(builtin.PyString).Value] = value
		case opcode.LOAD_CONST:
			stack.Append(instruction.Args)
		case opcode.LOAD_NAME:
			name := instruction.Args.(builtin.PyString)
			value, success := locals[name.Value]
			if success {
				stack.Append(value)
			} else {
				stack.Append(
					builtin.Builtin[name.Value],
				)
			}
		case opcode.BUILD_LIST:
			stack.Append(
				builtin.PyList{Value: stack.PopN(instruction.Arg)},
			)
		case opcode.CALL_FUNCTION:
			args := builtin.PyList{Value: stack.PopN(instruction.Arg)}
			function := stack.Pop().(builtin.PyFunction)
			stack.Append(function.Callable(args, builtin.None))
		case opcode.LIST_EXTEND:
			args := stack.Pop().(builtin.PyList)
			list := stack.Pop().(builtin.PyList)
			list.Extend(args.Value)
			stack.Append(list)
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
