package eval

import "github.com/AvoidMe/gpython/builtin"

type Frame struct {
	Stack  builtin.PyList
	Locals map[string]builtin.PyObject
}
