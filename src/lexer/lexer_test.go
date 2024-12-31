package lexer

import (
	"airy/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
	}{
		{token.Let, "let", 1},
		{token.Identifier, "x", 1},
		{token.Eq, "=", 1},
		{token.LParen, "(", 1},
		{token.LParen, "(", 1},
		{token.NumLiteral, "44", 1},
		{token.Plus, "+", 1},
		{token.NumLiteral, "23", 1},
		{token.RParen, ")", 1},
		{token.Mul, "*", 1},
		{token.LParen, "(", 1},
		{token.NumLiteral, "23", 1},
		{token.Minus, "-", 1},
		{token.NumLiteral, "12", 1},
		{token.RParen, ")", 1},
		{token.RParen, ")", 1},
		{token.Div, "/", 1},
		{token.NumLiteral, "4", 1},
		{token.Semicolon, ";", 1},
		//new line
		{token.Let, "let", 3},
		{token.Identifier, "y", 3},
		{token.Eq, "=", 3},
		{token.If, "if", 3},
		{token.LParen, "(", 3},
		{token.Identifier, "x", 3},
		{token.Gt, ">", 3},
		{token.NumLiteral, "100", 3},
		{token.RParen, ")", 3},
		{token.Then, "then", 3},
		{token.NumLiteral, "1", 3},
		{token.Else, "else", 3},
		{token.LParen, "(", 3},
		{token.NumLiteral, "0", 3},
		{token.RParen, ")", 3},
		{token.Semicolon, ";", 3},
		{token.EOF, "", 3},
	}
	l := NewLexer("testdata/bb.ul")
	for i, tt := range tests {
		tok := l.nextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%d, got=%d",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line nuymber wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}
	}
}
