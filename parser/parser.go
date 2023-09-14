package parser

type Parser struct {
	Keywords     [][]KeywordToken
	SoftKeywords []string
	StartRule    int

	tokens               []*Token
	tok                  *TokenState
	type_ignore_comments []*Token
	fill                 int
	mark                 int
	level                int
	error_indicator      int
	start_rule           int
	flags                int
	call_invalid_rules   bool
	parsing_started      bool
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
	value          string
}

const (
	MAXLEVEL  = 200
	MAXINDENT = 100
)

type TokenState struct {
	lineno                int
	pendin                int
	indent                int
	done                  int
	col_offset            int
	starting_col_offset   int
	atbol                 int
	tabsize               int
	level                 int
	async_def             bool
	async_def_nl          int
	async_def_indent      int
	decoding_erred        bool
	type_comments         bool
	cur                   byte
	start                 byte
	end                   byte
	interactive_src_start byte /* The start of the source parsed so far in interactive mode */
	interactive_src_end   byte /* The end of the source parsed so far in interactive mode */
	prompt, nextprompt    byte /* For interactive prompting */

	parenstack       [MAXLEVEL]byte
	parenlinenostack [MAXLEVEL]int
	parencolstack    [MAXLEVEL]int
	indstack         [MAXINDENT]int /* Stack of indents */
	altindstack      [MAXINDENT]int /* Stack of alternate indents */
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
	new_token := &Token{}
	_type := _PyTokenizer_Get(p.tok, new_token)

	// Record and skip '# type: ignore' comments
	for _type == TYPE_IGNORE {
		tag := new_token.value
		// Ownership of tag passes to the growable array
		p.type_ignore_comments = append(p.type_ignore_comments, &Token{lineno: p.tok.lineno, value: tag})
		_type = _PyTokenizer_Get(p.tok, new_token)
	}

	// If we have reached the end and we are in single input mode we need to insert a newline and reset the parsing
	if p.start_rule == Py_single_input && _type == ENDMARKER && p.parsing_started {
		_type = NEWLINE /* Add an extra newline */
		p.parsing_started = false

		if p.tok.indent > 0 && (p.flags&PyPARSE_DONT_IMPLY_DEDENT == 0) {
			p.tok.pendin = -p.tok.indent
			p.tok.indent = 0
		}
	} else {
		p.parsing_started = true
	}

	t := &Token{}
	p.tokens = append(p.tokens, t)
	return initialize_token(p, t, &new_token, _type)
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
