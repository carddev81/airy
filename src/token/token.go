package token

import (
	"fmt"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

type TokenType int

const (
	//symbols
	LParen TokenType = iota
	RParen
	LBracket
	RBracket
	Eq
	Minus
	Plus
	Semicolon
	Mul
	Div
	Mod
	Bang
	Gt
	Lt

	//keywords
	False
	Let
	True
	Else
	Then
	If

	//others
	EOF
	Identifier
	NumLiteral
	Illegal
	//Int
	//Underscore
)

func NewToken(tokenType TokenType, val byte, line int) Token {
	return Token{
		Type:    tokenType,
		Literal: string(val),
		Line:    line,
	}
}

var keywords = map[string]TokenType{
	"let":   Let,
	"false": False,
	"true":  True,
	"then":  Then,
	"if":    If,
	"else":  Else,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}

func (tk *Token) Print() {
	fmt.Printf("TOKEN: %v Type: %v line: %d\n", tk.Literal, tk.Type.getTypeString(), tk.Line)
}

func (tt TokenType) getTypeString() string {
	switch {
	case tt >= LParen && tt <= Lt:
		return "SYMBOL"
	case tt >= False && tt <= If:
		return "KEYWORD"
	case tt == NumLiteral:
		return "NUMBER_LITERAL"
	case tt == Identifier:
		return "IDENTIFIER"
	case tt == Illegal:
		return "ILLEGAL"
	case tt == EOF:
		return "END_OF_FILE"
	default:
		return ""
	}
}
