package opcode

import "main/pyobject"

type Opcode struct {
	Command int
	Arg     int
	Args    []pyobject.PyObject
}
