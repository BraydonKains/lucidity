package try

import (
	"../lexer"
	"bufio"
	// "fmt"
	"os"
)

func TryLex() []lexer.Token {
	test_file := "./src/try/lex.lu"
	file, err := os.Open(test_file)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	l := lexer.NewLexer(test_file, reader)

	tokens := make([]lexer.Token, 0)

	for {
		token := l.Lex()
		// fmt.Println(token)
		tokens = append(tokens, token)
		if token.Is(lexer.EOF) {
			return tokens
		}
	}
}
