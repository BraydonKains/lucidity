/*
lexer.go
Author: @BraydonKains

Credits:

"How to Write a Lexer in Go" by Aaron Raff
Link: https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go
I largely stole the structure of the lexer from here. I obviously had
to adapt it a fair bit to identify all the constructs of my language,
but without this article this lexer wouldn't have existed for a long time.

"How to Write a Parser in Go" by Sugu Sougoumarane
Link: https://www.youtube.com/watch?v=NG0s3-s3whY
Repo: https://github.com/sougou/parser_tutorial
This
*/

package lexer

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

const (
	// General
	EOF = iota
	ILLEGAL
	IDENT
	INT
	SEMI // I'm going to lex this for now but it's up in the air if I want to in the future
	LF
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	DQUOTE // Wait do I even need this
	COMMA
	STRLITERAL

	// Operators
	ADD
	SUB
	MUL
	DIV
	ASSIGN

	// Keywords
	IF
	UNLESS
	AND
	OR
	EQUALS
	TRUE
	FALSE

	// Types
	TYPEINT
	TYPESTR
	TYPEBOOL
)

var tokens = []string{
	EOF:        "EOF",
	ILLEGAL:    "ILLEGAL",
	LF:         "LF",
	IDENT:      "IDENT",
	INT:        "INT",
	SEMI:       ";",
	LPAREN:     "(",
	RPAREN:     ")",
	LBRACE:     "{",
	RBRACE:     "}",
	DQUOTE:     "\"",
	COMMA:      ",",
	STRLITERAL: "STRLITERAL",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	DIV:    "/",
	ASSIGN: "=",

	IF:     "IF",
	UNLESS: "UNLESS",
	AND:    "AND",
	OR:     "OR",
	EQUALS: "EQUALS",
	TRUE:   "TRUE",
	FALSE:  "FALSE",

	TYPEINT:  "TYPEINT",
	TYPESTR:  "TYPESTR",
	TYPEBOOL: "TYPEBOOL",
}

var terminal_symbols = map[rune]Token{
	';': SEMI,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	'+': ADD,
	'-': SUB,
	'*': MUL,
	'/': DIV,
	'=': ASSIGN,
	',': COMMA,
}

var keywords = map[string]Token{
	"if":      IF,
	"unless":  UNLESS,
	"and":     AND,
	"or":      OR,
	"equals":  EQUALS,
	"true":    TRUE,
	"false":   FALSE,
	"int":     TYPEINT,
	"string":  TYPESTR,
	"boolean": TYPEBOOL,
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

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		startPos := l.pos

		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			panic(err)
		}

		l.pos.column++

		if r == '\n' {
			l.nextLine()
			return startPos, LF, "\\n"
		}

		if token, ok := terminal_symbols[r]; ok {
			return l.pos, token, string(r)
		}

		if unicode.IsSpace(r) {
			continue
		}

		if unicode.IsDigit(r) {
			l.backup()
			literal := l.lexInt()
			return startPos, INT, literal
		}

		if unicode.IsLetter(r) || r == '_' {
			l.backup()
			literal, token := l.lexLetter()
			return startPos, token, literal
		}

		if r == '"' {
			if literal := l.lexStrLiteral(); literal[len(literal)-1:] == "\"" {
				return startPos, STRLITERAL, literal
			} else {
				return startPos, ILLEGAL, literal
			}
		}
	}
}

func (l *Lexer) nextLine() {
	l.pos.line++
	l.pos.column = 0
}

func (l *Lexer) backup() {
	err := l.reader.UnreadRune()
	if err != nil {
		panic(err)
	}

	l.pos.column--
}

func (l *Lexer) lexInt() string {
	var literal string

	for {
		r, _, err := l.reader.ReadRune()
		if err == io.EOF {
			return literal
		}

		if unicode.IsDigit(r) {
			literal += string(r)
		} else {
			l.backup()
			return literal
		}

		l.pos.column++
	}
}

func (l *Lexer) lexLetter() (string, Token) {
	var literal string

	for {
		r, _, err := l.reader.ReadRune()
		if err == io.EOF {
			return literal, IDENT
		}

		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			literal += string(r)
		} else {
			l.backup()
			if token, ok := keywords[literal]; ok {
				return literal, token
			} else {
				return literal, IDENT
			}
		}

		l.pos.column++
	}
}

func (l *Lexer) lexStrLiteral() string {
	literal := "\""

	for {
		r, _, err := l.reader.ReadRune()
		if err == io.EOF {
			return literal
		}

		literal += string(r)

		if r == '"' {
			return literal
		}
	}
}
