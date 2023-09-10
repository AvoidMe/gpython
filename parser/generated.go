// @generated by pegen from python.gram
package parser

import (
	"github.com/AvoidMe/gpython/builtin"
)

var reserved_keywords = [][]KeywordToken{
	{},
	{},
	{},
	{},
	{
		{"pass", 500},
		{"True", 501},
		{"None", 503},
	},
	{
		{"False", 502},
	},
}
var soft_keywords = []string{}
var (
	file_type                  = 1000
	statements_type            = 1001
	statement_type             = 1002
	simple_stmts_type          = 1003
	simple_stmt_type           = 1004
	assignment_type            = 1005
	annotated_rhs_type         = 1006
	expression_type            = 1007
	star_expressions_type      = 1008
	star_expression_type       = 1009
	disjunction_type           = 1010
	conjunction_type           = 1011
	inversion_type             = 1012
	comparison_type            = 1013
	bitwise_or_type            = 1014
	bitwise_xor_type           = 1015
	bitwise_and_type           = 1016
	shift_expr_type            = 1017
	sum_type                   = 1018
	term_type                  = 1019
	factor_type                = 1020
	power_type                 = 1021
	await_primary_type         = 1022
	primary_type               = 1023
	atom_type                  = 1024
	strings_type               = 1025
	star_targets_type          = 1026
	star_target_type           = 1027
	target_with_star_atom_type = 1028
	star_atom_type             = 1029
	_loop1_1_type              = 1030
	_loop0_3_type              = 1031
	_gather_2_type             = 1032
	_loop1_4_type              = 1033
	_loop1_5_type              = 1034
	_tmp_6_type                = 1035
)

// file: statements? $
func file_rule(p *Parser) mod_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res mod_ty
	_mark := p.mark
	{ // statements? $
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a any
		var endmarker_var *Token
		condition := func() bool {
			a = statements_rule(p)
			if p.error_indicator != 0 {
				return false
			}
			// statements?
			endmarker_var = _PyPegen_expect_token(p, ENDMARKER)
			if endmarker_var == nil {
				return false
			} // token='ENDMARKER'
			return true
		}
		if condition() {
			_res = _PyPegen_make_module(p, a)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// statements: statement+
func statements_rule(p *Parser) *asdl_stmt_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res *asdl_stmt_seq
	_mark := p.mark
	{ // statement+
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a *asdl_seq
		condition := func() bool {
			a = _loop1_1_rule(p)
			if a == nil {
				return false
			} // statement+
			return true
		}
		if condition() {
			_res = _PyPegen_seq_flatten(p, a)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// statement: simple_stmts
func statement_rule(p *Parser) *asdl_stmt_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res *asdl_stmt_seq
	_mark := p.mark
	{ // simple_stmts
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a *asdl_stmt_seq
		condition := func() bool {
			a = (*asdl_stmt_seq)(simple_stmts_rule(p))
			if a == nil {
				return false
			} // simple_stmts
			return true
		}
		if condition() {
			_res = a
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// simple_stmts: ';'.simple_stmt+ ';'? NEWLINE
func simple_stmts_rule(p *Parser) *asdl_stmt_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res *asdl_stmt_seq
	_mark := p.mark
	{ // ';'.simple_stmt+ ';'? NEWLINE
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _opt_var any
		UNUSED(_opt_var) // Silence compiler warnings
		var a *asdl_stmt_seq
		var newline_var *Token
		condition := func() bool {
			a = (*asdl_stmt_seq)(_gather_2_rule(p))
			if a == nil {
				return false
			} // ';'.simple_stmt+
			_opt_var = _PyPegen_expect_token(p, 13)
			if p.error_indicator != 0 {
				return false
			}
			// ';'?
			newline_var = _PyPegen_expect_token(p, NEWLINE)
			if newline_var == nil {
				return false
			} // token='NEWLINE'
			return true
		}
		if condition() {
			_res = a
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// simple_stmt: assignment | star_expressions | 'pass'
func simple_stmt_rule(p *Parser) stmt_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res stmt_ty
	if _PyPegen_is_memoized(p, simple_stmt_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	if p.mark == p.fill && _PyPegen_fill_token(p) < 0 {
		p.error_indicator = 1
		p.level--
		return nil
	}
	_start_lineno := p.tokens[_mark].lineno
	_start_col_offset := p.tokens[_mark].col_offset
	{ // assignment
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var assignment_var stmt_ty
		condition := func() bool {
			assignment_var = assignment_rule(p)
			if assignment_var == nil {
				return false
			} // assignment
			return true
		}
		if condition() {
			_res = assignment_var
			goto done
		}
		p.mark = _mark
	}
	{ // star_expressions
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var e expr_ty
		condition := func() bool {
			e = star_expressions_rule(p)
			if e == nil {
				return false
			} // star_expressions
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Expr(e, _start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	{ // 'pass'
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _keyword *Token
		condition := func() bool {
			_keyword = _PyPegen_expect_token(p, 500)
			if _keyword == nil {
				return false
			} // token='pass'
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Pass(_start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, simple_stmt_type, _res)
	p.level--
	return _res
}

// assignment: ((star_targets '='))+ (star_expressions) !'=' TYPE_COMMENT?
func assignment_rule(p *Parser) stmt_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res stmt_ty
	_mark := p.mark
	if p.mark == p.fill && _PyPegen_fill_token(p) < 0 {
		p.error_indicator = 1
		p.level--
		return nil
	}
	_start_lineno := p.tokens[_mark].lineno
	_start_col_offset := p.tokens[_mark].col_offset
	{ // ((star_targets '='))+ (star_expressions) !'=' TYPE_COMMENT?
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a *asdl_expr_seq
		var b expr_ty
		var tc any
		condition := func() bool {
			a = (*asdl_expr_seq)(_loop1_4_rule(p))
			if a == nil {
				return false
			} // ((star_targets '='))+
			b = star_expressions_rule(p)
			if b == nil {
				return false
			} // star_expressions
			_tmp := _PyPegen_lookahead_with_int(0, _PyPegen_expect_token, p, 22)
			if _tmp == 0 {
				return false
			} // token='='
			tc = _PyPegen_expect_token(p, TYPE_COMMENT)
			if p.error_indicator != 0 {
				return false
			}
			// TYPE_COMMENT?
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Assign(a, b, NEW_TYPE_COMMENT(p, tc), _start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// annotated_rhs: star_expressions
func annotated_rhs_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // star_expressions
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var star_expressions_var expr_ty
		condition := func() bool {
			star_expressions_var = star_expressions_rule(p)
			if star_expressions_var == nil {
				return false
			} // star_expressions
			return true
		}
		if condition() {
			_res = star_expressions_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// expression: disjunction
func expression_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, expression_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // disjunction
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var disjunction_var expr_ty
		condition := func() bool {
			disjunction_var = disjunction_rule(p)
			if disjunction_var == nil {
				return false
			} // disjunction
			return true
		}
		if condition() {
			_res = disjunction_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, expression_type, _res)
	p.level--
	return _res
}

// star_expressions: star_expression
func star_expressions_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // star_expression
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var star_expression_var expr_ty
		condition := func() bool {
			star_expression_var = star_expression_rule(p)
			if star_expression_var == nil {
				return false
			} // star_expression
			return true
		}
		if condition() {
			_res = star_expression_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// star_expression: expression
func star_expression_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, star_expression_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // expression
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var expression_var expr_ty
		condition := func() bool {
			expression_var = expression_rule(p)
			if expression_var == nil {
				return false
			} // expression
			return true
		}
		if condition() {
			_res = expression_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, star_expression_type, _res)
	p.level--
	return _res
}

// disjunction: conjunction
func disjunction_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, disjunction_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // conjunction
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var conjunction_var expr_ty
		condition := func() bool {
			conjunction_var = conjunction_rule(p)
			if conjunction_var == nil {
				return false
			} // conjunction
			return true
		}
		if condition() {
			_res = conjunction_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, disjunction_type, _res)
	p.level--
	return _res
}

// conjunction: inversion
func conjunction_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, conjunction_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // inversion
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var inversion_var expr_ty
		condition := func() bool {
			inversion_var = inversion_rule(p)
			if inversion_var == nil {
				return false
			} // inversion
			return true
		}
		if condition() {
			_res = inversion_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, conjunction_type, _res)
	p.level--
	return _res
}

// inversion: comparison
func inversion_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, inversion_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // comparison
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var comparison_var expr_ty
		condition := func() bool {
			comparison_var = comparison_rule(p)
			if comparison_var == nil {
				return false
			} // comparison
			return true
		}
		if condition() {
			_res = comparison_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, inversion_type, _res)
	p.level--
	return _res
}

// comparison: bitwise_or
func comparison_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // bitwise_or
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var bitwise_or_var expr_ty
		condition := func() bool {
			bitwise_or_var = bitwise_or_rule(p)
			if bitwise_or_var == nil {
				return false
			} // bitwise_or
			return true
		}
		if condition() {
			_res = bitwise_or_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// bitwise_or: bitwise_xor
func bitwise_or_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // bitwise_xor
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var bitwise_xor_var expr_ty
		condition := func() bool {
			bitwise_xor_var = bitwise_xor_rule(p)
			if bitwise_xor_var == nil {
				return false
			} // bitwise_xor
			return true
		}
		if condition() {
			_res = bitwise_xor_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// bitwise_xor: bitwise_and
func bitwise_xor_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // bitwise_and
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var bitwise_and_var expr_ty
		condition := func() bool {
			bitwise_and_var = bitwise_and_rule(p)
			if bitwise_and_var == nil {
				return false
			} // bitwise_and
			return true
		}
		if condition() {
			_res = bitwise_and_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// bitwise_and: shift_expr
func bitwise_and_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // shift_expr
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var shift_expr_var expr_ty
		condition := func() bool {
			shift_expr_var = shift_expr_rule(p)
			if shift_expr_var == nil {
				return false
			} // shift_expr
			return true
		}
		if condition() {
			_res = shift_expr_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// shift_expr: sum
func shift_expr_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // sum
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var sum_var expr_ty
		condition := func() bool {
			sum_var = sum_rule(p)
			if sum_var == nil {
				return false
			} // sum
			return true
		}
		if condition() {
			_res = sum_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// sum: term
func sum_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // term
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var term_var expr_ty
		condition := func() bool {
			term_var = term_rule(p)
			if term_var == nil {
				return false
			} // term
			return true
		}
		if condition() {
			_res = term_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// term: factor
func term_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // factor
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var factor_var expr_ty
		condition := func() bool {
			factor_var = factor_rule(p)
			if factor_var == nil {
				return false
			} // factor
			return true
		}
		if condition() {
			_res = factor_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// factor: power
func factor_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, factor_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // power
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var power_var expr_ty
		condition := func() bool {
			power_var = power_rule(p)
			if power_var == nil {
				return false
			} // power
			return true
		}
		if condition() {
			_res = power_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, factor_type, _res)
	p.level--
	return _res
}

// power: await_primary
func power_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // await_primary
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var await_primary_var expr_ty
		condition := func() bool {
			await_primary_var = await_primary_rule(p)
			if await_primary_var == nil {
				return false
			} // await_primary
			return true
		}
		if condition() {
			_res = await_primary_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// await_primary: primary
func await_primary_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, await_primary_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // primary
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var primary_var expr_ty
		condition := func() bool {
			primary_var = primary_rule(p)
			if primary_var == nil {
				return false
			} // primary
			return true
		}
		if condition() {
			_res = primary_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, await_primary_type, _res)
	p.level--
	return _res
}

// primary: atom
func primary_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // atom
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var atom_var expr_ty
		condition := func() bool {
			atom_var = atom_rule(p)
			if atom_var == nil {
				return false
			} // atom
			return true
		}
		if condition() {
			_res = atom_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// atom: NAME | 'True' | 'False' | 'None' | &STRING strings | NUMBER
func atom_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	if p.mark == p.fill && _PyPegen_fill_token(p) < 0 {
		p.error_indicator = 1
		p.level--
		return nil
	}
	_start_lineno := p.tokens[_mark].lineno
	_start_col_offset := p.tokens[_mark].col_offset
	{ // NAME
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var name_var expr_ty
		condition := func() bool {
			name_var = _PyPegen_name_token(p)
			if name_var == nil {
				return false
			} // NAME
			return true
		}
		if condition() {
			_res = name_var
			goto done
		}
		p.mark = _mark
	}
	{ // 'True'
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _keyword *Token
		condition := func() bool {
			_keyword = _PyPegen_expect_token(p, 501)
			if _keyword == nil {
				return false
			} // token='True'
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Constant(builtin.PyTrue, nil, _start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	{ // 'False'
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _keyword *Token
		condition := func() bool {
			_keyword = _PyPegen_expect_token(p, 502)
			if _keyword == nil {
				return false
			} // token='False'
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Constant(builtin.PyFalse, nil, _start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	{ // 'None'
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _keyword *Token
		condition := func() bool {
			_keyword = _PyPegen_expect_token(p, 503)
			if _keyword == nil {
				return false
			} // token='None'
			return true
		}
		if condition() {
			_token := _PyPegen_get_last_nonnwhitespace_token(p)
			if _token == nil {
				p.level--
				return nil
			}
			_end_lineno := _token.end_lineno
			_end_col_offset := _token.end_col_offset
			_res = _PyAST_Constant(builtin.PyNone, nil, _start_lineno, _start_col_offset, _end_lineno, _end_col_offset)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	{ // &STRING strings
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var strings_var expr_ty
		condition := func() bool {
			_tmp := _PyPegen_lookahead(1, _PyPegen_string_token, p)
			if _tmp == 0 {
				return false
			}
			strings_var = strings_rule(p)
			if strings_var == nil {
				return false
			} // strings
			return true
		}
		if condition() {
			_res = strings_var
			goto done
		}
		p.mark = _mark
	}
	{ // NUMBER
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var number_var expr_ty
		condition := func() bool {
			number_var = _PyPegen_number_token(p)
			if number_var == nil {
				return false
			} // NUMBER
			return true
		}
		if condition() {
			_res = number_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// strings: STRING+
func strings_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, strings_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // STRING+
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a *asdl_seq
		condition := func() bool {
			a = _loop1_5_rule(p)
			if a == nil {
				return false
			} // STRING+
			return true
		}
		if condition() {
			_res = _PyPegen_concatenate_strings(p, a)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, strings_type, _res)
	p.level--
	return _res
}

// star_targets: star_target !','
func star_targets_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // star_target !','
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a expr_ty
		condition := func() bool {
			a = star_target_rule(p)
			if a == nil {
				return false
			} // star_target
			_tmp := _PyPegen_lookahead_with_int(0, _PyPegen_expect_token, p, 12)
			if _tmp == 0 {
				return false
			} // token=','
			return true
		}
		if condition() {
			_res = a
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// star_target: target_with_star_atom
func star_target_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, star_target_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // target_with_star_atom
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var target_with_star_atom_var expr_ty
		condition := func() bool {
			target_with_star_atom_var = target_with_star_atom_rule(p)
			if target_with_star_atom_var == nil {
				return false
			} // target_with_star_atom
			return true
		}
		if condition() {
			_res = target_with_star_atom_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, star_target_type, _res)
	p.level--
	return _res
}

// target_with_star_atom: star_atom
func target_with_star_atom_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	if _PyPegen_is_memoized(p, target_with_star_atom_type, &_res) {
		p.level--
		return _res
	}
	_mark := p.mark
	{ // star_atom
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var star_atom_var expr_ty
		condition := func() bool {
			star_atom_var = star_atom_rule(p)
			if star_atom_var == nil {
				return false
			} // star_atom
			return true
		}
		if condition() {
			_res = star_atom_var
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	_PyPegen_insert_memo(p, _mark, target_with_star_atom_type, _res)
	p.level--
	return _res
}

// star_atom: NAME
func star_atom_rule(p *Parser) expr_ty {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res expr_ty
	_mark := p.mark
	{ // NAME
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var a expr_ty
		condition := func() bool {
			a = _PyPegen_name_token(p)
			if a == nil {
				return false
			} // NAME
			return true
		}
		if condition() {
			_res = _PyPegen_set_expr_context(p, a, Store)
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// _loop1_1: statement
func _loop1_1_rule(p *Parser) *asdl_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res any
	_mark := p.mark
	_start_mark := p.mark
	_children := []any{}
	_n := 0
	{ // statement
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var statement_var *asdl_stmt_seq
		condition := func() bool {
			statement_var = statement_rule(p)
			if statement_var == nil {
				return false
			} // statement
			return true
		}
		for condition() {
			_res = statement_var
			_children = append(_children, _res)
			_mark = p.mark
		}
		p.mark = _mark
	}
	if _n == 0 || p.error_indicator > 0 {
		p.level--
		return nil
	}
	_seq := _Py_asdl_generic_seq_new(_n, p.arena)
	for i := 0; i < _n; i++ {
		asdl_seq_SET_UNTYPED(_seq, i, _children[i])
	}
	_PyPegen_insert_memo(p, _start_mark, _loop1_1_type, _seq)
	p.level--
	return _seq
}

// _loop0_3: ';' simple_stmt
func _loop0_3_rule(p *Parser) *asdl_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res any
	_mark := p.mark
	_start_mark := p.mark
	_children := []any{}
	_n := 0
	{ // ';' simple_stmt
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _literal *Token
		var elem stmt_ty
		condition := func() bool {
			_literal = _PyPegen_expect_token(p, 13)
			if _literal == nil {
				return false
			} // token=';'
			elem = simple_stmt_rule(p)
			if elem == nil {
				return false
			} // simple_stmt
			return true
		}
		for condition() {
			_res = elem
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			_children = append(_children, _res)
			_mark = p.mark
		}
		p.mark = _mark
	}
	_seq := _Py_asdl_generic_seq_new(_n, p.arena)
	for i := 0; i < _n; i++ {
		asdl_seq_SET_UNTYPED(_seq, i, _children[i])
	}
	_PyPegen_insert_memo(p, _start_mark, _loop0_3_type, _seq)
	p.level--
	return _seq
}

// _gather_2: simple_stmt _loop0_3
func _gather_2_rule(p *Parser) *asdl_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res *asdl_seq
	_mark := p.mark
	{ // simple_stmt _loop0_3
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var elem stmt_ty
		var seq *asdl_seq
		condition := func() bool {
			elem = simple_stmt_rule(p)
			if elem == nil {
				return false
			} // simple_stmt
			seq = _loop0_3_rule(p)
			if seq == nil {
				return false
			} // _loop0_3
			return true
		}
		if condition() {
			_res = _PyPegen_seq_insert_in_front(p, elem, seq)
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

// _loop1_4: (star_targets '=')
func _loop1_4_rule(p *Parser) *asdl_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res any
	_mark := p.mark
	_start_mark := p.mark
	_children := []any{}
	_n := 0
	{ // (star_targets '=')
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _tmp_6_var any
		condition := func() bool {
			_tmp_6_var = _tmp_6_rule(p)
			if _tmp_6_var == nil {
				return false
			} // star_targets '='
			return true
		}
		for condition() {
			_res = _tmp_6_var
			_children = append(_children, _res)
			_mark = p.mark
		}
		p.mark = _mark
	}
	if _n == 0 || p.error_indicator > 0 {
		p.level--
		return nil
	}
	_seq := _Py_asdl_generic_seq_new(_n, p.arena)
	for i := 0; i < _n; i++ {
		asdl_seq_SET_UNTYPED(_seq, i, _children[i])
	}
	_PyPegen_insert_memo(p, _start_mark, _loop1_4_type, _seq)
	p.level--
	return _seq
}

// _loop1_5: STRING
func _loop1_5_rule(p *Parser) *asdl_seq {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res any
	_mark := p.mark
	_start_mark := p.mark
	_children := []any{}
	_n := 0
	{ // STRING
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var string_var expr_ty
		condition := func() bool {
			string_var = _PyPegen_string_token(p)
			if string_var == nil {
				return false
			} // STRING
			return true
		}
		for condition() {
			_res = string_var
			_children = append(_children, _res)
			_mark = p.mark
		}
		p.mark = _mark
	}
	if _n == 0 || p.error_indicator > 0 {
		p.level--
		return nil
	}
	_seq := _Py_asdl_generic_seq_new(_n, p.arena)
	for i := 0; i < _n; i++ {
		asdl_seq_SET_UNTYPED(_seq, i, _children[i])
	}
	_PyPegen_insert_memo(p, _start_mark, _loop1_5_type, _seq)
	p.level--
	return _seq
}

// _tmp_6: star_targets '='
func _tmp_6_rule(p *Parser) any {
	p.level++
	if p.error_indicator > 0 {
		p.level--
		return nil
	}
	var _res any
	_mark := p.mark
	{ // star_targets '='
		if p.error_indicator > 0 {
			p.level--
			return nil
		}
		var _literal *Token
		var z expr_ty
		condition := func() bool {
			z = star_targets_rule(p)
			if z == nil {
				return false
			} // star_targets
			_literal = _PyPegen_expect_token(p, 22)
			if _literal == nil {
				return false
			} // token='='
			return true
		}
		if condition() {
			_res = z
			if _res == nil && PyErr_Occurred() {
				p.error_indicator = 1
				p.level--
				return nil
			}
			goto done
		}
		p.mark = _mark
	}
	_res = nil
done:
	p.level--
	return _res
}

func _PyPegen_parse(p *Parser) any {
	// Initialize keywords
	p.Keywords = reserved_keywords
	p.SoftKeywords = soft_keywords

	// Run parser
	var result *AST
	switch p.StartRule {
	case Py_file_input:
		return file_rule(p)
	}

	return result
}
