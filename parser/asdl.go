package parser

type asdl_seq [][]any
type asdl_stmt_seq [][]any
type asdl_expr_seq [][]any

// asdl_seq == asdl_stmt_seq or asdl_seq == asdl_expr_seq
type ASDL_INTERFACE interface {
	asdl_seq | asdl_stmt_seq | asdl_expr_seq
}

func _Py_asdl_generic_seq_new(n int) asdl_seq {
	return make(asdl_seq, 1)
}
