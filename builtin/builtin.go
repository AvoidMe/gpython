package builtin

import (
	"fmt"
	"main/pyobject"
)

var Builtin = map[string]interface{}{
	"print": GPythonPrint,
}

func GPythonPrint(args *pyobject.PyObject, kwnames *pyobject.PyObject) pyobject.PyObject {
	GPythonPrintImplementation(args, &pyobject.PyObject{Value: " "}, &pyobject.PyObject{Value: "\n"}, &pyobject.None, false)
	return pyobject.None
}

// https://github.com/python/cpython/blob/main/Python/bltinmodule.c#L1987 (builtin_print_impl)
func GPythonPrintImplementation(
	args *pyobject.PyObject,
	sep *pyobject.PyObject,
	end *pyobject.PyObject,
	file *pyobject.PyObject,
	flush bool,
) {
	for _, arg := range args.Tuple {
		fmt.Print(arg.Value)
		fmt.Print(sep.Value)
	}
	fmt.Print(end.Value)
}
