package try

import (
	"../lexer"
	"bufio"
	"fmt"
	"os"
)

func TryLex() {
	file, err := os.Open("./src/try/lex.lu")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	l := lexer.NewLexer(reader)

	for {
		pos, token, literal := l.Lex()
		fmt.Println(pos, token, literal)
		if token == lexer.EOF {
			return
		}
	}
}
