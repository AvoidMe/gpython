package parser

func tok_get(tok *TokenState, token *Token) int {
	var c int
	var blankline, nonascii int

	var p_start, p_end *string

nextline:
	tok.start = 0
	tok.starting_col_offset = -1
	blankline = 0

	/* Get indentation level */
	if tok.atbol > 0 {
		col := 0
		altcol := 0
		tok.atbol = 0
		cont_line_col := 0
		for {
			c = tok_nextc(tok)
			if c == ' ' {
				col++
				altcol++
			} else if c == '\t' {
				col = (col/tok.tabsize + 1) * tok.tabsize
				altcol = (altcol/ALTTABSIZE + 1) * ALTTABSIZE
			} else if c == '\014' { /* Control-L (formfeed) */
				col = 0
				altcol = 0 /* For Emacs users */
			} else if c == '\\' {
				// Indentation cannot be split over multiple physical lines
				// using backslashes. This means that if we found a backslash
				// preceded by whitespace, **the first one we find** determines
				// the level of indentation of whatever comes next.
				if cont_line_col == 0 {
					cont_line_col = col
				}
				if c = tok_continuation_line(tok); c == -1 {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			} else {
				break
			}
		}
		tok_backup(tok, c)
		if c == '#' || c == '\n' {
			/* Lines with only whitespace and/or comments
			   shouldn't affect the indentation and are
			   not passed to the parser as NEWLINE tokens,
			   except *totally* empty lines in interactive
			   mode, which signal the end of a command group. */
			if col == 0 && c == '\n' && tok.prompt != nil {
				blankline = 0 /* Let it through */
			} else if tok.prompt != nil && tok.lineno == 1 {
				/* In interactive mode, if the first line contains
				   only spaces and/or a comment, let it through. */
				blankline = 0
				col = 0
				altcol = 0
			} else {
				blankline = 1 /* Ignore completely */
			}
			/* We can't jump back right here since we still
			   may need to skip to the end of a comment */
		}
		if blankline != 0 && tok.level == 0 {
			if cont_line_col > 0 {
				col = cont_line_col
				altcol = cont_line_col
			}
			if col == tok.indstack[tok.indent] {
				/* No change */
				if altcol != tok.altindstack[tok.indent] {
					return MAKE_TOKEN(indenterror(tok))
				}
			} else if col > tok.indstack[tok.indent] {
				/* Indent -- always one */
				if tok.indent+1 >= MAXINDENT {
					tok.done = E_TOODEEP
					tok.cur = tok.inp
					return MAKE_TOKEN(ERRORTOKEN)
				}
				if altcol <= tok.altindstack[tok.indent] {
					return MAKE_TOKEN(indenterror(tok))
				}
				tok.pendin++
				tok.indent++
				tok.indstack[tok.indent] = col
				tok.altindstack[tok.indent] = altcol
			} else /* col < tok.indstack[tok.indent] */ {
				/* Dedent -- any number, must be consistent */
				for tok.indent > 0 &&
					col < tok.indstack[tok.indent] {
					tok.pendin--
					tok.indent--
				}
				if col != tok.indstack[tok.indent] {
					tok.done = E_DEDENT
					tok.cur = tok.inp
					return MAKE_TOKEN(ERRORTOKEN)
				}
				if altcol != tok.altindstack[tok.indent] {
					return MAKE_TOKEN(indenterror(tok))
				}
			}
		}
	}

	tok.start = tok.cur
	tok.starting_col_offset = tok.col_offset

	/* Return pending indents/dedents */
	if tok.pendin != 0 {
		if tok.pendin < 0 {
			tok.pendin++
			return MAKE_TOKEN(DEDENT)
		} else {
			tok.pendin--
			return MAKE_TOKEN(INDENT)
		}
	}

	/* Peek ahead at the next character */
	c = tok_nextc(tok)
	tok_backup(tok, c)
	/* Check if we are closing an async function */
	if tok.async_def && blankline != 0 &&
		/* Due to some implementation artifacts of type comments,
		 * a TYPE_COMMENT at the start of a function won't set an
		 * indentation level and it will produce a NEWLINE after it.
		 * To avoid spuriously ending an async function due to this,
		 * wait until we have some non-newline char in front of us. */
		c != '\n' &&
		tok.level == 0 &&
		/* There was a NEWLINE after ASYNC DEF,
		   so we're past the signature. */
		tok.async_def_nl != 0 &&
		/* Current indentation level is less than where
		   the async function was defined */
		tok.async_def_indent >= tok.indent {
		tok.async_def = false
		tok.async_def_indent = 0
		tok.async_def_nl = 0
	}

again:
	tok.start = 0
	/* Skip spaces */
	for {
		c = tok_nextc(tok)
		if c == ' ' || c == '\t' || c == '\014' {
			break
		}
	}

	/* Set start of current token */
	if tok.cur == 0 {
		tok.start = 0
	} else {
		tok.start = tok.cur - 1
	}
	tok.starting_col_offset = tok.col_offset - 1

	/* Skip comment, unless it's a type comment */
	if c == '#' {
		var prefix, p, type_start *byte
		var current_starting_col_offset int

		for c != EOF && c != '\n' {
			c = tok_nextc(tok)
		}

		if tok.type_comments {
			p = tok.start
			current_starting_col_offset = tok.starting_col_offset
			prefix = type_comment_prefix
			for *prefix && p < tok.cur {
				if *prefix == ' ' {
					for *p == ' ' || *p == '\t' {
						p++
						current_starting_col_offset++
					}
				} else if *prefix == *p {
					p++
					current_starting_col_offset++
				} else {
					break
				}

				prefix++
			}

			/* This is a type comment if we matched all of type_comment_prefix. */
			if !*prefix {
				is_type_ignore := 1
				// +6 in order to skip the word 'ignore'
				const char *ignore_end = p + 6
				const int ignore_end_col_offset = current_starting_col_offset + 6
				tok_backup(tok, c) /* don't eat the newline or EOF */

				type_start = p

				/* A TYPE_IGNORE is "type: ignore" followed by the end of the token
				 * or anything ASCII and non-alphanumeric. */
				is_type_ignore = (tok.cur >= ignore_end && memcmp(p, "ignore", 6) == 0 &&
					!(tok.cur > ignore_end &&
						(ignore_end[0] >= 128 || Py_ISALNUM(ignore_end[0]))))

				if is_type_ignore {
					p_start = ignore_end
					p_end = tok.cur

					/* If this type ignore is the only thing on the line, consume the newline also. */
					if blankline {
						tok_nextc(tok)
						tok.atbol = 1
					}
					return MAKE_TYPE_COMMENT_TOKEN(TYPE_IGNORE, ignore_end_col_offset, tok.col_offset)
				} else {
					p_start = type_start
					p_end = tok.cur
					return MAKE_TYPE_COMMENT_TOKEN(TYPE_COMMENT, current_starting_col_offset, tok.col_offset)
				}
			}
		}
	}

	if tok.done == E_INTERACT_STOP {
		return MAKE_TOKEN(ENDMARKER)
	}

	/* Check for EOF and errors now */
	if c == EOF {
		if tok.level {
			return MAKE_TOKEN(ERRORTOKEN)
		}
		if tok.done == E_EOF {
			return MAKE_TOKEN(ENDMARKER)
		}
		return ERRORTOKEN
	}

	/* Identifier (most frequent token!) */
	nonascii = 0
	if is_potential_identifier_start(c) {
		/* Process the various legal combinations of b"", r"", u"", and f"". */
		saw_b := 0
		saw_r := 0
		saw_u := 0
		saw_f := 0
		for {
			if !(saw_b > 0 || saw_u > 0 || saw_f > 0) && (c == 'b' || c == 'B') {
				saw_b = 1
			} else if !(saw_b > 0 || saw_u > 0 || saw_r > 0 || saw_f > 0) && (c == 'u' || c == 'U') {
				/* Since this is a backwards compatibility support literal we don't
				   want to support it in arbitrary order like byte literals. */
				saw_u = 1
			} else if !(saw_r > 0 || saw_u > 0) && (c == 'r' || c == 'R') {
				/* ur"" and ru"" are not supported */
				saw_r = 1
			} else if !(saw_f > 0 || saw_b > 0 || saw_u > 0) && (c == 'f' || c == 'F') {
				saw_f = 1
			} else {
				break
			}
			c = tok_nextc(tok)
			if c == '"' || c == '\'' {
				goto letter_quote
			}
		}
		for is_potential_identifier_char(c) {
			if c >= 128 {
				nonascii = 1
			}
			c = tok_nextc(tok)
		}
		tok_backup(tok, c)
		if nonascii && !verify_identifier(tok) {
			return MAKE_TOKEN(ERRORTOKEN)
		}

		p_start = tok.start
		p_end = tok.cur

		/* async/await parsing block. */
		if tok.cur-tok.start == 5 && tok.start[0] == 'a' {
			/* May be an 'async' or 'await' token.  For Python 3.7 or
			   later we recognize them unconditionally.  For Python
			   3.5 or 3.6 we recognize 'async' in front of 'def', and
			   either one inside of 'async def'.  (Technically we
			   shouldn't recognize these at all for 3.4 or earlier,
			   but there's no *valid* Python 3.4 code that would be
			   rejected, and async functions will be rejected in a
			   later phase.) */
			if !tok.async_hacks || tok.async_def {
				/* Always recognize the keywords. */
				if memcmp(tok.start, "async", 5) == 0 {
					return MAKE_TOKEN(ASYNC)
				}
				if memcmp(tok.start, "await", 5) == 0 {
					return MAKE_TOKEN(AWAIT)
				}
			} else if memcmp(tok.start, "async", 5) == 0 {
				/* The current token is 'async'.
				   Look ahead one token to see if that is 'def'. */

				ahead_tok := TokenState{}
				ahead_token := Token{}
				ahead_token_kind := 0

				memcpy(&ahead_tok, tok, sizeof(ahead_tok))
				ahead_tok_kind = tok_get(&ahead_tok, &ahead_token)

				if ahead_tok_kind == NAME &&
					ahead_tok.cur-ahead_tok.start == 3 &&
					memcmp(ahead_tok.start, "def", 3) == 0 {
					/* The next token is going to be 'def', so instead of
					   returning a plain NAME token, return ASYNC. */
					tok.async_def_indent = tok.indent
					tok.async_def = 1
					return MAKE_TOKEN(ASYNC)
				}
			}
		}

		return MAKE_TOKEN(NAME)
	}

	/* Newline */
	if c == '\n' {
		tok.atbol = 1
		if blankline || tok.level > 0 {
			goto nextline
		}
		p_start = tok.start
		p_end = tok.cur - 1 /* Leave '\n' out of the string */
		tok.cont_line = 0
		if tok.async_def {
			/* We're somewhere inside an 'async def' function, and
			   we've encountered a NEWLINE after its signature. */
			tok.async_def_nl = 1
		}
		return MAKE_TOKEN(NEWLINE)
	}

	/* Period or number starting with period? */
	if c == '.' {
		c = tok_nextc(tok)
		if isdigit(c) {
			goto fraction
		} else if c == '.' {
			c = tok_nextc(tok)
			if c == '.' {
				p_start = tok.start
				p_end = tok.cur
				return MAKE_TOKEN(ELLIPSIS)
			} else {
				tok_backup(tok, c)
			}
			tok_backup(tok, '.')
		} else {
			tok_backup(tok, c)
		}
		p_start = tok.start
		p_end = tok.cur
		return MAKE_TOKEN(DOT)
	}

	/* Number */
	if isdigit(c) {
		if c == '0' {
			/* Hex, octal or binary -- maybe. */
			c = tok_nextc(tok)
			if c == 'x' || c == 'X' {
				/* Hex */
				c = tok_nextc(tok)
				for {
					if c == '_' {
						c = tok_nextc(tok)
					}
					if !isxdigit(c) {
						tok_backup(tok, c)
						return MAKE_TOKEN(syntaxerror(tok, "invalid hexadecimal literal"))
					}
					for {
						c = tok_nextc(tok)
						if !isxdigit(c) {
							break
						}
					}
					if c == '_' {
						break
					}
				}
				if !verify_end_of_number(tok, c, "hexadecimal") {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			} else if c == 'o' || c == 'O' {
				/* Octal */
				c = tok_nextc(tok)
				for {
					if c == '_' {
						c = tok_nextc(tok)
					}
					if c < '0' || c >= '8' {
						if isdigit(c) {
							return MAKE_TOKEN(syntaxerror(tok,
								"invalid digit '%c' in octal literal", c))
						} else {
							tok_backup(tok, c)
							return MAKE_TOKEN(syntaxerror(tok, "invalid octal literal"))
						}
					}
					for {
						c = tok_nextc(tok)
						if !('0' <= c && c < '8') {
							break
						}
					}
					if !(c == '_') {
						break
					}
				}
				if isdigit(c) {
					return MAKE_TOKEN(syntaxerror(tok,
						"invalid digit '%c' in octal literal", c))
				}
				if !verify_end_of_number(tok, c, "octal") {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			} else if c == 'b' || c == 'B' {
				/* Binary */
				c = tok_nextc(tok)
				for {
					if c == '_' {
						c = tok_nextc(tok)
					}
					if c != '0' && c != '1' {
						if isdigit(c) {
							return MAKE_TOKEN(syntaxerror(tok, "invalid digit '%c' in binary literal", c))
						} else {
							tok_backup(tok, c)
							return MAKE_TOKEN(syntaxerror(tok, "invalid binary literal"))
						}
					}
					for {
						c = tok_nextc(tok)
						if !(c == '0' || c == '1') {
							break
						}
					}
					if !(c == '_') {
						break
					}
				}
				if isdigit(c) {
					return MAKE_TOKEN(syntaxerror(tok, "invalid digit '%c' in binary literal", c))
				}
				if !verify_end_of_number(tok, c, "binary") {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			} else {
				nonzero := 0
				/* maybe old-style octal; c is first char of it */
				/* in any case, allow '0' as a literal */
				for {
					if c == '_' {
						c = tok_nextc(tok)
						if !isdigit(c) {
							tok_backup(tok, c)
							return MAKE_TOKEN(syntaxerror(tok, "invalid decimal literal"))
						}
					}
					if c != '0' {
						break
					}
					c = tok_nextc(tok)
				}
				zeros_end := tok.cur
				if isdigit(c) {
					nonzero = 1
					c = tok_decimal_tail(tok)
					if c == 0 {
						return MAKE_TOKEN(ERRORTOKEN)
					}
				}
				if c == '.' {
					c = tok_nextc(tok)
					goto fraction
				} else if c == 'e' || c == 'E' {
					goto exponent
				} else if c == 'j' || c == 'J' {
					goto imaginary
				} else if nonzero {
					/* Old-style octal: now disallowed. */
					tok_backup(tok, c)
					return MAKE_TOKEN(syntaxerror_known_range(
						tok, (tok.start + 1 - tok.line_start),
						(zeros_end - tok.line_start),
						"leading zeros in decimal integer "+
							"literals are not permitted; "+
							"use an 0o prefix for octal integers"))
				}
				if !verify_end_of_number(tok, c, "decimal") {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			}
		} else {
			/* Decimal */
			c = tok_decimal_tail(tok)
			if c == 0 {
				return MAKE_TOKEN(ERRORTOKEN)
			}
			{
				/* Accept floating point numbers. */
				if c == '.' {
					c = tok_nextc(tok)
				fraction:
					/* Fraction */
					if isdigit(c) {
						c = tok_decimal_tail(tok)
						if c == 0 {
							return MAKE_TOKEN(ERRORTOKEN)
						}
					}
				}
				if c == 'e' || c == 'E' {
					e := 0
				exponent:
					e = c
					/* Exponent part */
					c = tok_nextc(tok)
					if c == '+' || c == '-' {
						c = tok_nextc(tok)
						if !isdigit(c) {
							tok_backup(tok, c)
							return MAKE_TOKEN(syntaxerror(tok, "invalid decimal literal"))
						}
					} else if !isdigit(c) {
						tok_backup(tok, c)
						if !verify_end_of_number(tok, e, "decimal") {
							return MAKE_TOKEN(ERRORTOKEN)
						}
						tok_backup(tok, e)
						p_start = tok.start
						p_end = tok.cur
						return MAKE_TOKEN(NUMBER)
					}
					c = tok_decimal_tail(tok)
					if c == 0 {
						return MAKE_TOKEN(ERRORTOKEN)
					}
				}
				if c == 'j' || c == 'J' {
					/* Imaginary part */
				imaginary:
					c = tok_nextc(tok)
					if !verify_end_of_number(tok, c, "imaginary") {
						return MAKE_TOKEN(ERRORTOKEN)
					}
				} else if !verify_end_of_number(tok, c, "decimal") {
					return MAKE_TOKEN(ERRORTOKEN)
				}
			}
		}
		tok_backup(tok, c)
		p_start = tok.start
		p_end = tok.cur
		return MAKE_TOKEN(NUMBER)
	}

letter_quote:
	/* String */
	if c == '\'' || c == '"' {
		quote := c
		quote_size := 1 /* 1 or 3 */
		end_quote_size := 0

		/* Nodes of type STRING, especially multi line strings
		   must be handled differently in order to get both
		   the starting line number and the column offset right.
		   (cf. issue 16806) */
		tok.first_lineno = tok.lineno
		tok.multi_line_start = tok.line_start

		/* Find the quote size and start of string */
		c = tok_nextc(tok)
		if c == quote {
			c = tok_nextc(tok)
			if c == quote {
				quote_size = 3
			} else {
				end_quote_size = 1 /* empty string found */
			}
		}
		if c != quote {
			tok_backup(tok, c)
		}

		/* Get rest of string */
		for end_quote_size != quote_size {
			c = tok_nextc(tok)
			if tok.done == E_DECODE {
				break
			}
			if c == EOF || (quote_size == 1 && c == '\n') {
				assert(tok.multi_line_start != NULL)
				// shift the tok_state's location into
				// the start of string, and report the error
				// from the initial quote character
				tok.cur = tok.start
				tok.cur++
				tok.line_start = tok.multi_line_start
				start := tok.lineno
				tok.lineno = tok.first_lineno
				if quote_size == 3 {
					syntaxerror(tok, "unterminated triple-quoted string literal"+
						" (detected at line %d)", start)
					if c != '\n' {
						tok.done = E_EOFS
					}
					return MAKE_TOKEN(ERRORTOKEN)
				} else {
					syntaxerror(tok, "unterminated string literal (detected at"+
						" line %d)", start)
					if c != '\n' {
						tok.done = E_EOLS
					}
					return MAKE_TOKEN(ERRORTOKEN)
				}
			}
			if c == quote {
				end_quote_size += 1
			} else {
				end_quote_size = 0
				if c == '\\' {
					tok_nextc(tok) /* skip escaped char */
				}
			}
		}

		p_start = tok.start
		p_end = tok.cur
		return MAKE_TOKEN(STRING)
	}

	/* Line continuation */
	if c == '\\' {
		if c = tok_continuation_line(tok); c == -1 {
			return MAKE_TOKEN(ERRORTOKEN)
		}
		tok.cont_line = 1
		goto again /* Read next line */
	}

	/* Check for two-character token */
	{
		c2 := tok_nextc(tok)
		current_token := _PyToken_TwoChars(c, c2)
		if current_token != OP {
			c3 := tok_nextc(tok)
			current_token3 := _PyToken_ThreeChars(c, c2, c3)
			if current_token3 != OP {
				current_token = current_token3
			} else {
				tok_backup(tok, c3)
			}
			p_start = tok.start
			p_end = tok.cur
			return MAKE_TOKEN(current_token)
		}
		tok_backup(tok, c2)
	}

	/* Keep track of parentheses nesting level */
	switch c {
	case '(':
	case '[':
	case '{':
		if tok.level >= MAXLEVEL {
			return MAKE_TOKEN(syntaxerror(tok, "too many nested parentheses"))
		}
		tok.parenstack[tok.level] = c
		tok.parenlinenostack[tok.level] = tok.lineno
		tok.parencolstack[tok.level] = (int)(tok.start - tok.line_start)
		tok.level++
		break
	case ')':
	case ']':
	case '}':
		if !tok.level {
			return MAKE_TOKEN(syntaxerror(tok, "unmatched '%c'", c))
		}
		tok.level--
		opening := tok.parenstack[tok.level]
		if !((opening == '(' && c == ')') ||
			(opening == '[' && c == ']') ||
			(opening == '{' && c == '}')) {
			if tok.parenlinenostack[tok.level] != tok.lineno {
				return MAKE_TOKEN(syntaxerror(tok,
					"closing parenthesis '%c' does not match "+
						"opening parenthesis '%c' on line %d",
					c, opening, tok.parenlinenostack[tok.level]))
			} else {
				return MAKE_TOKEN(syntaxerror(tok,
					"closing parenthesis '%c' does not match "+
						"opening parenthesis '%c'",
					c, opening))
			}
		}
		break
	}

	if !Py_UNICODE_ISPRINTABLE(c) {
		var hex [9]byte
		PyOS_snprintf(hex, sizeof(hex), "%04X", c)
		return MAKE_TOKEN(syntaxerror(tok, "invalid non-printable character U+%s", hex))
	}

	/* Punctuation character */
	p_start = tok.start
	p_end = tok.cur
	return MAKE_TOKEN(_PyToken_OneChar(c))
}

func _PyTokenizer_Get(tok *TokenState, token *Token) int {
	result := tok_get(tok, token)
	if tok.decoding_erred {
		result = ERRORTOKEN
		tok.done = E_DECODE
	}
	return result
}
