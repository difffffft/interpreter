package ast

import (
	"e/token"
)

// Node 每个节点都需要实现Node接口
type Node interface {
	TokenValue() string
	//String() string
}

// Statement 语句
// 每个程序都有很多的语句
// 每个语句都有一个起始节点
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
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

// VarStatement 是一个表达式
// VarStatement 实现了Expression接口
type VarStatement struct {
	Token token.Token
	Value string
}

func (i *VarStatement) statementNode() {
}
func (i *VarStatement) expressionNode() {
}
func (i *VarStatement) TokenValue() string {
	return i.Token.Value
}

// LetStatement let词法单元
type LetStatement struct {
	Token token.Token
	//标识符
	Name *VarStatement
	//表达式
	Value Expression
}

//func (s *LetStatement) String() string {
//	return fmt.Sprintf(")
//}

func (l *LetStatement) statementNode() {
}
func (l *LetStatement) TokenValue() string {
	return l.Token.Value
}

// ReturnStatement return词法单元
type ReturnStatement struct {
	//每一个词法单元都包含一个关键字
	Token token.Token
	//表达式
	Value Expression
}

func (rs *ReturnStatement) statementNode() {
}
func (rs *ReturnStatement) TokenValue() string {
	return rs.Token.Value
}

// ExpressionStatement 表达式词法单元
type ExpressionStatement struct {
	//每一个词法单元都包含一个关键字
	Token token.Token
	//表达式
	Value Expression
}

func (es *ExpressionStatement) statementNode() {
}
func (es *ExpressionStatement) TokenValue() string {
	return es.Token.Value
}

// IntegerStatement 是一个表达式
// IntegerStatement 实现了Expression接口
type IntegerStatement struct {
	Token token.Token
	Value int64
}

func (es *IntegerStatement) statementNode() {
}
func (es *IntegerStatement) expressionNode() {
}
func (es *IntegerStatement) TokenValue() string {
	return es.Token.Value
}
