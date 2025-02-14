package main

import (
	"fmt"

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
}
