package builtin

import (
	"fmt"
	gpythonfunction "main/g_python_function"
	gpythonlist "main/g_python_list"
	gpythonstring "main/g_python_string"
	"main/pyobject"
)

var Builtin = map[string]pyobject.PyObject{
	"print": gpythonfunction.GPythonFunction{
		Callable:    GPythonPrint,
		StringValue: "<built-in function print>",
		ReprValue:   "<built-in function print>",
	},
}

func GPythonPrint(args pyobject.PyObject, kwnames pyobject.PyObject) pyobject.PyObject {
	GPythonPrintImplementation(args, gpythonstring.GpythonString{Str: " "}, gpythonstring.GpythonString{Str: "\n"}, pyobject.None, false)
	return pyobject.None
}

// https://github.com/python/cpython/blob/main/Python/bltinmodule.c#L1987 (builtin_print_impl)
func GPythonPrintImplementation(
	args pyobject.PyObject,
	sep pyobject.PyObject,
	end pyobject.PyObject,
	file pyobject.PyObject,
	flush bool,
) {
	list := args.(gpythonlist.GpythonList).List
	for _, arg := range list {
		fmt.Print(arg.String())
		fmt.Print(sep.String())
	}
	fmt.Print(end.String())
}
