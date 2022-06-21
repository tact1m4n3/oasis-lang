package parser

import (
	"fmt"
	"oasis/ast"
	"oasis/lexer"
	"oasis/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.EQ:     EQUALS,
	token.NEQ:    EQUALS,
	token.LT:     LESSGREATER,
	token.GT:     LESSGREATER,
	token.ADD:    SUM,
	token.SUB:    SUM,
	token.MUL:    PRODUCT,
	token.DIV:    PRODUCT,
	token.LPAREN: CALL,
}

type (
	prefixParseFn func() (ast.Expr, error)
	infixParseFn  func(ast.Expr) (ast.Expr, error)
)

type Parser struct {
	l *lexer.Lexer

	tok token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.advance()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdent)
	p.registerPrefix(token.INT, p.parseIntLit)
	p.registerPrefix(token.SUB, p.parsePrefixExpr)
	p.registerPrefix(token.NOT, p.parsePrefixExpr)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpr)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.ADD, p.parseInfixExpr)
	p.registerInfix(token.SUB, p.parseInfixExpr)
	p.registerInfix(token.MUL, p.parseInfixExpr)
	p.registerInfix(token.DIV, p.parseInfixExpr)
	p.registerInfix(token.EQ, p.parseInfixExpr)
	p.registerInfix(token.NEQ, p.parseInfixExpr)
	p.registerInfix(token.LT, p.parseInfixExpr)
	p.registerInfix(token.GT, p.parseInfixExpr)
	p.registerInfix(token.LPAREN, p.parseCallExpr)

	return p
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	var stmts []ast.Stmt
	for p.tok.Type != token.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return &ast.Program{Stmts: stmts}, nil
}

func (p *Parser) parseStmt() (ast.Stmt, error) {
	switch p.tok.Type {
	case token.LET:
		return p.parseLetStmt()
	case token.LBRACE:
		return p.parseBlockStmt()
	case token.IF:
		return p.parseIfStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	case token.FUNC:
		return p.parseFuncStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseExprStmt() (*ast.ExprStmt, error) {
	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	if p.tok.Type != token.SEMI {
		return nil, fmt.Errorf("expected %q, got %q", token.SEMI, p.tok.Type)
	}

	p.advance()

	return &ast.ExprStmt{Expr: expr}, nil
}

func (p *Parser) parseExpr(prec int) (ast.Expr, error) {
	prefix, ok := p.prefixParseFns[p.tok.Type]
	if !ok {
		return nil, fmt.Errorf("expected %q, %q, %q, %q or %q, got %q",
			token.IDENT, token.INT, token.SUB, token.NOT, token.LPAREN, p.tok.Type)
	}

	left, err := prefix()
	if err != nil {
		return nil, err
	}

	for p.tok.Type != token.SEMI && prec < p.getPrecedence() {
		infix, ok := p.infixParseFns[p.tok.Type]
		if !ok {
			return nil, fmt.Errorf("this should never have happened. what did you do???")
		}

		left, err = infix(left)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) parseIdent() (ast.Expr, error) {
	if p.tok.Type != token.IDENT {
		return nil, fmt.Errorf("expected %q, got %q", token.IDENT, p.tok.Type)
	}
	node := &ast.Ident{Value: p.tok.Lit}
	p.advance()
	return node, nil
}

func (p *Parser) parseIntLit() (ast.Expr, error) {
	if p.tok.Type != token.INT {
		return nil, fmt.Errorf("expected %q, got %q", token.INT, p.tok.Type)
	}
	node := &ast.IntLit{Value: p.tok.Lit}
	p.advance()
	return node, nil
}

func (p *Parser) parsePrefixExpr() (ast.Expr, error) {
	expr := &ast.PrefixExpr{Op: p.tok.Lit}
	p.advance()

	right, err := p.parseExpr(PREFIX)
	if err != nil {
		return nil, err
	}

	expr.Right = right

	return expr, nil
}

func (p *Parser) parseInfixExpr(left ast.Expr) (ast.Expr, error) {
	expr := &ast.InfixExpr{
		Left: left,
		Op:   p.tok.Lit,
	}

	prec := p.getPrecedence()
	p.advance()

	right, err := p.parseExpr(prec)
	if err != nil {
		return nil, err
	}

	expr.Right = right

	return expr, nil
}

func (p *Parser) parseGroupedExpr() (ast.Expr, error) {
	if p.tok.Type != token.LPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.LPAREN, p.tok.Type)
	}
	p.advance()

	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	if p.tok.Type != token.RPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.RPAREN, p.tok.Type)
	}
	p.advance()

	return expr, nil
}

func (p *Parser) parseCallExpr(left ast.Expr) (ast.Expr, error) {
	if p.tok.Type != token.LPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.LPAREN, p.tok.Type)
	}
	p.advance()

	var args []ast.Expr
	for p.tok.Type != token.EOF {
		expr, err := p.parseExpr(LOWEST)
		if err != nil {
			return nil, err
		}
		args = append(args, expr)

		if p.tok.Type == token.RPAREN {
			break
		} else if p.tok.Type != token.COMMA {
			return nil, fmt.Errorf("expected %q, got %q", token.COMMA, p.tok.Type)
		}
		p.advance()
	}

	if p.tok.Type != token.RPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.RPAREN, p.tok.Type)
	}
	p.advance()

	return &ast.CallExpr{Left: left, Args: args}, nil
}

func (p *Parser) parseLetStmt() (*ast.LetStmt, error) {
	if p.tok.Type != token.LET {
		return nil, fmt.Errorf("expected %q, got %q", token.LET, p.tok.Type)
	}
	p.advance()

	if p.tok.Type != token.IDENT {
		return nil, fmt.Errorf("expected %q, got %q", token.IDENT, p.tok.Type)
	}
	name := &ast.Ident{Value: p.tok.Lit}
	p.advance()

	if p.tok.Type != token.ASSIGN {
		return nil, fmt.Errorf("expected %q, got %q", token.ASSIGN, p.tok.Type)
	}
	p.advance()

	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	if p.tok.Type != token.SEMI {
		return nil, fmt.Errorf("expected %q, got %q", token.SEMI, p.tok.Type)
	}
	p.advance()

	return &ast.LetStmt{Name: name, Expr: expr}, nil
}

func (p *Parser) parseBlockStmt() (*ast.BlockStmt, error) {
	if p.tok.Type != token.LBRACE {
		return nil, fmt.Errorf("expected %q, got %q", token.LBRACE, p.tok.Type)
	}
	p.advance()

	var stmts []ast.Stmt
	for p.tok.Type != token.RBRACE && p.tok.Type != token.EOF {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}

	if p.tok.Type != token.RBRACE {
		return nil, fmt.Errorf("expected %q, got %q", token.RBRACE, p.tok.Type)
	}
	p.advance()

	return &ast.BlockStmt{Stmts: stmts}, nil
}

func (p *Parser) parseIfStmt() (*ast.IfStmt, error) {
	if p.tok.Type != token.IF {
		return nil, fmt.Errorf("expected %q, got %q", token.IF, p.tok.Type)
	}
	p.advance()

	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	ifBlock, err := p.parseBlockStmt()
	if err != nil {
		return nil, err
	}

	if p.tok.Type != token.ELSE {
		return &ast.IfStmt{Expr: expr, IfBlock: ifBlock}, nil
	}
	p.advance()

	elseBlock, err := p.parseBlockStmt()
	if err != nil {
		return nil, err
	}

	return &ast.IfStmt{Expr: expr, IfBlock: ifBlock, ElseBlock: elseBlock}, nil
}

func (p *Parser) parseReturnStmt() (*ast.ReturnStmt, error) {
	if p.tok.Type != token.RETURN {
		return nil, fmt.Errorf("expected %q, got %q", token.RETURN, p.tok.Type)
	}
	p.advance()

	if p.tok.Type == token.SEMI {
		p.advance()
		return &ast.ReturnStmt{}, nil
	}

	expr, err := p.parseExpr(LOWEST)
	if err != nil {
		return nil, err
	}

	if p.tok.Type != token.SEMI {
		return nil, fmt.Errorf("expected %q, got %q", token.SEMI, p.tok.Type)
	}
	p.advance()

	return &ast.ReturnStmt{Expr: expr}, nil
}

func (p *Parser) parseFuncStmt() (*ast.FuncStmt, error) {
	if p.tok.Type != token.FUNC {
		return nil, fmt.Errorf("expected %q, got %q", token.FUNC, p.tok.Type)
	}
	p.advance()

	if p.tok.Type != token.IDENT {
		return nil, fmt.Errorf("expected %q, got %q", token.IDENT, p.tok.Type)
	}
	name := &ast.Ident{Value: p.tok.Lit}
	p.advance()

	if p.tok.Type != token.LPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.LPAREN, p.tok.Type)
	}
	p.advance()

	argNames := []*ast.Ident{}
	for p.tok.Type != token.RPAREN && p.tok.Type != token.EOF {
		argNames = append(argNames, &ast.Ident{Value: p.tok.Lit})
		p.advance()

		if p.tok.Type == token.RPAREN {
			break
		} else if p.tok.Type != token.COMMA {
			return nil, fmt.Errorf("expected %q, got %q", token.COMMA, p.tok.Type)
		}
		p.advance()
	}

	if p.tok.Type != token.RPAREN {
		return nil, fmt.Errorf("expected %q, got %q", token.RPAREN, p.tok.Type)
	}
	p.advance()

	body, err := p.parseBlockStmt()
	if err != nil {
		return nil, err
	}

	return &ast.FuncStmt{Name: name, ArgNames: argNames, Body: body}, nil
}

func (p *Parser) advance() {
	p.tok = p.l.NextToken()
}

func (p *Parser) registerPrefix(tok token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tok] = fn
}

func (p *Parser) registerInfix(tok token.TokenType, fn infixParseFn) {
	p.infixParseFns[tok] = fn
}

func (p *Parser) getPrecedence() int {
	if prec, ok := precedences[p.tok.Type]; ok {
		return prec
	}
	return LOWEST
}
