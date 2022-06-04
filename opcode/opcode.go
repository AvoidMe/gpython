package opcode

import "main/pyobject"

type Opcode struct {
	Command int
	Args    []pyobject.PyObject
}
