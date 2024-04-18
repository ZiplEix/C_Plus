package main

import (
	"os"

	"github.com/ZiplEix/c_parser/src/lexer"
)

func main() {
	bytes, err := os.ReadFile("main.c")
	if err != nil {
		panic(err)
	}

	tokens := lexer.Tokensize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
