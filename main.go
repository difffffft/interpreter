package main

import (
	"e/lexer"
	"e/parser"
	"io/ioutil"
)

func main() {
	codeStr := ReadFile("./test/test.js")
	l := lexer.New(codeStr)
	p := parser.New(l)
	p.ParseProgram()
}

func ReadFile(fileName string) string {
	f, _ := ioutil.ReadFile(fileName)
	return string(f)
}
