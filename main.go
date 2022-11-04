package main

import (
	"e/lexer"
	"io/ioutil"
)

func main() {
	codeStr := ReadFile("./test/test.js")
	lexer.New(codeStr)
}

func ReadFile(fileName string) string {
	f, _ := ioutil.ReadFile(fileName)
	return string(f)
}
