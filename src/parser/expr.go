package parser

import (
	"fmt"
	"strconv"

	"github.com/ZiplEix/c_parser/src/ast"
	"github.com/ZiplEix/c_parser/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	// First parse the nud
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD HANDLER EXPECTED FOR TOKEN '%s' at index %d\n", lexer.TokenKindString(tokenKind), p.currentToken().Index))
	}

	left := nud_fn(p)

	// while we have a led and the current bp is < bp of current token
	// continue parsing the left hand side
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind := p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("LED HANDLER EXPECTED FOR TOKEN '%s'\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.CHARACTER:
		return ast.CharacterExpr{Value: p.advance().Value}
	case lexer.INTEGER:
		number, _ := strconv.ParseInt(p.advance().Value, 10, 64)
		return ast.IntegerExpr{Value: number}
	case lexer.FLOATING:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.FloatExpr{Value: number}
	case lexer.UNSIGNED_INTEGER:
		number, _ := strconv.ParseUint(p.advance().Value, 10, 64)
		return ast.UnsignedIntegerExpr{Value: number}
	case lexer.STRING:
		return ast.StringExpr{Value: p.advance().Value}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{Value: p.advance().Value}
	default:
		panic(fmt.Sprintf("Cannot create primary_expression from %s \n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

// func parse_call_expr(p *parser, left ast.Expr, _ binding_power) ast.Expr {
// 	p.expect(lexer.LPAREN) // consume the LPAREN

// 	args := []ast.Expr{}

// 	for p.currentTokenKind() != lexer.RPAREN {
// 		args = append(args, parse_expr(p, default_bp))

// 		if p.currentTokenKind() == lexer.COMMA {
// 			p.advance()
// 		}
// 	}

// 	p.expect(lexer.RPAREN)

// 	return ast.CallExpr{
// 		Func: left,
// 		Args: args,
// 	}
// }
