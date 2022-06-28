package parser

import (
	"oasis/lexer"
	"testing"
)

func TestExpressions(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"a", "a"},
		{"1", "1"},
		{"-1", "(-1)"},
		{"~2", "(~2)"},
		{"!false", "(!false)"},
		{"(10 + 5)", "(10 + 5)"},
		{"a = 10", "(a = 10)"},
		{"true && true", "(true && true)"},
		{"true || false", "(true || false)"},
		{"1 & 1", "(1 & 1)"},
		{"1 | 0", "(1 | 0)"},
		{"1 ^ 0", "(1 ^ 0)"},
		{"1 + 1", "(1 + 1)"},
		{"1 - 1", "(1 - 1)"},
		{"1 * 1", "(1 * 1)"},
		{"1 / 1", "(1 / 1)"},
		{"1 % 1", "(1 % 1)"},
		{"a()", "a()"},
		{"sum(1, 3)", "sum(1, 3, )"},
		{"{}", "{ }"},
		{"{ 10 }", "{ 10; }"},
		{"if true { 1 }", "if true { 1; }"},
		{"if true { 1 } else { 0 }", "if true { 1; } else { 0; }"},
		{"func() { 10 }", "func() { 10; }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr := p.parseExpr(LOWEST)
		if expr == nil {
			t.Fatalf("tests[%d]: %s", i, p.Error())
		}

		if expr.String() != tt.output {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.output, expr.String())
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"let a = 10", "let a = 10;"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt := p.parseLetStmt()
		if stmt == nil {
			t.Fatalf("tests[%d]: %s", i, p.Error())
		}

		if stmt.String() != tt.output {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.output, stmt.String())
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"return", "return;"},
		{"return 10", "return 10;"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt := p.parseReturnStmt()
		if stmt == nil {
			t.Fatalf("tests[%d]: %s", i, p.Error())
		}

		if stmt.String() != tt.output {
			t.Fatalf("tests[%d]: expected %q, got %q", i, tt.output, stmt.String())
		}
	}
}
