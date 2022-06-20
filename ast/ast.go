package ast

import (
	"bytes"
)

type Node interface {
	String() string
}

type Expr interface {
	Node
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Program struct {
	Stmts []Stmt
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Stmts {
		out.WriteString(stmt.String())
		out.WriteString(" ")
	}

	return out.String()
}

type Ident struct {
	Value string
}

func (i *Ident) exprNode()      {}
func (i *Ident) String() string { return i.Value }

type IntLit struct {
	Value string
}

func (il *IntLit) exprNode()      {}
func (il *IntLit) String() string { return il.Value }

type PrefixExpr struct {
	Op    string
	Right Expr
}

func (pe *PrefixExpr) exprNode() {}
func (pe *PrefixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Op)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

func (ie *InfixExpr) exprNode() {}
func (ie *InfixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Op)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type ExprStmt struct {
	Expr Expr
}

func (es *ExprStmt) stmtNode() {}
func (es *ExprStmt) String() string {
	var out bytes.Buffer

	out.WriteString(es.Expr.String())
	out.WriteString(";")

	return out.String()
}

type LetStmt struct {
	Name *Ident
	Expr Expr
}

func (ls *LetStmt) stmtNode() {}
func (ls *LetStmt) String() string {
	var out bytes.Buffer

	out.WriteString("let ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	out.WriteString(ls.Expr.String())
	out.WriteString(";")

	return out.String()
}

type BlockStmt struct {
	Stmts []Stmt
}

func (bs *BlockStmt) stmtNode() {}
func (bs *BlockStmt) String() string {
	var out bytes.Buffer

	out.WriteString("{ ")
	for _, stmt := range bs.Stmts {
		out.WriteString(stmt.String())
		out.WriteString(" ")
	}
	out.WriteString("}")

	return out.String()
}

type IfStmt struct {
	Expr      Expr
	IfBlock   *BlockStmt
	ElseBlock *BlockStmt
}

func (is *IfStmt) stmtNode() {}
func (is *IfStmt) String() string {
	var out bytes.Buffer

	return out.String()
}
