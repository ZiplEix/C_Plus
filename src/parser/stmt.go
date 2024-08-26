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

func parse_block_stmt(p *parser) ast.Stmt {
	p.expect(lexer.LBRACE)
	body := []ast.Stmt{}

	for p.hasTokens() && p.currentTokenKind() != lexer.RBRACE {
		body = append(body, parseStmt(p))
	}

	p.expect(lexer.RBRACE)

	return ast.BlockStmt{
		Body: body,
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
	isConst, isSigned, pointerLevel, varType, varName := parse_var_type_and_name(p)

	// var declaration without assigment
	if p.currentTokenKind() == lexer.SEMICOLON {
		p.advance()
		return &ast.VarDeclarationStmt{
			Name:         varName,
			IsConst:      isConst,
			IsSigned:     isSigned,
			PointerLevel: pointerLevel,
			Type:         varType,
		}
	} else if p.currentToken().Kind == lexer.LPAREN {
		// function declaration
		return parse_func_declaration_stmt(p, varName, isConst, isSigned, pointerLevel, varType)
	}

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

func parse_var_type_and_name(p *parser) (bool, bool, int, ast.VarType, string) {
	isConst := false
	isSigned := true
	pointerLevel := 0
	varType := ast.INT

	for isType(p.currentTokenKind()) || p.currentTokenKind() == lexer.CONST || p.currentTokenKind() == lexer.SIGNED || p.currentTokenKind() == lexer.UNSIGNED || p.currentTokenKind() == lexer.STAR {
		if p.currentTokenKind() == lexer.CONST {
			isConst = true
		}
		if p.currentTokenKind() == lexer.SIGNED {
			isSigned = true
		}
		if p.currentTokenKind() == lexer.UNSIGNED {
			isSigned = false
		}
		if p.currentTokenKind() == lexer.STAR {
			pointerLevel++
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

	varName := p.expect(lexer.IDENTIFIER).Value

	return isConst, isSigned, pointerLevel, varType, varName
}

func parse_function_param_and_body(p *parser) ([]ast.Parameter, []ast.Stmt) {
	functionParameters := make([]ast.Parameter, 0)

	p.expect(lexer.LPAREN)
	for p.hasTokens() && p.currentTokenKind() != lexer.RPAREN {
		isConst, isSigned, pointerLevel, varType, varName := parse_var_type_and_name(p)

		currentParameter := ast.Parameter{
			Name:         varName,
			IsConst:      isConst,
			IsSigned:     isSigned,
			PointerLevel: pointerLevel,
			Type:         varType,
		}

		// fmt.Printf("Current parameter:\n")
		// litter.Dump(currentParameter)

		functionParameters = append(functionParameters, currentParameter)

		if !p.currentToken().IsOneOfMany(lexer.COMMA, lexer.RPAREN) {
			p.expect(lexer.COMMA)
		}

		if p.currentTokenKind() == lexer.RPAREN {
			break
		}

		p.advance()
	}

	p.expect(lexer.RPAREN)

	functionBody := ast.ExpectStmt[ast.BlockStmt](parse_block_stmt(p)).Body

	return functionParameters, functionBody
}

func parse_func_declaration_stmt(p *parser, functionName string, isConst, isSigned bool, pointerLevel int, returnType ast.VarType) ast.Stmt {
	functionParameters, functionBody := parse_function_param_and_body(p)

	return &ast.FunctionDeclarationStmt{
		Name:         functionName,
		IsConst:      isConst,
		IsSigned:     isSigned,
		PointerLevel: pointerLevel,
		ReturnType:   returnType,
		Parameters:   functionParameters,
		Body:         functionBody,
	}
}

func parse_return_stmt(p *parser) ast.Stmt {
	p.advance()

	expr := parse_expr(p, default_bp)
	p.expect(lexer.SEMICOLON)

	return &ast.ReturnStmt{
		Expr: expr,
	}
}

func parse_comment_stmt(p *parser) ast.Stmt {
	var kind ast.VarType

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
	p.expect(lexer.INCLUDER) // skip the includer token

	return &ast.IncluderStmt{
		Value: p.advance().Value,
	}
}
