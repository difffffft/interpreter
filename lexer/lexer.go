package lexer

import (
	"e/token"
	"fmt"
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

// New 初始化
func New(input string) *Lexer {
	l := &Lexer{input: input}

	//初始化游标
	l.ReadChar()

	//读取第一个token
	var _token = l.NextToken()
	//直到token结束
	for _token.Type != token.EOF {
		_token = l.NextToken()
		fmt.Println(_token)
	}

	return l
}

// ReadChar 读取字符
// 游标向下移动
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = '0'
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// IsLetter 判断是否是字母和下划线开头
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// IsDigit 判断是否是数字开头
func IsDigit(ch byte, pointNum *int) bool {
	//if pointNum != nil {
	//	if *pointNum > 1 {
	//		return false
	//	}
	//	if ch == '.' {
	//		*pointNum++
	//		return true
	//	}
	//}
	return '0' <= ch && ch <= '9'
}

// skipWhitespace 跳过空格和换行符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.ReadChar()
	}
}

// ReadVar 读取变量
func (l *Lexer) ReadVar() string {
	position := l.position
	//是字母就继续往下读
	for IsLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

// ReadNumber 读取int类型的数字
func (l *Lexer) ReadNumber() string {
	position := l.position
	pointNum := 0
	//是数字就继续往下读
	for IsDigit(l.ch, &pointNum) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

// NextToken 返回当前字符对应的Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	//空格和回车不是token类型,但是也不属于语法错误,所以读取的时候,需要跳过
	l.skipWhitespace()

	//读取正常的token
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
	default:
		if IsLetter(l.ch) {
			//判断是变量还是关键字
			tok.Literal = l.ReadVar()
			tok.Type = token.IsKeyWords(tok.Literal)
			return tok
		} else if IsDigit(l.ch, nil) {
			//判断是变量还是关键字
			tok.Literal = l.ReadNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = "未知语法"
		}
	}
	l.ReadChar()
	return tok
}
