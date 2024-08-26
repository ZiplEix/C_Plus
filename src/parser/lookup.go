package parser

import (
	"github.com/ZiplEix/c_parser/src/ast"
	"github.com/ZiplEix/c_parser/src/lexer"
)

type binding_power int

const (
	default_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var (
	bp_lu   = bp_lookup{}
	nud_lu  = nud_lookup{}
	led_lu  = led_lookup{}
	stmt_lu = stmt_lookup{}
)

// led (Left Denotation) est utilisé pour les jetons qui apparaissent au
// milieu ou à la fin d'une expression (comme les opérateurs binaires), et
// cette fonction est responsable de l'analyse de ces opérations.
//
// Par exemple, lorsqu'un opérateur binaire comme "+", "-", "*", etc., est
// rencontré, ces cas sont gérés par des fonctions spécifiques définies dans
// led_lu.
func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

// nud (Null Denotation) est utilisé pour les jetons qui peuvent apparaître au
// début d'une expression (comme les identificateurs ou les littéraux), et cette
// fonction est responsable de l'analyse de ces expressions.
//
// Par exemple, une expression peut commencer par un nombre, une chaîne ou un
// identifiant, et ces cas sont gérés par des fonctions spécifiques définies
// dans nud_lu.
func nud(kind lexer.TokenKind, _ binding_power, nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookup() {
	// Logical
	led(lexer.LOGICAL_AND, logical, parse_binary_expr)
	led(lexer.LOGICAL_OR, logical, parse_binary_expr)

	// Relational
	led(lexer.EQUAL, relational, parse_binary_expr)
	led(lexer.NOT_EQUAL, relational, parse_binary_expr)
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUAL, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUAL, relational, parse_binary_expr)

	// Additive & Multiplicative
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.MINUS, additive, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// literals & symbols
	nud(lexer.INTEGER, primary, parse_primary_expr)
	nud(lexer.FLOATING, primary, parse_primary_expr)
	nud(lexer.CHARACTER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)

	// Computed / Call
	// led(lexer.LPAREN, call, parse_call_expr)

	// Statements
	stmt(lexer.MULTI_LINE_COMMENT, parse_comment_stmt)
	stmt(lexer.SINGLE_LINE_COMMENT, parse_comment_stmt)

	stmt(lexer.INCLUDER, parse_includer_stmt)

	stmt(lexer.RETURN, parse_return_stmt)

	stmt(lexer.VOID, parse_var_declaration_stmt)
	stmt(lexer.CHAR, parse_var_declaration_stmt)
	stmt(lexer.SHORT, parse_var_declaration_stmt)
	stmt(lexer.INT, parse_var_declaration_stmt)
	stmt(lexer.LONG, parse_var_declaration_stmt)
	stmt(lexer.FLOAT, parse_var_declaration_stmt)
	stmt(lexer.DOUBLE, parse_var_declaration_stmt)
	stmt(lexer.SIGNED, parse_var_declaration_stmt)
	stmt(lexer.UNSIGNED, parse_var_declaration_stmt)
	stmt(lexer.CONST, parse_var_declaration_stmt)
}
