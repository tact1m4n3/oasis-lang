package ast

import (
	"bytes"
	"oasis/token"
)

type Node interface {
	String() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
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
	Name  *Ident
	Value Expr
}

func (ls *LetStmt) stmtNode() {}
func (ls *LetStmt) String() string {
	var out bytes.Buffer

	out.WriteString("let ")
	out.WriteString(ls.Name.String())
	if ls.Value != nil {
		out.WriteString(" = ")
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStmt struct {
	Value Expr
}

func (rs *ReturnStmt) stmtNode() {}
func (rs *ReturnStmt) String() string {
	var out bytes.Buffer

	out.WriteString("return")
	if rs.Value != nil {
		out.WriteString(" ")
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")

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
	Op    token.Token
	Right Expr
}

func (pe *PrefixExpr) exprNode() {}
func (pe *PrefixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Op.String())
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpr struct {
	Left  Expr
	Op    token.Token
	Right Expr
}

func (ie *InfixExpr) exprNode() {}
func (ie *InfixExpr) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Op.String())
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type CallExpr struct {
	Func Expr
	Args []Expr
}

func (ce *CallExpr) exprNode() {}
func (ce *CallExpr) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Func.String())
	out.WriteString("(")
	for _, arg := range ce.Args {
		out.WriteString(arg.String())
		out.WriteString(", ")
	}
	out.WriteString(")")

	return out.String()
}

type BlockExpr struct {
	Stmts []Stmt
}

func (be *BlockExpr) exprNode() {}
func (be *BlockExpr) String() string {
	var out bytes.Buffer

	out.WriteString("{ ")
	for _, stmt := range be.Stmts {
		out.WriteString(stmt.String())
		out.WriteString(" ")
	}
	out.WriteString("}")

	return out.String()
}

type IfExpr struct {
	Condition Expr
	TrueCase  Expr
	FalseCase Expr
}

func (ie *IfExpr) exprNode() {}
func (ie *IfExpr) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.TrueCase.String())
	if ie.FalseCase != nil {
		out.WriteString(" else ")
		out.WriteString(ie.FalseCase.String())
	}

	return out.String()
}

type FuncLit struct {
	Params []*Ident
	Body   Expr
}

func (fl *FuncLit) exprNode() {}
func (fl *FuncLit) String() string {
	var out bytes.Buffer

	out.WriteString("func(")
	for _, param := range fl.Params {
		out.WriteString(param.String())
		out.WriteString(", ")
	}
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
