package ast

import (
	"log"

	"github.com/toastsandwich/tinylang/lexer"
)

type ASTNode struct {
	IncomingNodes []Node
	OutgoingNodes []Node
}

func (a *ASTNode) AddToIncomingNodes(node ...Node) {
	a.IncomingNodes = append(a.IncomingNodes, node...)
}

func (a *ASTNode) AddToOutgoingNode(node ...Node) {
	a.OutgoingNodes = append(a.OutgoingNodes, node...)
}

type AST struct {
	Lexer *lexer.Lexer
	Root  *ASTNode
}

func NewAST(in string) *AST {
	l := lexer.New(in)
	if l == nil {
		log.Fatal("lexer is nil, may be no input was passed...")
	}
	return &AST{
		Lexer: l,
		Root:  new(ASTNode),
	}
}

func (a *AST) Unamed() {
	for {
		t := a.Lexer.NextToken()
		for t.Type != lexer.DLIM {
			switch t.Type {
			case lexer.KWRD:
			}
		}
		if t.Type == lexer.EOF {
			break
		}
	}
}
