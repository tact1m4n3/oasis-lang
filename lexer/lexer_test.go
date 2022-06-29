package lexer

import (
	"oasis/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `a 10
= + - * / % & | ^ << >> ~ += -= *= /= %= &= |= ^= <<= >>=
&& || ! == != < <= > >=
, ;
() {}
let if else return func`

	tests := []struct {
		tok token.Token
		lit string
	}{
		{tok: token.IDENT, lit: "a"},
		{tok: token.INT, lit: "10"},
		{tok: token.SEMI, lit: ";"},

		{tok: token.ASSIGN, lit: "="},
		{tok: token.ADD, lit: "+"},
		{tok: token.SUB, lit: "-"},
		{tok: token.MUL, lit: "*"},
		{tok: token.DIV, lit: "/"},
		{tok: token.MOD, lit: "%"},
		{tok: token.AND, lit: "&"},
		{tok: token.OR, lit: "|"},
		{tok: token.XOR, lit: "^"},
		{tok: token.LSHIFT, lit: "<<"},
		{tok: token.RSHIFT, lit: ">>"},
		{tok: token.TILDE, lit: "~"},
		{tok: token.ADD_ASSIGN, lit: "+="},
		{tok: token.SUB_ASSIGN, lit: "-="},
		{tok: token.MUL_ASSIGN, lit: "*="},
		{tok: token.DIV_ASSIGN, lit: "/="},
		{tok: token.MOD_ASSIGN, lit: "%="},
		{tok: token.AND_ASSIGN, lit: "&="},
		{tok: token.OR_ASSIGN, lit: "|="},
		{tok: token.XOR_ASSIGN, lit: "^="},
		{tok: token.LSHIFT_ASSIGN, lit: "<<="},
		{tok: token.RSHIFT_ASSIGN, lit: ">>="},

		{tok: token.LAND, lit: "&&"},
		{tok: token.LOR, lit: "||"},
		{tok: token.NOT, lit: "!"},
		{tok: token.EQ, lit: "=="},
		{tok: token.NEQ, lit: "!="},
		{tok: token.LT, lit: "<"},
		{tok: token.LTE, lit: "<="},
		{tok: token.GT, lit: ">"},
		{tok: token.GTE, lit: ">="},

		{tok: token.COMMA, lit: ","},
		{tok: token.SEMI, lit: ";"},

		{tok: token.LPAREN, lit: "("},
		{tok: token.RPAREN, lit: ")"},
		{tok: token.LBRACE, lit: "{"},
		{tok: token.RBRACE, lit: "}"},
		{tok: token.SEMI, lit: ";"},

		{tok: token.LET, lit: "let"},
		{tok: token.IF, lit: "if"},
		{tok: token.ELSE, lit: "else"},
		{tok: token.RETURN, lit: "return"},
		{tok: token.FUNC, lit: "func"},
		{tok: token.SEMI, lit: ";"},
	}

	l := New(input)
	for _, tt := range tests {
		tok, lit := l.NextToken()

		if tok != tt.tok {
			t.Fatalf("wrong token type: expected %q, got %q", tt.tok, tok)
		}

		if lit != tt.lit {
			t.Fatalf("wrong literal: expected %q, got %q", tt.lit, lit)
		}
	}
}
