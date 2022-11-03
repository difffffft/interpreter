package main

import (
	"e/lexer"
)

func main() {
	input := `=+abc_(){},;`
	lexer.New(input)
}
