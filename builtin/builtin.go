package builtin

import (
	"fmt"
	gpythonlist "main/g_python_list"
	"main/pyobject"
)

var Builtin = map[string]interface{}{
	"print": GPythonPrint,
}

func GPythonPrint(args pyobject.PyObject, kwnames pyobject.PyObject) pyobject.PyObject {
	GPythonPrintImplementation(args, pyobject.PyObject{Value: " "}, pyobject.PyObject{Value: "\n"}, pyobject.None, false)
	return pyobject.None
}

type Printable interface {
	String() string
}

type Repr interface {
}

// https://github.com/python/cpython/blob/main/Python/bltinmodule.c#L1987 (builtin_print_impl)
func GPythonPrintImplementation(
	args pyobject.PyObject,
	sep pyobject.PyObject,
	end pyobject.PyObject,
	file pyobject.PyObject,
	flush bool,
) {
	list := args.Value.(gpythonlist.GpythonList).List
	for _, arg := range list {
		value := arg.Value.(Printable)
		fmt.Print(value.String())
		fmt.Print(sep.Value)
	}
	fmt.Print(end.Value)
}
