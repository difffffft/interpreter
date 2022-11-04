package main

import (
	"e/lexer"
)

func main() {
	input := `
		=+abc_(function)
		{let}$$,;12345.67.
`
	lexer.New(input)
}
