package parser

type AST struct{}

const (
	Load  = 1
	Store = 2
	Del   = 3

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

func _PyAST_Interactive(a *asdl_stmt_seq, arena any) mod_ty {
	return nil
}

func _PyAST_Expression(a expr_ty, arena any) mod_ty {
	return nil
}

func _PyAST_FunctionType(a any, b expr_ty, arena any) mod_ty {
	return nil
}
