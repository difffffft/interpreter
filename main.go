package main

import (
	"e/lexer"
)

func main() {
	input := `=+(){},;`
	l := lexer.New(input)

	//读取第一个token
	l.NextToken()
	l.NextToken()
	l.NextToken()
	l.NextToken()
	l.NextToken()
	l.NextToken()
	l.NextToken()
	l.NextToken()
	//fmt.Println(t)
}
