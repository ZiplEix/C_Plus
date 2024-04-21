package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota

	// COMMENTS

	SINGLE_LINE_COMMENT // //
	MULTI_LINE_COMMENT  // /* */

	// IDENTIFIERS
	IDENTIFIER // [a-zA-Z_][a-zA-Z0-9_]*

	// PUNCTUATION
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LBRACKET  // [
	RBRACKET  // ]
	COMMA     // ,
	SEMICOLON // ;
	COLON     // :
	DOT       // .
	QUESTION  // ?
	POUND     // #

	// LITERALS
	INTEGER          // -13 || 42
	UNSIGNED_INTEGER // 42u
	FLOATING         // 3.14
	CHARACTER        // '...'
	STRING           // "..."
	INCLUDER         // #include <...> || #include "..."

	// OPERATORS
	PLUS        // +
	MINUS       // -
	STAR        // *
	SLASH       // /
	PERCENT     // %
	ESPERLUETTE // &
	PIPE        // |
	CARET       // ^
	TILDE       // ~

	// COMPARISON
	EQUAL         // ==
	NOT_EQUAL     // !=
	LESS          // <
	LESS_EQUAL    // <=
	GREATER       // >
	GREATER_EQUAL // >=

	// ASSIGNMENT
	ASSIGN             // =
	PLUS_ASSIGN        // +=
	MINUS_ASSIGN       // -=
	STAR_ASSIGN        // *=
	SLASH_ASSIGN       // /=
	PERCENT_ASSIGN     // %=
	ESPERLUETTE_ASSIGN // &=
	PIPE_ASSIGN        // |=
	CARET_ASSIGN       // ^=
	SHIFT_LEFT_ASSIGN  // <<=
	SHIFT_RIGHT_ASSIGN // >>=
	ARROW              // ->

	// SHIFT
	SHIFT_LEFT  // <<
	SHIFT_RIGHT // >>

	// LOGICAL
	LOGICAL_AND // &&
	LOGICAL_OR  // ||
	LOGICAL_NOT // !

	//
	// KEYWORDS
	//

	// types
	VOID     // void ...
	CHAR     // char ...
	SHORT    // short ...
	INT      // int ...
	LONG     // long ...
	FLOAT    // float ...
	DOUBLE   // double ...
	SIGNED   // signed type ...
	UNSIGNED // unsigned type ...
	POINTER  // *

	SIZEOF // sizeof()
	// INCLUDE // include ...

	// data classes
	TYPEDEF
	STRUCT
	ENUM
	UNION

	// control flow
	IF
	ELSE
	FOR
	WHILE
	DO
	SWITCH
	CASE
	DEFAULT
	BREAK
	CONTINUE
	RETURN
	GOTO

	// storage classes
	EXTERN
	STATIC
	REGISTER
	AUTO
	CONST
	VOLATILE
)

var reservedKeywords = map[string]TokenKind{
	"void":     VOID,
	"char":     CHAR,
	"short":    SHORT,
	"int":      INT,
	"long":     LONG,
	"float":    FLOAT,
	"double":   DOUBLE,
	"signed":   SIGNED,
	"unsigned": UNSIGNED,

	"sizeof": SIZEOF,
	// "include": INCLUDE,

	"typedef": TYPEDEF,
	"struct":  STRUCT,
	"enum":    ENUM,
	"union":   UNION,

	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"while":    WHILE,
	"do":       DO,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"break":    BREAK,
	"continue": CONTINUE,
	"return":   RETURN,
	"goto":     GOTO,

	"extern":   EXTERN,
	"static":   STATIC,
	"register": REGISTER,
	"auto":     AUTO,
	"const":    CONST,
	"volatile": VOLATILE,
}

type Token struct {
	Kind  TokenKind
	Value string
	Line  int
	Col   int
	Index int
}

func (t Token) IsOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if t.Kind == expected {
			return true
		}
	}
	return false
}

func (t Token) Debug(index ...int) {
	if len(index) > 0 {
		fmt.Printf("% 3d: ", index[0])
	}

	if t.IsOneOfMany(INTEGER, UNSIGNED_INTEGER, FLOATING, CHARACTER, STRING, IDENTIFIER, INCLUDER, SINGLE_LINE_COMMENT, MULTI_LINE_COMMENT) {
		fmt.Printf("%s (%s)\n", TokenKindString(t.Kind), t.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(t.Kind))
	}
}

func NewToken(kind TokenKind, value string, line, col int) Token {
	return Token{
		Kind:  kind,
		Value: value,
		Line:  line,
		Col:   col,
	}
}

func TokenKindString(token TokenKind) string {
	switch token {
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case SINGLE_LINE_COMMENT:
		return "SINGLE_LINE_COMMENT"
	case MULTI_LINE_COMMENT:
		return "MULTI_LINE_COMMENT"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LBRACKET:
		return "LBRACKET"
	case RBRACKET:
		return "RBRACKET"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case COLON:
		return "COLON"
	case DOT:
		return "DOT"
	case QUESTION:
		return "QUESTION"
	case POUND:
		return "POUND"
	case INTEGER:
		return "INTEGER"
	case UNSIGNED_INTEGER:
		return "UNSIGNED_INTEGER"
	case FLOATING:
		return "FLOATING"
	case CHARACTER:
		return "CHARACTER"
	case STRING:
		return "STRING"
	case INCLUDER:
		return "INCLUDER"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case PERCENT:
		return "PERCENT"
	case ESPERLUETTE:
		return "ESPERLUETTE"
	case PIPE:
		return "PIPE"
	case CARET:
		return "CARET"
	case TILDE:
		return "TILDE"
	case EQUAL:
		return "EQUAL"
	case NOT_EQUAL:
		return "NOT_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case ASSIGN:
		return "ASSIGN"
	case PLUS_ASSIGN:
		return "PLUS_ASSIGN"
	case MINUS_ASSIGN:
		return "MINUS_ASSIGN"
	case STAR_ASSIGN:
		return "STAR_ASSIGN"
	case SLASH_ASSIGN:
		return "SLASH_ASSIGN"
	case PERCENT_ASSIGN:
		return "PERCENT_ASSIGN"
	case ESPERLUETTE_ASSIGN:
		return "ESPERLUETTE_ASSIGN"
	case PIPE_ASSIGN:
		return "PIPE_ASSIGN"
	case CARET_ASSIGN:
		return "CARET_ASSIGN"
	case SHIFT_LEFT_ASSIGN:
		return "SHIFT_LEFT_ASSIGN"
	case SHIFT_RIGHT_ASSIGN:
		return "SHIFT_RIGHT_ASSIGN"
	case ARROW:
		return "ARROW"
	case SHIFT_LEFT:
		return "SHIFT_LEFT"
	case SHIFT_RIGHT:
		return "SHIFT_RIGHT"
	case LOGICAL_AND:
		return "LOGICAL_AND"
	case LOGICAL_OR:
		return "LOGICAL_OR"
	case LOGICAL_NOT:
		return "LOGICAL_NOT"
	case VOID:
		return "VOID"
	case CHAR:
		return "CHAR"
	case SHORT:
		return "SHORT"
	case INT:
		return "INT"
	case LONG:
		return "LONG"
	case FLOAT:
		return "FLOAT"
	case DOUBLE:
		return "DOUBLE"
	case SIGNED:
		return "SIGNED"
	case UNSIGNED:
		return "UNSIGNED"
	case POINTER:
		return "POINTER"
	case SIZEOF:
		return "SIZEOF"
	// case INCLUDE:
	// 	return "INCLUDE"
	case TYPEDEF:
		return "TYPEDEF"
	case STRUCT:
		return "STRUCT"
	case ENUM:
		return "ENUM"
	case UNION:
		return "UNION"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case FOR:
		return "FOR"
	case WHILE:
		return "WHILE"
	case DO:
		return "DO"
	case SWITCH:
		return "SWITCH"
	case CASE:
		return "CASE"
	case DEFAULT:
		return "DEFAULT"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case RETURN:
		return "RETURN"
	case GOTO:
		return "GOTO"
	case EXTERN:
		return "EXTERN"
	case STATIC:
		return "STATIC"
	case REGISTER:
		return "REGISTER"
	case AUTO:
		return "AUTO"
	case CONST:
		return "CONST"
	case VOLATILE:
		return "VOLATILE"
	default:
		return "UNKNOWN"
	}
}
