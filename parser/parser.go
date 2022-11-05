package parser

import (
	"e/ast"
	"e/lexer"
	"e/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string
}

// ParseProgram 开始分析程序,得到语法分析树
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {

			//语法分析过程中如果遇到语法分析错误,就暴露给用户
			//并停止下面的语法分析
			if len(p.errors) > 0 {
				fmt.Println(p.errors)
				return program
			} else {
				fmt.Println(stmt)
			}

			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

//语法分析遇到 期望token错误
//期望token错误表示，你的下一个token和上一个token不配套使用时
func (p *Parser) addExpectError(t string) {
	msg := fmt.Sprintf("语法出现错误\n期望得到的类型: %s\n而程序的类型:%s\n位置:%d\n值:%s", t, p.peekToken.Type, p.l.GetPosition(), p.peekToken.Value)
	p.errors = append(p.errors, msg)
}

//记录当前token
//并找到下一个token
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.ParseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) currTokenIs(t string) bool {
	return p.currToken.Type == t
}

//判断下一个token和自己的预期是否想符合
func (p *Parser) peekTokenIs(t string) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addExpectError(t)
		return false
	}
}
func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}
	if !p.expectPeek(token.VAR) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Value}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	if !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	return p
}
