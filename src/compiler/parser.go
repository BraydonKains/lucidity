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
}

func (p *Parser) ReadToken() Token {
	token := p.tokens[p.pos]
	p.pos += 1
	return token
}

func ParseValidity(tokens []Token) bool {
	for _, token := range tokens {
		fmt.Println(token)
	}
	return true
}
