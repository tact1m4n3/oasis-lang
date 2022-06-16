package parser

import (
	"fmt"
	"oasis/ast"
	"oasis/lexer"
	"oasis/token"
)

type Parser struct {
	l      *lexer.Lexer
	curTok token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.advance()
	return p
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	var stmts []ast.Stmt
	for p.curTok.Type != token.EOF {
		stmt, err := p.ParseStatement()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return &ast.Program{Stmts: stmts}, nil
}

func (p *Parser) ParseStatement() (ast.Stmt, error) {
	switch p.curTok.Type {
	case token.LET:
		return p.ParseLetStmt()
	default:
		return p.ParseExprStmt()
	}
}

func (p *Parser) ParseLetStmt() (*ast.LetStmt, error) {
	p.advance()

	if p.curTok.Type != token.IDENT {
		return nil, fmt.Errorf("expected %q, got %q", token.IDENT, p.curTok.Type)
	}

	name := p.curTok.Lit
	namePos := p.curTok.Pos

	p.advance()

	if p.curTok.Type != token.ASSIGN {
		return nil, fmt.Errorf("expected %q, got %q", token.ASSIGN, p.curTok.Type)
	}

	p.advance()

	expr, err := p.ParseExpr()
	if err != nil {
		return nil, err
	}

	if p.curTok.Type != token.SEMI {
		return nil, fmt.Errorf("expected %q, got %q", token.SEMI, p.curTok.Type)
	}

	p.advance()

	return &ast.LetStmt{Name: name, NamePos: namePos, Expr: expr}, nil
}

func (p *Parser) ParseExprStmt() (*ast.ExprStmt, error) {
	expr, err := p.ParseExpr()
	if err != nil {
		return nil, err
	}

	if p.curTok.Type != token.SEMI {
		return nil, fmt.Errorf("expected %q, got %q", token.SEMI, p.curTok.Type)
	}

	p.advance()

	return &ast.ExprStmt{Expr: expr}, nil
}

func (p *Parser) ParseExpr() (ast.Expr, error) {
	left, err := p.ParseTerm()
	if err != nil {
		return nil, err
	}

	for p.curTok.Type == token.ADD || p.curTok.Type == token.SUB {
		op := p.curTok.Lit

		p.advance()

		right, err := p.ParseExpr()
		if err != nil {
			return nil, err
		}

		left = &ast.BinExpr{Left: left, Op: op, Right: right}
	}

	return left, nil
}

func (p *Parser) ParseTerm() (ast.Expr, error) {
	left, err := p.ParseFactor()
	if err != nil {
		return nil, err
	}

	for p.curTok.Type == token.MUL || p.curTok.Type == token.DIV {
		op := p.curTok.Lit

		p.advance()

		right, err := p.ParseTerm()
		if err != nil {
			return nil, err
		}

		left = &ast.BinExpr{Left: left, Op: op, Right: right}
	}

	return left, nil
}

func (p *Parser) ParseFactor() (ast.Expr, error) {
	switch tok := p.curTok; tok.Type {
	case token.IDENT:
		node := &ast.Ident{Value: tok.Lit, ValuePos: tok.Pos}
		p.advance()
		return node, nil
	case token.INT:
		node := &ast.IntLit{Value: tok.Lit, ValuePos: tok.Pos}
		p.advance()
		return node, nil
	case token.SUB, token.NOT:
		node := &ast.UnaryExpr{Op: tok.Lit, OpPos: tok.Pos}

		p.advance()

		right, err := p.ParseFactor()
		if err != nil {
			return nil, err
		}

		node.Right = right

		return node, nil
	case token.LPAREN:
		p.advance()

		expr, err := p.ParseExpr()
		if err != nil {
			return nil, err
		}

		if p.curTok.Type != token.RPAREN {
			return nil, fmt.Errorf("expected %q, got %q", token.RPAREN, p.curTok.Type)
		}

		p.advance()

		return expr, nil
	default:
		return nil, fmt.Errorf("expected %q, %q, %q or %q, got %q",
			token.IDENT, token.INT, token.SUB, token.LPAREN, tok.Type)
	}
}

func (p *Parser) advance() {
	p.curTok = p.l.NextToken()
}
