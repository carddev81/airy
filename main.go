package main

import (
	"fmt"
	"log"
	"os"

	"airy/src/lexer"
)

const helpStr = `Usage: ul [filepath]`

func main() {
	args := os.Args
	if len(args) > 2 {
		log.Fatalf("No input file specified:\n %s", helpStr)
	}
	lexer := lexer.NewLexer(args[1])
	tokens := lexer.Tokenize()
	fmt.Println("final token stream: ")
	for _, token := range tokens {
		token.Print()
	}
}
