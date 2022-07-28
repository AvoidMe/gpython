package eval

import (
	"log"

	"github.com/AvoidMe/gpython/builtin"
	"github.com/AvoidMe/gpython/opcode"
)

func EvalInstructions(instructions []opcode.Instruction) builtin.PyObject {
	var returnValue builtin.PyObject
	returnValue = builtin.PyNone
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
		case opcode.NOP:
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
		case opcode.BINARY_SUBSTRACT:
			a := frame.Stack.Pop()
			b := frame.Stack.Pop().(builtin.PyBinarySubstract) // TODO: error handling
			result := b.BinarySubstract(a)                     // TODO: error handling
			frame.Stack.Append(result)
		case opcode.BINARY_SUBSCR:
			sub := frame.Stack.Pop()
			container := frame.Stack.Pop().(builtin.PyGetItem)
			item, _ := container.GetItem(sub) // TODO: error handling
			frame.Stack.Append(item)
		case opcode.STORE_SUBSCR:
			sub := frame.Stack.Pop()
			container := frame.Stack.Pop().(builtin.PySetItem)
			v := frame.Stack.Pop()
			/* container[sub] = v */
			container.SetItem(sub, v)
			frame.Stack.Append(container)
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
		case opcode.BUILD_MAP:
			dict := &builtin.PyDict{}
			for i := 0; i < instruction.Arg; i++ {
				value := frame.Stack.Pop()
				key := frame.Stack.Pop()
				dict.SetItem(key, value)
			}
			frame.Stack.Append(dict)
		case opcode.COMPARE_OP:
			a := frame.Stack.Pop()
			b := frame.Stack.Pop()
			switch instruction.Args.(*builtin.PyString).Value {
			case builtin.PyEq.Value:
				frame.Stack.Append(a.Equal(b))
			default:
				panic("Not implemented comparsion opcode")
			}
		case opcode.IS_OP:
			a := frame.Stack.Pop()
			b := frame.Stack.Pop()
			if a == b {
				frame.Stack.Append(builtin.PyTrue)
			} else {
				frame.Stack.Append(builtin.PyFalse)
			}
		case opcode.CALL_FUNCTION:
			args := &builtin.PyList{Value: frame.Stack.PopN(instruction.Arg)}
			function := frame.Stack.Pop().(*builtin.PyFunction)
			frame.Stack.Append(function.Callable(args, builtin.PyNone))
		case opcode.BUILD_CONST_KEY_MAP:
			dict := &builtin.PyDict{}
			keys := frame.Stack.Pop().(*builtin.PyList)
			for i := instruction.Arg - 1; i >= 0; i-- {
				index := &builtin.PyInt{Value: int64(i)}
				key, _ := keys.GetItem(index) // TODO: add error checking
				value := frame.Stack.Pop()
				dict.SetItem(key, value)
			}
			frame.Stack.Append(dict)
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
