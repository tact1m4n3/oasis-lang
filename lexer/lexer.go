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

func (l *Lexer) NextToken() (token.Token, string) {
	l.skipWhitespace()

	if l.insertSemi && (l.ch == 0 || l.ch == '\n' || l.ch == '}') {
		l.insertSemi = false
		return token.SEMI, ";"
	}

	var tok token.Token
	var lit string
	l.insertSemi = false
	switch l.ch {
	case 0:
		tok = token.EOF
		lit = ""
	case '=':
		if l.peek() == '=' {
			l.advance()
			tok = token.EQ
			lit = "=="
		} else {
			tok = token.ASSIGN
			lit = "="
		}
	case '+':
		if l.peek() == '=' {
			l.advance()
			tok = token.ADD_ASSIGN
			lit = "+="
		} else {
			tok = token.ADD
			lit = "+"
		}
	case '-':
		if l.peek() == '=' {
			l.advance()
			tok = token.SUB_ASSIGN
			lit = "-="
		} else {
			tok = token.SUB
			lit = "-"
		}
	case '*':
		if l.peek() == '=' {
			l.advance()
			tok = token.MUL_ASSIGN
			lit = "*="
		} else {
			tok = token.MUL
			lit = "*"
		}
	case '/':
		if l.peek() == '=' {
			l.advance()
			tok = token.DIV_ASSIGN
			lit = "/="
		} else {
			tok = token.DIV
			lit = "/"
		}
	case '%':
		if l.peek() == '=' {
			l.advance()
			tok = token.MOD_ASSIGN
			lit = "%="
		} else {
			tok = token.MOD
			lit = "%"
		}
	case '&':
		if l.peek() == '=' {
			l.advance()
			tok = token.AND_ASSIGN
			lit = "&="
		} else if l.peek() == '&' {
			l.advance()
			tok = token.LAND
			lit = "&&"
		} else {
			tok = token.AND
			lit = "&"
		}
	case '|':
		if l.peek() == '=' {
			l.advance()
			tok = token.OR_ASSIGN
			lit = "|="
		} else if l.peek() == '|' {
			l.advance()
			tok = token.LOR
			lit = "||"
		} else {
			tok = token.OR
			lit = "|"
		}
	case '^':
		if l.peek() == '=' {
			l.advance()
			tok = token.XOR_ASSIGN
			lit = "^="
		} else {
			tok = token.XOR
			lit = "^"
		}
	case '<':
		if l.peek() == '<' {
			l.advance()
			if l.peek() == '=' {
				l.advance()
				tok = token.LSHIFT_ASSIGN
				lit = "<<="
			} else {
				tok = token.LSHIFT
				lit = "<<"
			}
		} else if l.peek() == '=' {
			l.advance()
			tok = token.LTE
			lit = "<="
		} else {
			tok = token.LT
			lit = "<"
		}
	case '>':
		if l.peek() == '>' {
			l.advance()
			if l.peek() == '=' {
				l.advance()
				tok = token.RSHIFT_ASSIGN
				lit = ">>="
			} else {
				tok = token.RSHIFT
				lit = ">>"
			}
		} else if l.peek() == '=' {
			l.advance()
			tok = token.GTE
			lit = ">="
		} else {
			tok = token.GT
			lit = ">"
		}
	case '~':
		tok = token.TILDE
		lit = "~"
	case '!':
		if l.peek() == '=' {
			l.advance()
			tok = token.NEQ
			lit = "!="
		} else {
			tok = token.NOT
			lit = "!"
		}
	case ',':
		tok = token.COMMA
		lit = ","
	case ';':
		tok = token.SEMI
		lit = ";"
	case '(':
		tok = token.LPAREN
		lit = "("
	case ')':
		l.insertSemi = true
		tok = token.RPAREN
		lit = ")"
	case '{':
		tok = token.LBRACE
		lit = "{"
	case '}':
		l.insertSemi = true
		tok = token.RBRACE
		lit = "}"
	default:
		if isLetter(l.ch) {
			l.insertSemi = true
			lit = l.readIdent()
			tok = token.LookupIdent(lit)
			return tok, lit
		} else if isDigit(l.ch) {
			l.insertSemi = true
			tok = token.INT
			lit = l.readNumber()
			return tok, lit
		} else {
			return token.ILLEGAL, string(l.ch)
		}
	}

	l.advance()

	return tok, lit
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

func (l *Lexer) readIdent() string {
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
