package lexer

import (
	"fmt"
	"oasis/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let a0 = 1
let b = 2

fn add(x, y) { return x + y }

let c = add(a0, b)

if c == 3 { print(true) }
else {
	print(false)
}
`

	tests := []token.Token{
		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "a0"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.INT, Lit: "1"},
		{Type: token.SEMI, Lit: ";"},

		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "b"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.INT, Lit: "2"},
		{Type: token.SEMI, Lit: ";"},

		{Type: token.FUNC, Lit: "fn"},
		{Type: token.IDENT, Lit: "add"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "x"},
		{Type: token.COMMA, Lit: ","},
		{Type: token.IDENT, Lit: "y"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.LBRACE, Lit: "{"},
		{Type: token.RETURN, Lit: "return"},
		{Type: token.IDENT, Lit: "x"},
		{Type: token.ADD, Lit: "+"},
		{Type: token.IDENT, Lit: "y"},
		{Type: token.SEMI, Lit: ";"},
		{Type: token.RBRACE, Lit: "}"},

		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "c"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.IDENT, Lit: "add"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "a0"},
		{Type: token.COMMA, Lit: ","},
		{Type: token.IDENT, Lit: "b"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.SEMI, Lit: ";"},

		{Type: token.IF, Lit: "if"},
		{Type: token.IDENT, Lit: "c"},
		{Type: token.EQ, Lit: "=="},
		{Type: token.INT, Lit: "3"},
		{Type: token.LBRACE, Lit: "{"},

		{Type: token.IDENT, Lit: "print"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "true"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.SEMI, Lit: ";"},

		{Type: token.RBRACE, Lit: "}"},
		{Type: token.ELSE, Lit: "else"},
		{Type: token.LBRACE, Lit: "{"},

		{Type: token.IDENT, Lit: "print"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "false"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.SEMI, Lit: ";"},

		{Type: token.RBRACE, Lit: "}"},

		{Type: token.EOF, Lit: ""},
	}

	l := New(input)
	for _, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok.Type)
		if tok.Type != tt.Type {
			t.Fatalf("wrong token type(%d): expected %q, got %q", tok.Pos, tt.Type, tok.Type)
		}

		if tok.Lit != tt.Lit {
			t.Fatalf("wrong literal: expected %q, got %q", tt.Lit, tok.Lit)
		}
	}
}
