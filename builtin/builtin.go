package builtin

import (
	"fmt"
)

var Builtin = map[string]PyObject{
	"print": &PyFunction{
		Callable:    GPythonPrint,
		StringValue: "<built-in function print>",
		ReprValue:   "<built-in function print>",
	},
}

func GPythonPrint(args PyObject, kwnames PyObject) PyObject {
	GPythonPrintImplementation(args, NewPyString(" "), NewPyString("\n"), PyNone, false)
	return PyNone
}

// https://github.com/python/cpython/blob/main/Python/bltinmodule.c#L1987 (builtin_print_impl)
func GPythonPrintImplementation(
	args PyObject,
	sep PyObject,
	end PyObject,
	file PyObject,
	flush bool,
) {
	for _, arg := range args.(*PyList).Value {
		fmt.Print(arg.String())
		fmt.Print(sep.String())
	}
	fmt.Print(end.String())
}
