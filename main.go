package main

import (
	"fmt"

	"github.com/toastsandwich/tinylang/ast"
	"github.com/toastsandwich/tinylang/lexer"
)

func main() {
	input := `
let x = 1;
let y = 2;
let z = x > y;
`
	lex := lexer.New(input)
	if lex == nil {
		fmt.Println("no input given to lexer")
		return
	}
	for {
		tok := lex.NextToken()
		fmt.Printf("%s: %s\n", tok.Type.String(), tok.Value)
		if tok.Type == lexer.EOF {
			break
		}
	}

	a := ast.NewAST(input)

	va := &ast.Identifier{Value: "var_a"}
	vb := &ast.Identifier{Value: "var_b"}

	leta := &ast.LetStatement{
		Identifier: va,
		Value:      "1000",
	}
	letb := &ast.LetStatement{
		Identifier: vb,
		Value:      "100",
	}
	iF := &ast.IfStatement{
		Condition: &ast.BinaryExpression{
			Left:     va,
			Right:    vb,
			Operator: "==",
		},
	}
	a.Root.AddToOutgoingNode(leta, letb, iF)
	for _, n := range a.Root.OutgoingNodes {
		fmt.Println(n.GenerateGo())
	}
}
