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
		innerList, success := arg.Value.(gpythonlist.GpythonList)
		if success {
			fmt.Print(innerList.String())
		} else {
			fmt.Print(arg.Value)
		}
		fmt.Print(sep.Value)
	}
	fmt.Print(end.Value)
}
