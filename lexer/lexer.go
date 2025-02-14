package lexer

import (
	"unicode"
)

type TokenType int

func (t TokenType) String() string {
	return []string{
		"end of file",
		"illegal",
		"white space",
		"delimiter",
		"integar literal",
		"left parenthesis",
		"right parenthesis",
		"left bracket",
		"right bracket",
		"operator",
		"identifier",
		"keyword",
	}[t]
}

/*
let x = 100;
let y = 10;
let z = x + y;
if (z > 100) {
  z = 100;
}
*/

const (
	EOF TokenType = iota
	ILL
	WHSP
	DLIM
	INTL
	LPAR
	RPAR
	LBRA
	RBRA
	OPER
	IDNT
	KWRD
)

var (
	Operator = map[string]struct{}{
		"=":  {},
		">":  {},
		"<":  {},
		"!":  {},
		"&":  {},
		"|":  {},
		"+":  {},
		"-":  {},
		"*":  {},
		"/":  {},
		"%":  {},
		">=": {},
		"<=": {},
		"==": {},
		"++": {},
		"--": {},
		"+=": {},
		"-=": {},
		"*=": {},
		"/=": {},
		"&&": {},
		"||": {},
	}

	Keyword = map[string]struct{}{
		"let": {},
	}
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Input   string
	Current int
	Next    int
}

func New(input string) *Lexer {
	l := &Lexer{Input: input}
	if len(l.Input) == 0 {
		return nil
	}
	if len(l.Input) > 1 {
		l.Next = 1
	}
	return l
}

// reads value where the cursor is
func (l *Lexer) Read() byte {
	if l.Current >= len(l.Input) {
		return 0
	}
	return l.Input[l.Current]
}

// reads the immediate next value where cursor is
func (l *Lexer) ReadNext() byte {
	if l.Next >= len(l.Input) {
		return 0 // EOF
	}
	return l.Input[l.Next]
}

// first advances the cursor and then reads the value
func (l *Lexer) Advance() {
	l.Current++
	l.Next = l.Current + 1
}

func (l *Lexer) NextToken() Token {
	defer l.Advance()

	if unicode.IsSpace(rune(l.Read())) {
		return Token{
			Type: WHSP,
		}
	}

	// check for delimiter, EOF, Parenthesis and Brackets
	switch l.Read() {
	case 0:
		return Token{
			Type: EOF,
		}
	case ';':
		return Token{
			Type:  DLIM,
			Value: ";",
		}
	case '(':
		return Token{
			Type:  LPAR,
			Value: "(",
		}
	case '{':
		return Token{
			Type:  LBRA,
			Value: "{",
		}
	case ')':
		return Token{
			Type:  RPAR,
			Value: ")",
		}
	case '}':
		return Token{
			Type:  RBRA,
			Value: "}",
		}
	}

	if _, ok := Operator[string(l.Read())]; ok {
		val := string(l.Read())
		nxt := string(l.ReadNext())
		if _, ok := Operator[val+nxt]; ok {
			val += nxt
			l.Advance()
		}
		return Token{
			Type:  OPER,
			Value: val,
		}
	}

	if unicode.IsLetter(rune(l.Read())) {
		val := string(l.Read())
		for unicode.IsLetter(rune(l.ReadNext())) {
			l.Advance()
			val += string(l.Read())
		}
		if _, ok := Keyword[val]; ok {
			return Token{
				Type:  KWRD,
				Value: val,
			}
		}
		return Token{
			Type:  IDNT,
			Value: val,
		}
	}

	if unicode.IsDigit(rune(l.Read())) {
		val := string(l.Read())
		for unicode.IsDigit(rune(l.ReadNext())) {
			l.Advance()
			val += string(l.Read())
		}
		return Token{
			Type:  INTL,
			Value: val,
		}
	}

	return Token{Type: ILL}
}
