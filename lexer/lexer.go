package lexer

import (
	"oasis/token"
)

type Lexer struct {
	input string

	pos     int
	readPos int

	ch         byte
	insertSemi bool
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if l.insertSemi && (l.ch == 0 || l.ch == '\n' || l.ch == '}') {
		l.insertSemi = false
		return token.Token{Type: token.SEMI, Lit: ";"}
	}

	l.insertSemi = false
	switch l.ch {
	case 0:
		tok = token.Token{Type: token.EOF, Lit: ""}
	case '+':
		tok = token.Token{Type: token.ADD, Lit: "+"}
	case '-':
		tok = token.Token{Type: token.SUB, Lit: "-"}
	case '*':
		tok = token.Token{Type: token.MUL, Lit: "*"}
	case '/':
		tok = token.Token{Type: token.DIV, Lit: "/"}
	case '=':
		if l.peek() == '=' {
			l.advance()
			tok = token.Token{Type: token.EQ, Lit: "=="}
		} else {
			tok = token.Token{Type: token.ASSIGN, Lit: "="}
		}
	case '<':
		tok = token.Token{Type: token.LT, Lit: "<"}
	case '>':
		tok = token.Token{Type: token.GT, Lit: ">"}
	case '!':
		if l.peek() == '=' {
			l.advance()
			tok = token.Token{Type: token.NEQ, Lit: "!="}
		} else {
			tok = token.Token{Type: token.NOT, Lit: "!"}
		}
	case ',':
		tok = token.Token{Type: token.COMMA, Lit: ","}
	case ';':
		tok = token.Token{Type: token.SEMI, Lit: ";"}
	case '(':
		tok = token.Token{Type: token.LPAREN, Lit: "("}
	case ')':
		l.insertSemi = true
		tok = token.Token{Type: token.RPAREN, Lit: ")"}
	case '{':
		tok = token.Token{Type: token.LBRACE, Lit: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Lit: "}"}
	default:
		if isLetter(l.ch) {
			l.insertSemi = true
			tok.Lit = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Lit)
			return tok
		} else if isDigit(l.ch) {
			l.insertSemi = true
			tok.Type = token.INT
			tok.Lit = l.readNumber()
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Lit: string(l.ch)}
		}
	}

	l.advance()

	return tok
}

func (l *Lexer) advance() {
	if l.readPos < len(l.input) {
		l.ch = l.input[l.readPos]
	} else {
		l.ch = 0
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peek() byte {
	if l.readPos < len(l.input) {
		return l.input[l.readPos]
	}
	return 0
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' && !l.insertSemi {
		l.advance()
	}
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) || isDigit(l.ch) {
		l.advance()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.advance()
	}
	return l.input[pos:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
