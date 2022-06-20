package parser

import (
	"oasis/lexer"
	"testing"
)

func TestParseExprStmt(t *testing.T) {
	tests := []struct {
		input string
		expr  string
	}{
		{"10", "10"},
		{"-10", "(-10)"},
		{"20 + 20 / 2", "(20 + (20 / 2))"},
		{"(20 + 20) / 2", "((20 + 20) / 2)"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		es, err := p.parseExprStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.expr != es.Expr.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.expr, es.Expr.String())
		}
	}
}

func TestParseLetStmt(t *testing.T) {
	tests := []struct {
		input string
		stmt  string
	}{
		{"let a = 10;", "let a = 10;"},
		{"let b = -10;", "let b = (-10);"},
		{"let c = a + b;", "let c = (a + b);"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		ls, err := p.parseLetStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.stmt != ls.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.stmt, ls.String())
		}
	}
}

func TestParseBlockStmt(t *testing.T) {
	tests := []struct {
		input string
		stmt  string
	}{
		{"{ let a = 10; a + 10 }", "{ let a = 10; (a + 10); }"},
		{`{
	let x = 10
	let y = 10
	let z = x + y
	z
}`, "{ let x = 10; let y = 10; let z = (x + y); z; }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		bs, err := p.parseBlockStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.stmt != bs.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.stmt, bs.String())
		}
	}
}

func TestParseIfStmt(t *testing.T) {
	tests := []struct {
		input string
		stmt  string
	}{
		{"if a > b { a }", "if (a > b) { a; }"},
		{"if a > b { a } else { b }", "if (a > b) { a; } else { b; }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		is, err := p.parseIfStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.stmt != is.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.stmt, is.String())
		}
	}
}

func TestParseFuncStmt(t *testing.T) {
	tests := []struct {
		input string
		stmt  string
	}{
		{"fn nothing() { 10 }", "fn nothing() { 10; }"},
		{"fn sum(a, b) { return a + b }", "fn sum(a, b, ) { return (a + b); }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		is, err := p.parseFuncStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.stmt != is.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.stmt, is.String())
		}
	}
}
