package token

// ILLEGAL 未知的词法单元
const ILLEGAL = "ILLEGAL"

// EOF 文件结束符
const EOF = "EOF"

// VAR 标识符
const VAR = "VAR"

// INT 字面量
const INT = "INT"
const FLOAT = "FLOAT"

// ASSIGN ADD 运算符
const ASSIGN = "="
const ADD = "+"

// COMMA SEMICOLON 分隔符
const COMMA = ","
const SEMICOLON = ";"

// LPAREN RPAREN LBRACE RBRACE 分隔符
const LPAREN = "("
const RPAREN = ")"
const LBRACE = "{"
const RBRACE = "}"

const FUNCTION = "FUNCTION"
const LET = "LET"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

//关键字
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
}

// IsKeyWords 判断是否是关键字还是变量
func IsKeyWords(s string) TokenType {
	if tok, ok := keywords[s]; ok {
		return tok
	}
	return VAR
}
