package try

import (
	"../compiler"
)

func TryParse(tokens []compiler.Token) bool {
	return compiler.ParseValidity(tokens)
}
