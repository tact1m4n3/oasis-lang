package lexer

import (
	"oasis/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "a 10 = + - * / % & | ^ << >> ~ && || ! == != < <= > >= , ; () {} let if else return fn"

	tests := []struct {
		tok token.Token
		lit string
	}{
		{tok: token.IDENT, lit: "a"},
		{tok: token.INT, lit: "10"},
		{tok: token.ASSIGN, lit: "="},
		{tok: token.PLUS, lit: "+"},
		{tok: token.MINUS, lit: "-"},
		{tok: token.ASTERISK, lit: "*"},
		{tok: token.SLASH, lit: "/"},
		{tok: token.MOD, lit: "%"},
		{tok: token.AND, lit: "&"},
		{tok: token.OR, lit: "|"},
		{tok: token.XOR, lit: "^"},
		{tok: token.LSHIFT, lit: "<<"},
		{tok: token.RSHIFT, lit: ">>"},
		{tok: token.NOT, lit: "~"},
		{tok: token.LAND, lit: "&&"},
		{tok: token.LOR, lit: "||"},
		{tok: token.BANG, lit: "!"},
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

		{tok: token.LET, lit: "let"},
		{tok: token.IF, lit: "if"},
		{tok: token.ELSE, lit: "else"},
		{tok: token.RETURN, lit: "return"},
		{tok: token.FUNC, lit: "fn"},
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
