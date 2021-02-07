package compiler

import (
	"fmt"
)

type TokenId int

type Position struct {
	line   int
	column int
}

type Token struct {
	file    string
	pos     Position
	id      TokenId
	literal string
}

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

var terminalSymbols = map[rune]TokenId{
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

var controlKeywords = map[string]TokenId{
	"if":     IF,
	"unless": UNLESS,
	"and":    AND,
	"or":     OR,
	"not":    NOT,
	"equals": EQUALS,
	"true":   TRUE,
	"false":  FALSE,
}

var typeKeywords = map[string]TokenId{
	"int":     TYPEINT,
	"string":  TYPESTR,
	"boolean": TYPEBOOL,
}

func (t TokenId) String() string {
	return tokens[t]
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %d %s %s ", t.file, t.pos, t.id, t.literal)
}

func (t Token) Is(id TokenId) bool {
	return t.id == id
}

func (t Token) IsKeyword() bool {
	_, controlOk := controlKeywords[t.id]
	_, typeOk := typeKeywords[t.id]

	return controlOk || typeOk
}

func (t Token) IsType() bool {
	_, ok := typeKeywords[t.id]
	return ok
}
