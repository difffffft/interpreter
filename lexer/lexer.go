package lexer

import (
	"e/token"
	"fmt"
	"testing"
)

type Lexer struct {
	//源代码
	input string

	//分析器当前所处位置
	position int

	//分析器当前正在被读取的位置
	readPosition int

	//正在被读取的字符
	ch byte
}

//初始化
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

//读取字符
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	fmt.Println("当前位置:", l.position)
	fmt.Println("当前位置的字符:", string(l.ch))
	fmt.Println("将要读取的位置:", l.readPosition)
}

//返回当前字符对应的Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
	case '+':
		tok = token.Token{Type: token.ADD, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case '0':
		tok = token.Token{Type: token.EOF, Literal: ""}
	}
	l.ReadChar()
	return tok
}

func TestNextToken(t *testing.T) {

}
