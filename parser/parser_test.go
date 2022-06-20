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
			t.Fatalf("tests[%d]: expr: expected %q, got %q", i, tt.expr, es.Expr.String())
		}
	}
}

func TestParseLetStmt(t *testing.T) {
	tests := []struct {
		input string
		name  string
		expr  string
	}{
		{"let a = 10;", "a", "10"},
		{"let b = -10;", "b", "(-10)"},
		{"let c = a + b;", "c", "(a + b)"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		ls, err := p.parseLetStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.name != ls.Name.String() {
			t.Fatalf("tests[%d]: name: expected %q, got %q", i, tt.name, ls.Name.String())
		}

		if tt.expr != ls.Expr.String() {
			t.Fatalf("tests[%d]: expr: expected %q, got %q", i, tt.expr, ls.Expr.String())
		}
	}
}

func TestParseBlockStmt(t *testing.T) {
	tests := []struct {
		input string
		block string
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

		if tt.block != bs.String() {
			t.Fatalf("tests[%d]: expr: expected %q, got %q", i, tt.block, bs.String())
		}
	}
}
