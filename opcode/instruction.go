package opcode

import "main/pyobject"

type Instruction struct {
	Opcode int
	Arg    int
	Args   pyobject.PyObject
}
