package ast

import (
	"e/token"
)

// Node 每个节点都需要实现Node接口
type Node interface {

	// TokenValue 返回词法单元的字面量
	TokenValue() string
}

// Statement 语句
// 每个程序都有很多的语句
// 每个语句都有一个起始节点
type Statement interface {
	Node
	statementNode()
}

type Experssion interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) statementNode() {}

func (p *Program) TokenValue() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	}
	return ""
}

// Identifier 标识符词法单元
type IdentifierStatement struct {
	Token token.Token
	Value string
}

func (i *IdentifierStatement) statementNode() {
}
func (i *IdentifierStatement) TokenValue() string {
	return i.Token.Value
}

// LetStatement let词法单元
type LetStatement struct {
	Token token.Token
	//标识符
	Name *IdentifierStatement
	//表达式
	Value Experssion
}

//func (s *LetStatement) String() string {
//	return fmt.Sprintf(")
//}

func (l *LetStatement) statementNode() {
}
func (l *LetStatement) TokenValue() string {
	return l.Token.Value
}
