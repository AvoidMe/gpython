package builtin

import (
	"fmt"
	"main/pyobject"
)

var Builtin = map[string]interface{}{
	"print": GPythonPrint,
}

func GPythonPrint(arg pyobject.PyObject) {
	fmt.Println(arg.Value)
}
