package builtin

import (
	"fmt"
)

var Builtin = map[string]PyObject{
	// TODO: probably, we can reduce string/repr to only names and
	//		 leave string/repr implementation to PyFunction itself
	"print": &PyFunction{
		Callable:    PyPrint,
		StringValue: "<built-in function print>",
		ReprValue:   "<built-in function print>",
	},
	"hash": &PyFunction{
		Callable:    PyHash,
		StringValue: "<built-in function hash>",
		ReprValue:   "<built-in function hash>",
	},
}

func PyPrint(args PyObject, kwnames PyObject) PyObject {
	PyPrintImplementation(args, NewPyString(" "), NewPyString("\n"), PyNone, PyFalse)
	return PyNone
}

// https://github.com/python/cpython/blob/main/Python/bltinmodule.c#L1987 (builtin_print_impl)
func PyPrintImplementation(
	args PyObject,
	sep PyObject,
	end PyObject,
	file PyObject,
	flush *PyBool,
) {
	for _, arg := range args.(*PyList).Value {
		fmt.Print(arg.String())
		fmt.Print(sep.String())
	}
	fmt.Print(end.String())
}

func PyHash(args PyObject, kwnames PyObject) PyObject {
	// TODO: check that args and kwnames contains only one argument
	argList, _ := args.(*PyList)
	arg, _ := argList.GetItem(&PyInt{Value: 0})
	hash, _ := PyHashImplementation(arg) // TODO: check error
	return &PyInt{Value: int64(hash)}
}

func PyHashImplementation(obj PyObject) (uint64, error) {
	return obj.Hash()
}
