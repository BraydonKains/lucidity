package main

import (
	"./src/try"
	"fmt"
)

func main() {
	tokens := try.TryLex()
	fmt.Println(try.TryParse(tokens))
}
