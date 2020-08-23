/*
parser.go
Author: @BraydonKains

I have no idea what I'm doing
*/

package compiler

import (
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
	end    int
}

func (p *Parser) ReadToken() Token {
	token := p.tokens[p.pos]
	p.pos += 1
	return token
}

func newParser(tokens []Token) Parser {
	return Parser{tokens, 0, len(tokens)}
}

func ParseValidity(tokens []Token) bool {
	for _, token := range tokens {
		fmt.Println(token)
	}

	// p := newParser(tokens)

	for {
		// token := p.ReadToken()
		// if token.IsType() {
		// 	p.parseAssignment()
		// }

		return true
	}
}
