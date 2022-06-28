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
	ASSIGN
	LOR
	LAND
	XOR
	OR
	AND
	EQUALS
	COMPARISON
	SHIFT
	TERM
	FACTOR
	PREFIX
	CALL
)

var precedences = map[token.Token]int{
	token.ASSIGN:   ASSIGN,
	token.LAND:     LAND,
	token.LOR:      LOR,
	token.AND:      AND,
	token.OR:       OR,
	token.XOR:      XOR,
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       COMPARISON,
	token.LTE:      COMPARISON,
	token.GT:       COMPARISON,
	token.GTE:      COMPARISON,
	token.LSHIFT:   SHIFT,
	token.RSHIFT:   SHIFT,
	token.PLUS:     TERM,
	token.MINUS:    TERM,
	token.ASTERISK: FACTOR,
	token.SLASH:    FACTOR,
	token.MOD:      FACTOR,
	token.LPAREN:   CALL,
}

type (
	prefixParseFn func() ast.Expr
	infixParseFn  func(ast.Expr) ast.Expr
)

type Parser struct {
	l   *lexer.Lexer
	err error

	tok token.Token
	lit string

	prefixParseFns map[token.Token]prefixParseFn
	infixParseFns  map[token.Token]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.advance()

	p.prefixParseFns = make(map[token.Token]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdent)
	p.registerPrefix(token.INT, p.parseIntLit)
	p.registerPrefix(token.MINUS, p.parsePrefixExpr)
	p.registerPrefix(token.NOT, p.parsePrefixExpr)
	p.registerPrefix(token.BANG, p.parsePrefixExpr)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpr)
	p.registerPrefix(token.LBRACE, p.parseBlockExpr)
	p.registerPrefix(token.IF, p.parseIfExpr)
	p.registerPrefix(token.FUNC, p.parseFuncLit)

	p.infixParseFns = make(map[token.Token]infixParseFn)
	p.registerInfix(token.ASSIGN, p.parseInfixExpr)
	p.registerInfix(token.LAND, p.parseInfixExpr)
	p.registerInfix(token.LOR, p.parseInfixExpr)
	p.registerInfix(token.AND, p.parseInfixExpr)
	p.registerInfix(token.OR, p.parseInfixExpr)
	p.registerInfix(token.XOR, p.parseInfixExpr)
	p.registerInfix(token.EQ, p.parseInfixExpr)
	p.registerInfix(token.NEQ, p.parseInfixExpr)
	p.registerInfix(token.LT, p.parseInfixExpr)
	p.registerInfix(token.LTE, p.parseInfixExpr)
	p.registerInfix(token.GT, p.parseInfixExpr)
	p.registerInfix(token.GTE, p.parseInfixExpr)
	p.registerInfix(token.LSHIFT, p.parseInfixExpr)
	p.registerInfix(token.RSHIFT, p.parseInfixExpr)
	p.registerInfix(token.PLUS, p.parseInfixExpr)
	p.registerInfix(token.MINUS, p.parseInfixExpr)
	p.registerInfix(token.ASTERISK, p.parseInfixExpr)
	p.registerInfix(token.SLASH, p.parseInfixExpr)
	p.registerInfix(token.MOD, p.parseInfixExpr)
	p.registerInfix(token.LPAREN, p.parseCallExpr)

	return p
}

func (p *Parser) Error() error {
	return p.err
}

func (p *Parser) ParseProgram() *ast.Program {
	stmts := []ast.Stmt{}

	for p.tok != token.EOF {
		stmt := p.parseStmt()
		if stmt == nil {
			return nil
		}
		stmts = append(stmts, stmt)
	}

	return &ast.Program{Stmts: stmts}
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.tok {
	case token.LET:
		return p.parseLetStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *Parser) parseExprStmt() ast.Stmt {
	expr := p.parseExpr(LOWEST)
	if expr == nil {
		return nil
	}

	if !p.expect(token.SEMI) {
		return nil
	}
	p.advance()

	return &ast.ExprStmt{Expr: expr}
}

func (p *Parser) parseLetStmt() ast.Stmt {
	p.advance()

	if !p.expect(token.IDENT) {
		return nil
	}
	name := &ast.Ident{Value: p.lit}
	p.advance()

	if !p.expect(token.ASSIGN) {
		return nil
	}
	p.advance()

	value := p.parseExpr(LOWEST)
	if value == nil {
		return nil
	}

	if !p.expect(token.SEMI) {
		return nil
	}
	p.advance()

	return &ast.LetStmt{Name: name, Value: value}
}

func (p *Parser) parseReturnStmt() ast.Stmt {
	p.advance()

	if p.tok == token.SEMI {
		p.advance()
		return &ast.ReturnStmt{}
	}

	value := p.parseExpr(LOWEST)
	if value == nil {
		return nil
	}

	if !p.expect(token.SEMI) {
		return nil
	}
	p.advance()

	return &ast.ReturnStmt{Value: value}
}

func (p *Parser) parseExpr(prec int) ast.Expr {
	prefix := p.prefixParseFns[p.tok]
	if prefix == nil {
		p.err = fmt.Errorf("expected %q, %q, %q, %q, %q or %q, got %q",
			token.IDENT, token.INT, token.MINUS, token.NOT, token.BANG, token.LPAREN, p.tok)
		return nil
	}

	left := prefix()
	if left == nil {
		return nil
	}

	for p.tok != token.RPAREN && p.tok != token.SEMI && p.tok != token.EOF && prec < p.curPrecedence() {
		infix := p.infixParseFns[p.tok]
		if infix == nil {
			return nil
		}

		left = infix(left)
		if left == nil {
			return nil
		}
	}

	return left
}

func (p *Parser) parseIdent() ast.Expr {
	node := &ast.Ident{Value: p.lit}
	p.advance()
	return node
}

func (p *Parser) parseIntLit() ast.Expr {
	node := &ast.IntLit{Value: p.lit}
	p.advance()
	return node
}

func (p *Parser) parsePrefixExpr() ast.Expr {
	op := p.tok
	p.advance()

	right := p.parseExpr(PREFIX)
	if right == nil {
		return nil
	}

	return &ast.PrefixExpr{Op: op, Right: right}
}

func (p *Parser) parseInfixExpr(left ast.Expr) ast.Expr {
	op := p.tok
	prec := p.curPrecedence()
	p.advance()

	right := p.parseExpr(prec)
	if right == nil {
		return nil
	}

	return &ast.InfixExpr{Left: left, Op: op, Right: right}
}

func (p *Parser) parseGroupedExpr() ast.Expr {
	p.advance()

	expr := p.parseExpr(LOWEST)
	if expr == nil {
		return expr
	}

	if !p.expect(token.RPAREN) {
		return nil
	}
	p.advance()

	return expr
}

func (p *Parser) parseCallExpr(left ast.Expr) ast.Expr {
	p.advance()

	args := p.parseCallArgs()
	if args == nil {
		return nil
	}

	if !p.expect(token.RPAREN) {
		return nil
	}
	p.advance()

	return &ast.CallExpr{Func: left, Args: args}
}

func (p *Parser) parseCallArgs() []ast.Expr {
	args := []ast.Expr{}

	if p.tok == token.RPAREN {
		return args
	}

	arg := p.parseExpr(LOWEST)
	if arg == nil {
		return nil
	}
	args = append(args, arg)
	for p.tok == token.COMMA {
		p.advance()

		if p.tok == token.RPAREN {
			break
		}

		arg = p.parseExpr(LOWEST)
		if arg == nil {
			return nil
		}
		args = append(args, arg)
	}

	return args
}

func (p *Parser) parseBlockExpr() ast.Expr {
	p.advance()

	stmts := []ast.Stmt{}
	for p.tok != token.RBRACE && p.tok != token.EOF {
		stmt := p.parseStmt()
		if stmt == nil {
			return nil
		}
		stmts = append(stmts, stmt)
	}

	if !p.expect(token.RBRACE) {
		return nil
	}
	p.advance()

	return &ast.BlockExpr{Stmts: stmts}
}

func (p *Parser) parseIfExpr() ast.Expr {
	p.advance()

	condition := p.parseExpr(LOWEST)
	if condition == nil {
		return nil
	}

	trueCase := p.parseExpr(LOWEST)
	if trueCase == nil {
		return nil
	}

	if p.tok == token.ELSE {
		p.advance()

		falseCase := p.parseExpr(LOWEST)
		if falseCase == nil {
			return nil
		}

		return &ast.IfExpr{Condition: condition, TrueCase: trueCase, FalseCase: falseCase}
	}

	return &ast.IfExpr{Condition: condition, TrueCase: trueCase}
}

func (p *Parser) parseFuncLit() ast.Expr {
	p.advance()

	if !p.expect(token.LPAREN) {
		return nil
	}
	p.advance()

	params := p.parseFuncParams()
	if params == nil {
		return nil
	}

	if !p.expect(token.RPAREN) {
		return nil
	}
	p.advance()

	body := p.parseBlockExpr()
	if body == nil {
		return nil
	}

	return &ast.FuncLit{Params: params, Body: body}
}

func (p *Parser) parseFuncParams() []*ast.Ident {
	params := []*ast.Ident{}

	if p.tok == token.RPAREN {
		return params
	}

	if !p.expect(token.IDENT) {
		return nil
	}
	params = append(params, &ast.Ident{Value: p.lit})
	p.advance()

	for p.tok == token.COMMA {
		p.advance()

		if p.tok == token.RPAREN {
			break
		}

		if !p.expect(token.IDENT) {
			return nil
		}
		params = append(params, &ast.Ident{Value: p.lit})
		p.advance()
	}

	return params
}

func (p *Parser) advance() {
	p.tok, p.lit = p.l.NextToken()
}

func (p *Parser) registerPrefix(tok token.Token, fn prefixParseFn) {
	p.prefixParseFns[tok] = fn
}

func (p *Parser) registerInfix(tok token.Token, fn infixParseFn) {
	p.infixParseFns[tok] = fn
}

func (p *Parser) curPrecedence() int {
	if prec, ok := precedences[p.tok]; ok {
		return prec
	}
	return LOWEST
}

func (p *Parser) expect(tok token.Token) bool {
	if p.tok != tok {
		p.err = fmt.Errorf("expected %q, got %q", tok, p.tok)
		return false
	}
	return true
}
