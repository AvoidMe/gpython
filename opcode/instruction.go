package opcode

import "github.com/AvoidMe/gpython/builtin"

type Instruction struct {
	Opcode int
	Arg    int
	Args   builtin.PyObject
}
