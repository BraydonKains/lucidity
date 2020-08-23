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

package compiler

import (
	"bufio"
	"io"
	"unicode"
)

type Lexer struct {
	filename string
	pos      Position
	reader   *bufio.Reader
}

func NewLexer(filename string, reader *bufio.Reader) *Lexer {
	return &Lexer{
		filename: filename,
		pos:      Position{line: 1, column: 0},
		reader:   reader,
	}
}

func (l *Lexer) Lex() Token {
	file := l.filename
	for {
		startPos := l.pos

		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{file, startPos, EOF, ""}
			}

			panic(err)
		}

		l.pos.column++

		if r == '\n' {
			l.nextLine()
			return Token{file, startPos, LF, "\\n"}
		}

		if token, ok := terminalSymbols[r]; ok {
			return Token{file, startPos, token, string(r)}
		}

		if unicode.IsSpace(r) {
			continue
		}

		if unicode.IsDigit(r) {
			l.backup()
			literal := l.lexInt()
			return Token{file, startPos, INT, literal}
		}

		if unicode.IsLetter(r) || r == '_' {
			l.backup()
			literal, token := l.lexLetter()
			return Token{file, startPos, token, literal}
		}

		if r == '"' {
			if literal := l.lexStrLiteral(); literal[len(literal)-1:] == "\"" {
				return Token{file, startPos, STRLITERAL, literal}
			} else {
				return Token{file, startPos, ILLEGAL, literal}
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

func (l *Lexer) lexLetter() (string, TokenId) {
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
