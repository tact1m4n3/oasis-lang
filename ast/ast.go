package ast

import (
	"bytes"
)

type Node interface {
	StartPos() int
	EndPos() int
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

func (p *Program) StartPos() int {
	if len(p.Stmts) == 0 {
		return 0
	}
	return p.Stmts[0].StartPos()
}

func (p *Program) EndPos() int {
	if len(p.Stmts) == 0 {
		return 0
	}
	return p.Stmts[len(p.Stmts)-1].EndPos()
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Stmts {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	return out.String()
}

type Ident struct {
	Value    string
	ValuePos int
}

func (i *Ident) exprNode()      {}
func (i *Ident) StartPos() int  { return i.ValuePos }
func (i *Ident) EndPos() int    { return i.ValuePos + len(i.Value) }
func (i *Ident) String() string { return i.Value }

type IntLit struct {
	Value    string
	ValuePos int
}

func (il *IntLit) exprNode()      {}
func (il *IntLit) StartPos() int  { return il.ValuePos }
func (il *IntLit) EndPos() int    { return il.ValuePos + len(il.Value) }
func (il *IntLit) String() string { return il.Value }

type UnaryExpr struct {
	Op    string
	OpPos int
	Right Expr
}

func (ue *UnaryExpr) exprNode()     {}
func (ue *UnaryExpr) StartPos() int { return ue.OpPos }
func (ue *UnaryExpr) EndPos() int   { return ue.Right.EndPos() }
func (ue *UnaryExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ue.Op)
	out.WriteString(ue.Right.String())
	out.WriteString(")")

	return out.String()
}

type BinExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

func (be *BinExpr) exprNode()     {}
func (be *BinExpr) StartPos() int { return be.Left.StartPos() }
func (be *BinExpr) EndPos() int   { return be.Right.EndPos() }
func (be *BinExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(be.Left.String())
	out.WriteString(" ")
	out.WriteString(be.Op)
	out.WriteString(" ")
	out.WriteString(be.Right.String())
	out.WriteString(")")

	return out.String()
}

type ExprStmt struct {
	Expr Expr
}

func (es *ExprStmt) stmtNode()     {}
func (es *ExprStmt) StartPos() int { return es.Expr.StartPos() }
func (es *ExprStmt) EndPos() int   { return es.Expr.EndPos() }
func (es *ExprStmt) String() string {
	var out bytes.Buffer

	out.WriteString(es.Expr.String())
	out.WriteString(";")

	return out.String()
}

type LetStmt struct {
	Name    string
	NamePos int
	Expr    Expr
}

func (ls *LetStmt) stmtNode()     {}
func (ls *LetStmt) StartPos() int { return ls.NamePos }
func (ls *LetStmt) EndPos() int   { return ls.Expr.EndPos() }
func (ls *LetStmt) String() string {
	var out bytes.Buffer

	out.WriteString("let ")
	out.WriteString(ls.Name)
	out.WriteString(" = ")
	out.WriteString(ls.Expr.String())
	out.WriteString(";")

	return out.String()
}
