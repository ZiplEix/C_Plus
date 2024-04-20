package ast

import "github.com/ZiplEix/c_parser/src/lexer"

//
// LITERAL EXPRESSION
//

type IntegerExpr struct {
	Value int64
}

func (e IntegerExpr) expr() {}

type FloatExpr struct {
	Value float64
}

func (e FloatExpr) expr() {}

type UnsignedIntegerExpr struct {
	Value uint64
}

func (e UnsignedIntegerExpr) expr() {}

type CharacterExpr struct {
	Value string
}

func (e CharacterExpr) expr() {}

type StringExpr struct {
	Value string
}

func (e StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (e SymbolExpr) expr() {}

//
// BINARY EXPRESSION
//

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (e BinaryExpr) expr() {}
