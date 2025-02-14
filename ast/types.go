package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	GenerateGo() string
}

type Statement interface {
	Node
	IsStatement()
}

func (*LetStatement) IsStatement() {}
func (*IfStatement) IsStatement()  {}

type Expression interface {
	Node
	IsExpression()
}

func (*BinaryExpression) IsExpression() {}
func (*Identifier) IsExpression()       {}
func (*UnaryExpression) IsExpression()  {}

type LetStatement struct {
	Identifier *Identifier
	Value      string
}

func (l *LetStatement) GenerateGo() string {
	return fmt.Sprintf("%s := %s\n", l.Identifier.Value, l.Value)
}

type IfStatement struct {
	Condition  Expression
	Statements []Statement
}

func (i *IfStatement) GenerateGo() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("if %s {\n", i.Condition.GenerateGo()))
	for _, s := range i.Statements {
		builder.WriteString(s.GenerateGo())
	}
	builder.WriteString("}\n")
	return builder.String()
}

// Note do not add \n in any Expressions

type UnaryExpression struct {
	Operand  Expression
	Operator string
}

func (u *UnaryExpression) GenerateGo() string {
	return fmt.Sprintf("%s%s", u.Operand.GenerateGo(), u.Operator)
}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (b *BinaryExpression) GenerateGo() string {
	return fmt.Sprintf("%s %s %s", b.Left.GenerateGo(), b.Operator, b.Right.GenerateGo())
}

type Identifier struct {
	Value string
}

func (i *Identifier) GenerateGo() string {
	return fmt.Sprintf("%s", i.Value)
}
