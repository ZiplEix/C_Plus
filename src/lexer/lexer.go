package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func Tokensize(source string) []Token {
	lex := createLexer(source)

	// iterate while there are still characters to lex
	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token at pos %d, near '%s'\n", lex.pos, lex.remainder()))
		}
	}

	lex.push(NewToken(EOF, "", 0, 0))
	return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value, 0, 0))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			// Whitespace & comments("//", "/* */")
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`(?s)\/\*.*?\*\/`), skipHandler},

			// LITERALS
			{regexp.MustCompile(`[0-9]+`), integerHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), floatHandler},
			{regexp.MustCompile(`"((?:[^"\\]|\\.)*?)"`), stringHandler},
			{regexp.MustCompile(`'(\\[^\n']|[^\\'\n]|\\.)'`), charHandler},
			{regexp.MustCompile(`#include\s*<.*>`), includerHandler},
			{regexp.MustCompile(`#include\s*".*"`), includerHandler},

			// SYMBOLS
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},

			// PUNCTUATION
			{regexp.MustCompile(`\(`), defaultHandler(LPAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(RPAREN, ")")},
			{regexp.MustCompile(`\{`), defaultHandler(LBRACE, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(RBRACE, "}")},
			{regexp.MustCompile(`\[`), defaultHandler(LBRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(RBRACKET, "]")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`#`), defaultHandler(POUND, "#")},

			// OPERATORS
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`==`), defaultHandler(PERCENT, "%")},
			{regexp.MustCompile(`!=`), defaultHandler(ESPERLUETTE, "&")},
			{regexp.MustCompile(`!=`), defaultHandler(PIPE, "|")},
			{regexp.MustCompile(`!=`), defaultHandler(CARET, "^")},
			{regexp.MustCompile(`!=`), defaultHandler(TILDE, "~")},

			// COMPARISON
			{regexp.MustCompile(`==`), defaultHandler(EQUAL, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUAL, "!=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUAL, "<=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUAL, ">=")},

			// ASSIGNMENT
			{regexp.MustCompile(`=`), defaultHandler(ASSIGN, "=")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_ASSIGN, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_ASSIGN, "-=")},
			{regexp.MustCompile(`\*=`), defaultHandler(STAR_ASSIGN, "*=")},
			{regexp.MustCompile(`/=`), defaultHandler(SLASH_ASSIGN, "/=")},
			{regexp.MustCompile(`%=`), defaultHandler(PERCENT_ASSIGN, "%=")},
			{regexp.MustCompile(`&=`), defaultHandler(ESPERLUETTE_ASSIGN, "&=")},
			{regexp.MustCompile(`\|=`), defaultHandler(PIPE_ASSIGN, "|=")},
			{regexp.MustCompile(`\^=`), defaultHandler(CARET_ASSIGN, "^=")},
			{regexp.MustCompile(`<<=`), defaultHandler(SHIFT_LEFT_ASSIGN, "<<=")},
			{regexp.MustCompile(`>>=`), defaultHandler(SHIFT_RIGHT_ASSIGN, ">>=")},
			{regexp.MustCompile(`->`), defaultHandler(ARROW, "->")},

			// SHIFT
			{regexp.MustCompile(`<<`), defaultHandler(SHIFT_LEFT, "<<")},
			{regexp.MustCompile(`>>`), defaultHandler(SHIFT_RIGHT, ">>")},

			// LOGICAL
			{regexp.MustCompile(`&&`), defaultHandler(LOGICAL_AND, "&&")},
			{regexp.MustCompile(`\|\|`), defaultHandler(LOGICAL_OR, "||")},
			{regexp.MustCompile(`!`), defaultHandler(LOGICAL_NOT, "!")},
		},
	}
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func integerHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(INT, match, 0, 0))
	lex.advanceN(len(match))
}

func floatHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(FLOAT, match, 0, 0))
	lex.advanceN(len(match))
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral, 0, 0))
	lex.advanceN(len(stringLiteral) + 2)
}

func charHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	charLiteral := lex.remainder()[match[0]:match[1]]

	lex.push(NewToken(CHAR, charLiteral, 0, 0))
	lex.advanceN(len(charLiteral) + 2)
}
func includerHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	includeLiteral := lex.remainder()[match[0]:match[1]]

	lex.push(NewToken(INCLUDER, includeLiteral, 0, 0))
	lex.advanceN(len(includeLiteral) + 2)
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remainder())

	if kind, exists := reservedKeywords[value]; exists {
		lex.push(NewToken(kind, value, 0, 0))
	} else {
		lex.push(NewToken(IDENTIFIER, value, 0, 0))
	}

	lex.advanceN(len(value))
}
