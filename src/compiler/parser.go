/*
parser.go
Author: @BraydonKains
*/

package compiler

import (
	// "errors"
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
	end    int
}

func newParser(tokens []Token) Parser {
	return Parser{tokens, 0, len(tokens)}
}

func (p *Parser) readToken() Token {
	token := p.tokens[p.pos]
	p.pos += 1
	return token
}

func (p *Parser) peekToken() Token {
	return p.tokens[p.pos]
}

func ParseValidity(tokens []Token) bool {
	for _, token := range tokens {
		fmt.Println(token)
	}

	parser := newParser(tokens)

	for parser.pos <= parser.end {
		token := parser.peekToken()
		switch {
		case token.IsType():
			err := parser.parseAssignment()
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
	}

	return true
}
