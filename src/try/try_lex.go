package try

import (
	"../compiler"
	"bufio"
	"os"
)

func TryLex() []compiler.Token {
	test_file := "./src/try/lex.lu"
	file, err := os.Open(test_file)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	l := compiler.NewLexer(test_file, reader)

	tokens := make([]compiler.Token, 0)

	for {
		token := l.Lex()
		tokens = append(tokens, token)
		if token.Is(compiler.EOF) {
			return tokens
		}
	}
}
