package parser

type Parser struct {
	Keywords     [][]KeywordToken
	SoftKeywords []string
	StartRule    int

	tokens             []*Token
	fill               int
	mark               int
	level              int
	error_indicator    int
	call_invalid_rules bool
}

type _mod_ty struct{}
type mod_ty *_mod_ty

type _expr_ty struct{}
type expr_ty *_expr_ty

type _stmt_ty struct{}
type stmt_ty *_stmt_ty

type _arg_ty struct{}
type arg_ty *_arg_ty

type _STMT interface {
	asdl_seq | stmt_ty | expr_ty
}

type Token struct {
	_type          int
	lineno         int
	col_offset     int
	end_lineno     int
	end_col_offset int
}

func (this *Parser) PyErr_Occurred() bool {
	return false
}

func _PyPegen_expect_token(p *Parser, _type int) *Token {
	if p.mark == p.fill {
		if _PyPegen_fill_token(p) < 0 {
			p.error_indicator = 1
			return nil
		}
	}
	t := p.tokens[p.mark]
	if t._type != _type {
		return nil
	}
	p.mark += 1
	return t
}

func _PyPegen_seq_flatten[T ASDL_INTERFACE](p *Parser, seqs T) T {
	flattened_seq := make(T, 1)
	l := len(seqs)
	for i := 0; i < l; i++ {
		inner_seq := seqs[i]
		for j := 0; j < len(inner_seq); j++ {
			flattened_seq[0] = append(flattened_seq[0], inner_seq[j])
		}
	}
	return flattened_seq
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

func _PyPegen_seq_insert_in_front(p *Parser, elem stmt_ty, seq asdl_seq) asdl_seq {
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

func _PyPegen_concatenate_strings(p *Parser, a asdl_seq) expr_ty {
	return nil
}

// Probably could be removed at all
func UNUSED(any) {}
