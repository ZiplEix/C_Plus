package main

import (
	"fmt"
	"os"

	"github.com/ZiplEix/c_parser/src/lexer"
	"github.com/ZiplEix/c_parser/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, err := os.ReadFile("main.c")
	if err != nil {
		panic(err)
	}

	tokens := lexer.Tokensize(string(bytes))

	fmt.Printf("------\n")
	fmt.Printf("TOKENS\n")
	fmt.Printf("------\n")

	for index, token := range tokens {
		token.Debug(index)
	}

	fmt.Printf("\n------\n")
	fmt.Printf("AST\n")
	fmt.Printf("------\n")

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
