package try

import (
	"../lexer"
	"bufio"
	"fmt"
	"os"
)

func TryLex() {
	test_file := "./src/try/lex.lu"
	file, err := os.Open(test_file)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	l := lexer.NewLexer(reader, test_file)

	for {
		token := l.Lex()
		fmt.Println(token)
		if token.Is(lexer.EOF) {
			return
		}
	}
}
