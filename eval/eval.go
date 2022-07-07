package eval

import (
	"log"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
)

func EvalInstructions(instructions []opcode.Instruction) builtin.PyObject {
	var returnValue builtin.PyObject
	frame := Frame{
		Stack:  builtin.PyList{},
		Locals: map[string]builtin.PyObject{},
	}
	for _, instruction := range instructions {
		log.Printf("Stack: %v\n", frame.Stack)
		log.Printf("Evaluating instruction: %v\n", instruction)
		switch instruction.Opcode {
		case opcode.POP_TOP:
			frame.Stack.Pop()
		case opcode.BINARY_ADD:
			// Main idea for custom objects is:
			//	handle error while casting like this:
			//	 a, err := stack.Pop().(builtin.PybinaryAdd)
			//	if error occurs, than assume __dict__ has __add__ method,
			//  else: print same error as cpython does
			a := frame.Stack.Pop()
			b := frame.Stack.Pop().(builtin.PyBinaryAdd) // TODO: error handling
			result := b.BinaryAdd(a)                     // TODO: error handling
			frame.Stack.Append(result)
		case opcode.STORE_NAME:
			value := frame.Stack.Pop()
			frame.Locals[instruction.Args.(*builtin.PyString).Value] = value
		case opcode.LOAD_CONST:
			frame.Stack.Append(instruction.Args)
		case opcode.LOAD_NAME:
			name := instruction.Args.(*builtin.PyString)
			value, success := frame.Locals[name.Value]
			if success {
				frame.Stack.Append(value)
			} else {
				frame.Stack.Append(
					builtin.Builtin[name.Value],
				)
			}
		case opcode.BUILD_LIST:
			frame.Stack.Append(
				&builtin.PyList{Value: frame.Stack.PopN(instruction.Arg)},
			)
		case opcode.COMPARE_OP:
			a := frame.Stack.Pop()
			b := frame.Stack.Pop()
			switch instruction.Args.(*builtin.PyString).Value {
			case builtin.PyEq.Value:
				frame.Stack.Append(a.Equal(b))
			default:
				panic("Not implemented comparsion opcode")
			}
		case opcode.CALL_FUNCTION:
			args := &builtin.PyList{Value: frame.Stack.PopN(instruction.Arg)}
			function := frame.Stack.Pop().(*builtin.PyFunction)
			frame.Stack.Append(function.Callable(args, builtin.PyNone))
		case opcode.LIST_EXTEND:
			args := frame.Stack.Pop().(*builtin.PyList)
			list := frame.Stack.Pop().(*builtin.PyList)
			list.Extend(args.Value)
			frame.Stack.Append(list)
		case opcode.RETURN_VALUE:
			returnValue = frame.Stack.Pop()
		default:
			log.Printf("Undefined opcode: %v\n", instruction)
		}
	}
	return returnValue
}
