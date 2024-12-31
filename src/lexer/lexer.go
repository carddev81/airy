package lexer

import (
	"airy/src/token"
	"log"
	"os"
	"unicode"
)

type Lexer struct {
	src         string
	pos         int
	readPos     int
	ch          byte
	currentLine int
	tokens      []token.Token
}

func NewLexer(filepath string) Lexer {
	src, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	log.Printf("read file: %s", filepath)
	lex := Lexer{
		src:         string(src),
		tokens:      []token.Token{},
		currentLine: 1,
	}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPos >= len(lex.src) {
		lex.ch = 0
	} else {
		lex.ch = lex.src[lex.readPos]
	}
	lex.pos = lex.readPos
	lex.readPos += 1
}

func (lxr *Lexer) Tokenize() []token.Token {
	for lxr.pos <= len(lxr.src) || lxr.ch != 0 {
		token := lxr.nextToken()
		lxr.tokens = append(lxr.tokens, token)
	}
	return lxr.tokens
}

func (lxr *Lexer) nextToken() token.Token {
	var toke token.Token
	lxr.skipWhitespace()
	switch lxr.ch {
	case '(':
		toke = token.NewToken(token.LParen, lxr.ch, lxr.currentLine)
	case ')':
		toke = token.NewToken(token.RParen, lxr.ch, lxr.currentLine)
	case '[':
		toke = token.NewToken(token.LBracket, lxr.ch, lxr.currentLine)
	case ']':
		toke = token.NewToken(token.RBracket, lxr.ch, lxr.currentLine)
	case '-':
		toke = token.NewToken(token.Minus, lxr.ch, lxr.currentLine)
	case '+':
		toke = token.NewToken(token.Plus, lxr.ch, lxr.currentLine)
	case '*':
		toke = token.NewToken(token.Mul, lxr.ch, lxr.currentLine)
	case '/':
		toke = token.NewToken(token.Div, lxr.ch, lxr.currentLine)
	case '=':
		toke = token.NewToken(token.Eq, lxr.ch, lxr.currentLine)
	case ';':
		toke = token.NewToken(token.Semicolon, lxr.ch, lxr.currentLine)
	case '!':
		toke = token.NewToken(token.Bang, lxr.ch, lxr.currentLine)
	case '%':
		toke = token.NewToken(token.Mod, lxr.ch, lxr.currentLine)
	case '<':
		toke = token.NewToken(token.Lt, lxr.ch, lxr.currentLine)
	case '>':
		toke = token.NewToken(token.Gt, lxr.ch, lxr.currentLine)
	case 0:
		toke.Literal = ""
		toke.Type = token.EOF
		toke.Line = lxr.currentLine
	default:
		if unicode.IsLetter(rune(lxr.ch)) {
			toke.Literal = lxr.readIdentifier() //keeps reading until next ...
			toke.Type = token.LookupIdent(toke.Literal)
			toke.Line = lxr.currentLine
			return toke
		} else if unicode.IsDigit(rune(lxr.ch)) {
			toke.Literal = lxr.readNumber()
			toke.Type = token.NumLiteral
			toke.Line = lxr.currentLine
			return toke
		} else {
			toke = token.NewToken(token.Illegal, lxr.ch, lxr.currentLine)
		}
	}
	lxr.readChar()
	return toke
}

func (lxr *Lexer) skipWhitespace() {
	for lxr.ch == ' ' || lxr.ch == '\t' || lxr.ch == '\n' || lxr.ch == '\r' {
		if lxr.ch == '\n' {
			lxr.currentLine++
		}
		lxr.readChar()
	}
}

func (lxr *Lexer) readIdentifier() string {
	position := lxr.pos
	for unicode.IsLetter(rune(lxr.ch)) {
		lxr.readChar()
	}
	return lxr.src[position:lxr.pos]
}

func (lxr *Lexer) readNumber() string {
	position := lxr.pos
	for unicode.IsDigit(rune(lxr.ch)) {
		lxr.readChar()
	}
	return lxr.src[position:lxr.pos]
}
