package parser

type Parser struct {
	Keywords     [][]KeywordToken
	SoftKeywords []string
	StartRule    int

	tokens             []Token
	fill               int
	mark               int
	level              int
	error_indicator    int
	arena              any
	call_invalid_rules bool
}

// *_ty
type _mod_ty struct{}
type mod_ty *_mod_ty

type _expr_ty struct{}
type expr_ty *_expr_ty

type _stmt_ty struct{}
type stmt_ty *_stmt_ty

type _arg_ty struct{}
type arg_ty *_arg_ty

// asdl_*
type asdl_stmt_seq struct{}
type asdl_seq struct{}
type asdl_expr_seq struct{}

type _STMT interface {
	*asdl_seq | stmt_ty | expr_ty
}

type Token struct {
	lineno         int
	col_offset     int
	end_lineno     int
	end_col_offset int
}

func PyErr_Occurred() bool {
	return false
}

func _PyPegen_expect_token(p *Parser, marker int) *Token {
	return nil
}

func _PyPegen_seq_flatten(p *Parser, a *asdl_seq) *asdl_stmt_seq {
	return nil
}

func _PyPegen_fill_token(p *Parser) int {
	return 0
}

func _PyPegen_get_last_nonnwhitespace_token(p *Parser) *Token {
	return nil
}

func _PyPegen_is_memoized(p *Parser, _type int, pres any) bool {
	return true
}

func _PyPegen_name_token(p *Parser) expr_ty {
	return nil
}

func _PyPegen_make_module(p *Parser, a any) mod_ty {
	return nil
}

func _PyPegen_insert_memo[STMT_TYPE _STMT](p *Parser, a, b int, stmt STMT_TYPE) {
}

func _PyPegen_set_expr_context(p *Parser, a expr_ty, context expr_context) expr_ty {
	return nil
}

func _PyPegen_seq_insert_in_front(p *Parser, elem stmt_ty, seq *asdl_seq) *asdl_seq {
	return nil
}

func _PyPegen_lookahead_with_int(positive int, f func(*Parser, int) *Token, p *Parser, arg int) int {
	return 0
}

func _PyPegen_lookahead(positive int, f func(*Parser) expr_ty, p *Parser) int {
	return 0
}

func _PyPegen_string_token(p *Parser) expr_ty {
	return nil
}

func _PyPegen_number_token(p *Parser) expr_ty {
	return nil
}

func _Py_asdl_generic_seq_new(n int, arena any) *asdl_seq {
	return nil
}

func _PyPegen_concatenate_strings(p *Parser, a *asdl_seq) expr_ty {
	return nil
}

func asdl_seq_SET_UNTYPED(seq *asdl_seq, i int, child any) {
}

// Probably could be removed at all
func UNUSED(any) {}
