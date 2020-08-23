/*
parser.go
Author: @BraydonKains

I have no idea what I'm doing
*/

package parser

import (
	"../lexer/"
	"fmt"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func (p *Parser) ReadToken() lexer.Token {
	token := p.tokens[p.pos]
	p.pos += 1
	return token
}

func ParseValidity(tokens []lexer.Token) bool {
	for _, token := range tokens {
		fmt.Println(token)
	}
	return true
}
