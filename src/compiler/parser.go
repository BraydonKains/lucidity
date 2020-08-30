/*
parser.go
Author: @BraydonKains

I have no idea what I'm doing
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

func (p *Parser) parseAssignment() ([]Token, error) {
	tokens := make([]Token, 0)
	for {
		token := p.readToken()

	}
	return nil, nil
}

func ParseValidity(tokens []Token) bool {
	for _, token := range tokens {
		fmt.Println(token)
	}

	p := newParser(tokens)

	for {
		token := p.peekToken()
		if token.IsType() {
			_, err := p.parseAssignment()
			if err != nil {
				fmt.Println(err)
				return false
			}
		}

		return true
	}
}
