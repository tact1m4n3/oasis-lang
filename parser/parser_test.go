package parser

import (
	"oasis/lexer"
	"testing"
)

func TestParseExprStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"10", "10;"},
		{"-10", "(-10);"},
		{"20 + 20 / 2", "(20 + (20 / 2));"},
		{"(20 + 20) / 2", "((20 + 20) / 2);"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func TestParseLetStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"let a = 10;", "let a = 10;"},
		{"let b = -10;", "let b = (-10);"},
		{"let c = a + b;", "let c = (a + b);"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func TestParseBlockStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
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

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func TestParseIfStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"if a > b { a }", "if (a > b) { a; }"},
		{"if a > b { a } else { b }", "if (a > b) { a; } else { b; }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func TestParseReturnStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"fn nothing() { 10 }", "fn nothing() { 10; }"},
		{"fn sum(a, b) { return a + b }", "fn sum(a, b, ) { return (a + b); }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func TestParseFuncStmt(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"fn nothing() { 10 }", "fn nothing() { 10; }"},
		{"fn sum(a, b) { return a + b }", "fn sum(a, b, ) { return (a + b); }"},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		if !testStmtParsing(t, i, p, tt.output) {
			t.Fail()
		}
	}
}

func testStmtParsing(t *testing.T, i int, p *Parser, output string) bool {
	stmt, err := p.parseStmt()
	if err != nil {
		t.Errorf("tests[%d]: parser error: %s", i, err)
		return false
	}

	if stmt.String() != output {
		t.Errorf("tests[%d]: expected %q, got %q", i, output, stmt.String())
		return false
	}

	return true
}
