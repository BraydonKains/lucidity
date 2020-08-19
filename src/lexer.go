/*
lexer.go
Author: @BraydonKains

Credits:

"How to Write a Lexer in Go" by Aaron Raff
Link: https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go

"How to Write a Parser in Go" by Sugu Sougoumarane
Link: https://www.youtube.com/watch?v=NG0s3-s3whY
Repo: https://github.com/sougou/parser_tutorial
*/

package lexer

type Token int

const (
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI // I'm going to lex this for now but it's up in the air if I want to in the future

	ADD
	SUB
	MUL
	DIV
	ASSIGN

	IF
	UNLESS
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	DIV:    "/",
	ASSIGN: "=",

	IF:     "IF",
	UNLESS: "UNLESS",
}

func (t Token) String() string {
	return tokens[t]
}

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}
