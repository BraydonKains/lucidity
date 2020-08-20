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
	LPAREN
	RPAREN
	LBRACE
	RBRACE

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
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",
	LPAREN:  "(",
	RPAREN:  ")",
	LBRACE:  "{",
	RBRACE:  "}",

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
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			panic(err)
		}

		l.pos.column++

		switch r {
		case '\n':
			l.nextLine()
		case ';':
			return l.pos, SEMI, ";"
		case '(':
			return l.pos, LPAREN, "("
		case ')':
			return l.pos, RPAREN, ")"
		case '{':
			return l.pos, LBRACE, "{"
		case '}':
			return l.pos, RBRACE, "}"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				literal := l.lexInt()
				return startPos, INT, literal
			} else if unicode.IsLetter(r) || r == '_' {
				startPos := l.pos
				l.backup()
				literal, token := l.lexLetter()
				return startPos, token, literal
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
