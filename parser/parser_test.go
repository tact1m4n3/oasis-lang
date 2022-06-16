package parser

import (
	"oasis/lexer"
	"testing"
)

func TestParseLetStmt(t *testing.T) {
	tests := []struct {
		input string
		name  string
		expr  string
	}{
		{"let a = 10;", "a", "10"},
		{"let b = -10;", "b", "(-10)"},
		{"let c = 20 + 20 / 2;", "c", "(20 + (20 / 2))"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt, err := p.ParseLetStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.name != stmt.Name {
			t.Fatalf("tests[%d]: name: expected %q, got %q", i, tt.name, stmt.Name)
		}

		if tt.expr != stmt.Expr.String() {
			t.Fatalf("tests[%d]: expr: expected %q, got %q", i, tt.expr, stmt.Expr.String())
		}
	}
}

func TestParseExprStmt(t *testing.T) {
	tests := []struct {
		input string
		expr  string
	}{
		{"10;", "10"},
		{"-10;", "(-10)"},
		{"20 + 20 / 2;", "(20 + (20 / 2))"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt, err := p.ParseExprStmt()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.expr != stmt.Expr.String() {
			t.Fatalf("tests[%d]: expr: expected %q, got %q", i, tt.expr, stmt.Expr.String())
		}
	}
}

func TestParseExpr(t *testing.T) {
	tests := []struct {
		input string
		expr  string
	}{
		{"a + b", "(a + b)"},
		{"-10 + 10", "((-10) + 10)"},
		{"15 / 5 + 10 / 5", "((15 / 5) + (10 / 5))"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr, err := p.ParseExpr()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.expr != expr.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.expr, expr.String())
		}
	}
}

func TestParseTerm(t *testing.T) {
	tests := []struct {
		input string
		expr  string
	}{
		{"a * b", "(a * b)"},
		{"-10 * 10", "((-10) * 10)"},
		{"15 * 10 / 5", "(15 * (10 / 5))"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr, err := p.ParseTerm()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.expr != expr.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.expr, expr.String())
		}
	}
}

func TestParseFactor(t *testing.T) {
	tests := []struct {
		input string
		expr  string
	}{
		{"a", "a"},
		{"10", "10"},
		{"-10", "(-10)"},
		{"!true", "(!true)"},
		{"-(10 + 10)", "(-(10 + 10))"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr, err := p.ParseFactor()
		if err != nil {
			t.Fatalf("tests[%d]: parser error: %s", i, err)
		}

		if tt.expr != expr.String() {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.expr, expr.String())
		}
	}
}
