package parser

import (
	"github.com/ZiplEix/c_parser/src/ast"
	"github.com/ZiplEix/c_parser/src/lexer"
)

func parseStmt(p *parser) ast.Stmt {
	stmt_fn, exist := stmt_lu[p.currentTokenKind()]

	if exist {
		return stmt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)

	return &ast.ExprStmt{
		Expr: expression,
	}
}

func isType(tokenKind lexer.TokenKind) bool {
	switch tokenKind {
	case lexer.INT, lexer.CHAR, lexer.FLOAT, lexer.DOUBLE, lexer.VOID, lexer.SHORT, lexer.LONG:
		return true
	}

	return false
}

func parse_var_declaration_stmt(p *parser) ast.Stmt {
	isConst := false
	isSigned := true
	pointerLevel := 0
	varType := ast.INT

	for isType(p.currentTokenKind()) || p.currentTokenKind() == lexer.CONST || p.currentTokenKind() == lexer.SIGNED || p.currentTokenKind() == lexer.UNSIGNED {
		if p.currentTokenKind() == lexer.CONST {
			isConst = true
		}
		if p.currentTokenKind() == lexer.SIGNED {
			isSigned = true
		}
		if p.currentTokenKind() == lexer.UNSIGNED {
			isSigned = false
		}

		if isType(p.currentTokenKind()) {
			switch p.currentTokenKind() {
			case lexer.VOID:
				varType = ast.VOID
			case lexer.INT:
				varType = ast.INT
			case lexer.CHAR:
				varType = ast.CHAR
			case lexer.FLOAT:
				varType = ast.FLOAT
			case lexer.DOUBLE:
				varType = ast.DOUBLE
			}
		}

		p.advance()
	}

	for p.currentTokenKind() == lexer.STAR {
		pointerLevel++
		p.advance()
	}

	varName := p.expectError(lexer.IDENTIFIER, "Expected an identifier").Value

	p.expect(lexer.ASSIGN)
	assignedExpr := parse_expr(p, assignment)
	p.expect(lexer.SEMICOLON)

	return &ast.VarDeclarationStmt{
		Name:         varName,
		IsConst:      isConst,
		IsSigned:     isSigned,
		PointerLevel: pointerLevel,
		Type:         varType,
		AssignedExpr: assignedExpr,
	}
}

func parse_return_stmt(p *parser) ast.Stmt {
	expr := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)

	return &ast.ReturnStmt{
		Expr: expr,
	}
}

func parse_comment_stmt(p *parser) ast.Stmt {
	var kind ast.Type

	if p.currentTokenKind() == lexer.MULTI_LINE_COMMENT {
		kind = ast.MULTI_LINE_COMMENT
	} else {
		kind = ast.SINGLE_LINE_COMMENT
	}

	return &ast.CommentStmt{
		Value: p.advance().Value,
		Type:  kind,
	}
}

func parse_includer_stmt(p *parser) ast.Stmt {
	return &ast.IncluderStmt{
		Value: p.advance().Value,
	}
}
