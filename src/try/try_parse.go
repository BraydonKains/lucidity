package try

import (
	"../lexer"
	"../parser"
)

func TryParse(tokens []lexer.Token) bool {
	return parser.ParseValidity(tokens)
}
