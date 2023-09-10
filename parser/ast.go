package parser

import (
	"github.com/AvoidMe/gpython/builtin"
)

type AST struct{}

type expr_context int

const (
	Load  expr_context = 1
	Store expr_context = 2
	Del   expr_context = 3
)

const (
	And = 1
	Or  = 2

	Add      = 1
	Sub      = 2
	Mult     = 3
	MatMult  = 4
	Div      = 5
	Mod      = 6
	Pow      = 7
	LShift   = 8
	RShift   = 9
	BitOr    = 10
	BitXor   = 11
	BitAnd   = 12
	FloorDiv = 13

	Invert = 1
	Not    = 2
	UAdd   = 3
	USub   = 4

	Eq    = 1
	NotEq = 2
	Lt    = 3
	LtE   = 4
	Gt    = 5
	GtE   = 6
	Is    = 7
	IsNot = 8
	In    = 9
	NotIn = 10
)

func _PyAST_Pass(start_lineno, start_col_offset, end_lineno, end_col_offset int) stmt_ty {
	return nil
}

func _PyAST_Expr(e expr_ty, start_lineno, start_col_offset, end_lineno, end_col_offset int) stmt_ty {
	return nil
}

func _PyAST_Assign(a *asdl_expr_seq, b expr_ty, type_comment builtin.PyObject, start_lineno, start_col_offset, end_lineno, end_col_offset int) stmt_ty {
	return nil
}

func _PyAST_Constant(constant builtin.PyObject, kind *string, start_lineno, start_col_offset, end_lineno, end_col_offset int) expr_ty {
	return nil
}

func NEW_TYPE_COMMENT(p *Parser, token any) builtin.PyObject {
	return nil
}
