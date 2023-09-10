import ast
import re
from dataclasses import dataclass, field
from enum import Enum
import os.path
import token
from typing import IO, Any, Dict, Optional, Sequence, Set, Text, Tuple, List

from pegen import grammar
from pegen.grammar import (
    Alt,
    Cut,
    Forced,
    Gather,
    GrammarVisitor,
    Group,
    Lookahead,
    NamedItem,
    NameLeaf,
    NegativeLookahead,
    Opt,
    PositiveLookahead,
    Repeat0,
    Repeat1,
    Rhs,
    Rule,
    Leaf,
    StringLeaf,
)
from pegen.parser_generator import ParserGenerator


class NodeTypes(Enum):
    NAME_TOKEN = 0
    NUMBER_TOKEN = 1
    STRING_TOKEN = 2
    GENERIC_TOKEN = 3
    KEYWORD = 4
    SOFT_KEYWORD = 5
    CUT_OPERATOR = 6


BASE_NODETYPES = {
    "NAME": NodeTypes.NAME_TOKEN,
    "NUMBER": NodeTypes.NUMBER_TOKEN,
    "STRING": NodeTypes.STRING_TOKEN,
    "SOFT_KEYWORD": NodeTypes.SOFT_KEYWORD,
}


EXTENSION_PREFIX = """\
package parser

import (
    "github.com/AvoidMe/gpython/builtin"
)
"""


EXTENSION_SUFFIX = """
"""


@dataclass
class FunctionCall:
    function: str
    arguments: List[Any] = field(default_factory=list)
    assigned_variable: Optional[str] = None
    assigned_variable_type: Optional[str] = None
    return_type: Optional[str] = None
    nodetype: Optional[NodeTypes] = None
    force_true: bool = False
    comment: Optional[str] = None

    def __str__(self) -> str:
        parts = []
        parts.append(self.function)
        if self.arguments:
            parts.append(f"({', '.join(map(str, self.arguments))})")
        if self.assigned_variable_type:
            parts.append(")")
        if self.force_true:
            parts.append("\nif p.error_indicator != 0 {\n")
            parts.append("  return false\n")
            parts.append("}\n")
        else:
            if self.assigned_variable:
                parts.append(f"\nif {self.assigned_variable} == nil {{ return false }}")
        if self.assigned_variable:
            if self.assigned_variable_type:
                parts = [
                    self.assigned_variable,
                    " = ",
                    "(",
                    self.assigned_variable_type,
                    ")(",
                    *parts,
                ]
            else:
                parts = [self.assigned_variable, " = ", *parts]
        else:
            parts = [
                "_tmp := ",
                *parts,
                "\nif _tmp == 0 { return false }",
            ]
        if self.comment:
            parts.append(f"  // {self.comment}")
        return "".join(parts)


class GoCallMakerVisitor(GrammarVisitor):
    def __init__(
        self,
        parser_generator: ParserGenerator,
        exact_tokens: Dict[str, int],
        non_exact_tokens: Set[str],
    ):
        self.gen = parser_generator
        self.exact_tokens = exact_tokens
        self.non_exact_tokens = non_exact_tokens
        self.cache: Dict[Any, FunctionCall] = {}
        self.cleanup_statements: List[str] = []

    def keyword_helper(self, keyword: str) -> FunctionCall:
        return FunctionCall(
            assigned_variable="_keyword",
            function="_PyPegen_expect_token",
            arguments=["p", self.gen.keywords[keyword]],
            return_type="*Token",
            nodetype=NodeTypes.KEYWORD,
            comment=f"token='{keyword}'",
        )

    def soft_keyword_helper(self, value: str) -> FunctionCall:
        return FunctionCall(
            assigned_variable="_keyword",
            function="_PyPegen_expect_soft_keyword",
            arguments=["p", value],
            return_type="expr_ty",
            nodetype=NodeTypes.SOFT_KEYWORD,
            comment=f"soft_keyword='{value}'",
        )

    def visit_NameLeaf(self, node: NameLeaf) -> FunctionCall:
        name = node.value
        if name in self.non_exact_tokens:
            if name in BASE_NODETYPES:
                return FunctionCall(
                    assigned_variable=f"{name.lower()}_var",
                    function=f"_PyPegen_{name.lower()}_token",
                    arguments=["p"],
                    nodetype=BASE_NODETYPES[name],
                    return_type="expr_ty",
                    comment=name,
                )
            return FunctionCall(
                assigned_variable=f"{name.lower()}_var",
                function="_PyPegen_expect_token",
                arguments=["p", name],
                nodetype=NodeTypes.GENERIC_TOKEN,
                return_type="*Token",
                comment=f"token='{name}'",
            )

        type = None
        rule = self.gen.all_rules.get(name.lower())
        if rule is not None:
            type = "*asdl_seq" if rule.is_loop() or rule.is_gather() else rule.type

        return FunctionCall(
            assigned_variable=f"{name}_var",
            function=f"{name}_rule",
            arguments=["p"],
            return_type=type,
            comment=f"{node}",
        )

    def visit_StringLeaf(self, node: StringLeaf) -> FunctionCall:
        val = ast.literal_eval(node.value)
        if re.match(r"[a-zA-Z_]\w*\Z", val):  # This is a keyword
            if node.value.endswith("'"):
                return self.keyword_helper(val)
            else:
                return self.soft_keyword_helper(node.value)
        else:
            assert val in self.exact_tokens, f"{node.value} is not a known literal"
            type = self.exact_tokens[val]
            return FunctionCall(
                assigned_variable="_literal",
                function="_PyPegen_expect_token",
                arguments=["p", type],
                nodetype=NodeTypes.GENERIC_TOKEN,
                return_type="*Token",
                comment=f"token='{val}'",
            )

    def visit_Rhs(self, node: Rhs) -> FunctionCall:
        if node in self.cache:
            return self.cache[node]
        if node.can_be_inlined:
            self.cache[node] = self.generate_call(node.alts[0].items[0])
        else:
            name = self.gen.artifical_rule_from_rhs(node)
            self.cache[node] = FunctionCall(
                assigned_variable=f"{name}_var",
                function=f"{name}_rule",
                arguments=["p"],
                comment=f"{node}",
            )
        return self.cache[node]

    def visit_NamedItem(self, node: NamedItem) -> FunctionCall:
        call = self.generate_call(node.item)
        if node.name:
            call.assigned_variable = node.name
        if node.type:
            call.assigned_variable_type = node.type
        return call

    def lookahead_call_helper(self, node: Lookahead, positive: int) -> FunctionCall:
        call = self.generate_call(node.node)
        if call.nodetype == NodeTypes.NAME_TOKEN:
            return FunctionCall(
                function="_PyPegen_lookahead_with_name",
                arguments=[positive, call.function, *call.arguments],
                return_type="int",
            )
        elif call.nodetype == NodeTypes.SOFT_KEYWORD:
            return FunctionCall(
                function="_PyPegen_lookahead_with_string",
                arguments=[positive, call.function, *call.arguments],
                return_type="int",
            )
        elif call.nodetype in {NodeTypes.GENERIC_TOKEN, NodeTypes.KEYWORD}:
            return FunctionCall(
                function="_PyPegen_lookahead_with_int",
                arguments=[positive, call.function, *call.arguments],
                return_type="int",
                comment=f"token={node.node}",
            )
        else:
            return FunctionCall(
                function="_PyPegen_lookahead",
                arguments=[positive, call.function, *call.arguments],
                return_type="int",
            )

    def visit_PositiveLookahead(self, node: PositiveLookahead) -> FunctionCall:
        return self.lookahead_call_helper(node, 1)

    def visit_NegativeLookahead(self, node: NegativeLookahead) -> FunctionCall:
        return self.lookahead_call_helper(node, 0)

    def visit_Forced(self, node: Forced) -> FunctionCall:
        call = self.generate_call(node.node)
        if isinstance(node.node, Leaf):
            assert isinstance(node.node, Leaf)
            val = ast.literal_eval(node.node.value)
            assert val in self.exact_tokens, f"{node.node.value} is not a known literal"
            type = self.exact_tokens[val]
            return FunctionCall(
                assigned_variable="_literal",
                function="_PyPegen_expect_forced_token",
                arguments=["p", type, f'"{val}"'],
                nodetype=NodeTypes.GENERIC_TOKEN,
                return_type="*Token",
                comment=f"forced_token='{val}'",
            )
        if isinstance(node.node, Group):
            call = self.visit(node.node.rhs)
            call.assigned_variable = None
            call.comment = None
            return FunctionCall(
                assigned_variable="_literal",
                function="_PyPegen_expect_forced_result",
                arguments=["p", str(call), f'"{node.node.rhs!s}"'],
                return_type="any",
                comment=f"forced_token=({node.node.rhs!s})",
            )
        else:
            raise NotImplementedError(f"Forced tokens don't work with {node.node} nodes")

    def visit_Opt(self, node: Opt) -> FunctionCall:
        call = self.generate_call(node.node)
        return FunctionCall(
            assigned_variable="_opt_var",
            function=call.function,
            arguments=call.arguments,
            force_true=True,
            comment=f"{node}",
        )

    def visit_Repeat0(self, node: Repeat0) -> FunctionCall:
        if node in self.cache:
            return self.cache[node]
        name = self.gen.artificial_rule_from_repeat(node.node, False)
        self.cache[node] = FunctionCall(
            assigned_variable=f"{name}_var",
            function=f"{name}_rule",
            arguments=["p"],
            return_type="*asdl_seq",
            comment=f"{node}",
        )
        return self.cache[node]

    def visit_Repeat1(self, node: Repeat1) -> FunctionCall:
        if node in self.cache:
            return self.cache[node]
        name = self.gen.artificial_rule_from_repeat(node.node, True)
        self.cache[node] = FunctionCall(
            assigned_variable=f"{name}_var",
            function=f"{name}_rule",
            arguments=["p"],
            return_type="*asdl_seq",
            comment=f"{node}",
        )
        return self.cache[node]

    def visit_Gather(self, node: Gather) -> FunctionCall:
        if node in self.cache:
            return self.cache[node]
        name = self.gen.artifical_rule_from_gather(node)
        self.cache[node] = FunctionCall(
            assigned_variable=f"{name}_var",
            function=f"{name}_rule",
            arguments=["p"],
            return_type="*asdl_seq",
            comment=f"{node}",
        )
        return self.cache[node]

    def visit_Group(self, node: Group) -> FunctionCall:
        return self.generate_call(node.rhs)

    def visit_Cut(self, node: Cut) -> FunctionCall:
        return FunctionCall(
            assigned_variable="_cut_var",
            return_type="int",
            function="1",
            nodetype=NodeTypes.CUT_OPERATOR,
        )

    def generate_call(self, node: Any) -> FunctionCall:
        return super().visit(node)


class GoParserGenerator(ParserGenerator, GrammarVisitor):
    def __init__(
        self,
        grammar: grammar.Grammar,
        tokens: Dict[int, str],
        exact_tokens: Dict[str, int],
        non_exact_tokens: Set[str],
        file: Optional[IO[Text]],
        debug: bool = False,
        skip_actions: bool = False,
    ):
        super().__init__(grammar, set(tokens.values()), file)
        self.callmakervisitor: GoCallMakerVisitor = GoCallMakerVisitor(
            self, exact_tokens, non_exact_tokens
        )
        self._varname_counter = 0
        self.debug = debug
        self.skip_actions = skip_actions
        self.cleanup_statements: List[str] = []

    def generate(self, filename: str) -> None:
        self.collect_rules()
        basename = os.path.basename(filename)
        self.print(f"// @generated by pegen from {basename}")
        header = self.grammar.metas.get("header", EXTENSION_PREFIX)
        if header:
            self.print(header)
        subheader = self.grammar.metas.get("subheader", "")
        if subheader:
            self.print(subheader)
        self._setup_keywords()
        self._setup_soft_keywords()
        self.print("var (")
        with self.indent():
            for i, (rulename, rule) in enumerate(self.all_rules.items(), 1000):
                comment = "  // Left-recursive" if rule.left_recursive else ""
                self.print(f"{rulename}_type = {i}{comment}")
        self.print(")")
        for rulename, rule in list(self.all_rules.items()):
            self.print()
            if rule.left_recursive:
                self.print("// Left-recursive")
            self.visit(rule)
        if self.skip_actions:
            mode = 0
        else:
            mode = int(self.rules["start"].type == "mod_ty") if "start" in self.rules else 1
            if mode == 1 and self.grammar.metas.get("bytecode"):
                mode += 1
        modulename = self.grammar.metas.get("modulename", "parse")
        trailer = self.grammar.metas.get("trailer", EXTENSION_SUFFIX)
        if trailer:
            self.print(trailer.rstrip("\n") % dict(mode=mode, modulename=modulename))

    def _group_keywords_by_length(self) -> Dict[int, List[Tuple[str, int]]]:
        groups: Dict[int, List[Tuple[str, int]]] = {}
        for keyword_str, keyword_type in self.keywords.items():
            length = len(keyword_str)
            if length in groups:
                groups[length].append((keyword_str, keyword_type))
            else:
                groups[length] = [(keyword_str, keyword_type)]
        return groups

    def _setup_keywords(self) -> None:
        groups = self._group_keywords_by_length()
        self.print("var reserved_keywords = [][]KeywordToken{")
        with self.indent():
            num_groups = max(groups) + 1 if groups else 1
            for keywords_length in range(num_groups):
                if keywords_length not in groups.keys():
                    self.print("{},")
                else:
                    self.print("{")
                    with self.indent():
                        for keyword_str, keyword_type in groups[keywords_length]:
                            self.print(f'{{"{keyword_str}", {keyword_type}}},')
                    self.print("},")
        self.print("}")

    def _setup_soft_keywords(self) -> None:
        soft_keywords = sorted(self.soft_keywords)
        self.print("var soft_keywords = []string{")
        with self.indent():
            for keyword in soft_keywords:
                self.print(f'"{keyword}",')
        self.print("}")

    def add_level(self) -> None:
        self.print("p.level++")

    def remove_level(self) -> None:
        self.print("p.level--")

    def add_return(self, ret_val: str) -> None:
        for stmt in self.cleanup_statements:
            self.print(stmt)
        self.remove_level()
        self.print(f"return {ret_val}")

    def unique_varname(self, name: str = "tmpvar") -> str:
        new_var = name + "_" + str(self._varname_counter)
        self._varname_counter += 1
        return new_var

    def call_with_errorcheck_return(self, call_text: str, returnval: str) -> None:
        error_var = self.unique_varname()
        self.print(f"{error_var} := {call_text}")
        self.print(f"if {error_var} {{")
        with self.indent():
            self.add_return(returnval)
        self.print("}")

    def call_with_errorcheck_goto(self, call_text: str, goto_target: str) -> None:
        error_var = self.unique_varname()
        self.print(f"{error_var} := {call_text}")
        self.print(f"if {error_var} {{")
        with self.indent():
            self.print(f"goto {goto_target}")
        self.print("}")

    def out_of_memory_return(
        self,
        expr: str,
        cleanup_code: Optional[str] = None,
    ) -> None:
        self.print(f"if {expr} {{")
        with self.indent():
            if cleanup_code is not None:
                self.print(cleanup_code)
            self.print("p.error_indicator = 1")
            self.print("PyErr_NoMemory()")
            self.add_return("nil")
        self.print("}")

    def out_of_memory_goto(self, expr: str, goto_target: str) -> None:
        self.print(f"if {expr} {{")
        with self.indent():
            self.print("PyErr_NoMemory()")
            self.print(f"goto {goto_target}")
        self.print("}")

    def _set_up_token_start_metadata_extraction(self) -> None:
        self.print("if p.mark == p.fill && _PyPegen_fill_token(p) < 0 {")
        with self.indent():
            self.print("p.error_indicator = 1")
            self.add_return("nil")
        self.print("}")
        self.print("_start_lineno := p.tokens[_mark].lineno")
        self.print("_start_col_offset := p.tokens[_mark].col_offset")

    def _set_up_token_end_metadata_extraction(self) -> None:
        self.print("_token := _PyPegen_get_last_nonnwhitespace_token(p)")
        self.print("if (_token == nil) {")
        with self.indent():
            self.add_return("nil")
        self.print("}")
        self.print("_end_lineno := _token.end_lineno")
        self.print("_end_col_offset := _token.end_col_offset")

    def _check_for_errors(self) -> None:
        self.print("if p.error_indicator > 0 {")
        with self.indent():
            self.add_return("nil")
        self.print("}")

    def _set_up_rule_memoization(self, node: Rule, result_type: str) -> None:
        with self.indent():
            self.add_level()
            self.print(f"var {result_type} _res")
            self.print(f"if _PyPegen_is_memoized(p, {node.name}_type, &_res) {{")
            with self.indent():
                self.add_return("_res")
            self.print("}")
            self.print("_mark := p.mark")
            self.print("_resmark := p.mark")
            self.print("for {")
            with self.indent():
                self.call_with_errorcheck_return(
                    f"_PyPegen_update_memo(p, _mark, {node.name}_type, _res)", "_res"
                )
                self.print("p.mark = _mark")
                self.print(f"_raw := {node.name}_raw(p)")
                self.print("if p.error_indicator > 0 {")
                with self.indent():
                    self.add_return("nil")
                self.print("}")
                self.print("if _raw == nil || p.mark <= _resmark {")
                with self.indent():
                    self.print("break")
                self.print("}")
                self.print("_resmark = p.mark")
                self.print("_res = _raw")
            self.print("}")
            self.print("p.mark = _resmark")
            self.add_return("_res")
        self.print("}")
        self.print(f"func {node.name}_raw(p *Parser) {result_type} {{")

    def _should_memoize(self, node: Rule) -> bool:
        return node.memo and not node.left_recursive

    def _handle_default_rule_body(self, node: Rule, rhs: Rhs, result_type: str) -> None:
        memoize = self._should_memoize(node)

        with self.indent():
            self.add_level()
            self._check_for_errors()
            self.print(f"var _res {result_type}")
            if memoize:
                self.print(f"if _PyPegen_is_memoized(p, {node.name}_type, &_res) {{")
                with self.indent():
                    self.add_return("_res")
                self.print("}")
            self.print("_mark := p.mark")
            if any(alt.action and "EXTRA" in alt.action for alt in rhs.alts):
                self._set_up_token_start_metadata_extraction()
            self.visit(
                rhs,
                is_loop=False,
                is_gather=node.is_gather(),
                rulename=node.name,
            )
            # if self.debug:
            #     self.print(f'D(fprintf(stderr, "Fail at %d: {node.name}\\n", p->mark));')
            self.print("_res = nil")
        self.print("  done:")
        with self.indent():
            if memoize:
                self.print(f"_PyPegen_insert_memo(p, _mark, {node.name}_type, _res)")
            self.add_return("_res")

    def _handle_loop_rule_body(self, node: Rule, rhs: Rhs) -> None:
        memoize = self._should_memoize(node)
        is_repeat1 = node.name.startswith("_loop1")

        with self.indent():
            self.add_level()
            self._check_for_errors()
            self.print("var _res any")
            if memoize:
                self.print(f"if (_PyPegen_is_memoized(p, {node.name}_type, &_res)) {{")
                with self.indent():
                    self.add_return("_res")
                self.print("}")
            self.print("_mark := p.mark")
            self.print("_start_mark := p.mark")
            self.print("_children := []any{}")
            self.print("_n := 0")
            if any(alt.action and "EXTRA" in alt.action for alt in rhs.alts):
                self._set_up_token_start_metadata_extraction()
            self.visit(
                rhs,
                is_loop=True,
                is_gather=node.is_gather(),
                rulename=node.name,
            )
            if is_repeat1:
                self.print("if _n == 0 || p.error_indicator > 0 {")
                with self.indent():
                    self.add_return("nil")
                self.print("}")
            self.print("_seq := _Py_asdl_generic_seq_new(_n, p.arena)")
            self.print("for i := 0; i < _n; i++ { asdl_seq_SET_UNTYPED(_seq, i, _children[i]) }")
            if node.name:
                self.print(f"_PyPegen_insert_memo(p, _start_mark, {node.name}_type, _seq)")
            self.add_return("_seq")

    def visit_Rule(self, node: Rule) -> None:
        is_loop = node.is_loop()
        is_gather = node.is_gather()
        rhs = node.flatten()
        if is_loop or is_gather:
            result_type = "*asdl_seq"
        elif node.type:
            result_type = node.type
        else:
            result_type = "any"

        for line in str(node).splitlines():
            self.print(f"// {line}")

        self.print(f"func {node.name}_rule(p *Parser) {result_type} {{")

        if node.left_recursive and node.leader:
            self._set_up_rule_memoization(node, result_type)

        if node.name.endswith("without_invalid"):
            with self.indent():
                self.print("_prev_call_invalid := p.call_invalid_rules")
                self.print("p.call_invalid_rules = 0")
                self.cleanup_statements.append("p.call_invalid_rules = _prev_call_invalid")

        if is_loop:
            self._handle_loop_rule_body(node, rhs)
        else:
            self._handle_default_rule_body(node, rhs, result_type)

        if node.name.endswith("without_invalid"):
            self.cleanup_statements.pop()

        self.print("}")

    def visit_NamedItem(self, node: NamedItem) -> None:
        call = self.callmakervisitor.generate_call(node)
        if call.assigned_variable:
            call.assigned_variable = self.dedupe(call.assigned_variable)
        self.print(call)

    def visit_Rhs(
        self, node: Rhs, is_loop: bool, is_gather: bool, rulename: Optional[str]
    ) -> None:
        if is_loop:
            assert len(node.alts) == 1
        for alt in node.alts:
            self.visit(alt, is_loop=is_loop, is_gather=is_gather, rulename=rulename)

    def join_conditions(self, keyword: str, node: Any) -> None:
        self.print("condition := func () bool {")
        with self.indent():
            for item in node.items:
                self.visit(item)
            self.print("return true")
        self.print("}")
        self.print(f"{keyword} condition() {{")

    def emit_action(self, node: Alt, cleanup_code: Optional[str] = None) -> None:
        node.action = node.action.replace(
            "EXTRA", "_start_lineno, _start_col_offset, _end_lineno, _end_col_offset"
        )
        self.print(f"_res = {node.action}")

        self.print("if _res == nil && PyErr_Occurred() {")
        with self.indent():
            self.print("p.error_indicator = 1")
            if cleanup_code:
                self.print(cleanup_code)
            self.add_return("nil")
        self.print("}")

    def emit_default_action(self, is_gather: bool, node: Alt) -> None:
        if len(self.local_variable_names) > 1:
            if is_gather:
                assert len(self.local_variable_names) == 2
                self.print(
                    f"_res = _PyPegen_seq_insert_in_front(p, "
                    f"{self.local_variable_names[0]}, {self.local_variable_names[1]})"
                )
            else:
                self.print(
                    f"_res = _PyPegen_dummy_name(p, {', '.join(self.local_variable_names)})"
                )
        else:
            self.print(f"_res = {self.local_variable_names[0]}")

    def emit_dummy_action(self) -> None:
        self.print("_res = _PyPegen_dummy_name(p)")

    def handle_alt_normal(self, node: Alt, is_gather: bool, rulename: Optional[str]) -> None:
        self.join_conditions(keyword="if", node=node)
        # We have parsed successfully all the conditions for the option.
        with self.indent():
            # Prepare to emit the rule action and do so
            if node.action and "EXTRA" in node.action:
                self._set_up_token_end_metadata_extraction()
            if self.skip_actions:
                self.emit_dummy_action()
            elif node.action:
                self.emit_action(node)
            else:
                self.emit_default_action(is_gather, node)

            # As the current option has parsed correctly, do not continue with the rest.
            self.print("goto done")
        self.print("}")

    def handle_alt_loop(self, node: Alt, is_gather: bool, rulename: Optional[str]) -> None:
        # Condition of the main body of the alternative
        self.join_conditions(keyword="for", node=node)
        # We have parsed successfully one item!
        with self.indent():
            # Prepare to emit the rule action and do so
            if node.action and "EXTRA" in node.action:
                self._set_up_token_end_metadata_extraction()
            if self.skip_actions:
                self.emit_dummy_action()
            elif node.action:
                self.emit_action(node, cleanup_code=None)
            else:
                self.emit_default_action(is_gather, node)

            # Add the result of rule to the temporary buffer of children. This buffer
            # will populate later an asdl_seq with all elements to return.
            self.print("_children = append(_children, _res)")
            self.print("_mark = p.mark")
        self.print("}")

    def visit_Alt(
        self, node: Alt, is_loop: bool, is_gather: bool, rulename: Optional[str]
    ) -> None:
        if len(node.items) == 1 and str(node.items[0]).startswith("invalid_"):
            self.print(f"if p.call_invalid_rules {{ // {node}")
        else:
            self.print(f"{{ // {node}")
        with self.indent():
            self._check_for_errors()
            # Prepare variable declarations for the alternative
            vars = self.collect_vars(node)
            for v, var_type in sorted(item for item in vars.items() if item[0] is not None):
                if not var_type:
                    var_type = "any"
                else:
                    var_type += " "
                if v == "_cut_var":
                    # cut_var must be initialized
                    self.print(f"{v} := 0")
                else:
                    self.print(f"var {v} {var_type}")
                if v and v.startswith("_opt_var"):
                    self.print(f"UNUSED({v}) // Silence compiler warnings")

            with self.local_variable_context():
                if is_loop:
                    self.handle_alt_loop(node, is_gather, rulename)
                else:
                    self.handle_alt_normal(node, is_gather, rulename)

            self.print("p.mark = _mark")
            if "_cut_var" in vars:
                self.print("if _cut_var {")
                with self.indent():
                    self.add_return("nil")
                self.print("}")
        self.print("}")

    def collect_vars(self, node: Alt) -> Dict[Optional[str], Optional[str]]:
        types = {}
        with self.local_variable_context():
            for item in node.items:
                name, type = self.add_var(item)
                types[name] = type
        return types

    def add_var(self, node: NamedItem) -> Tuple[Optional[str], Optional[str]]:
        call = self.callmakervisitor.generate_call(node.item)
        name = node.name if node.name else call.assigned_variable
        if name is not None:
            name = self.dedupe(name)
        return_type = call.return_type if node.type is None else node.type
        return name, return_type
