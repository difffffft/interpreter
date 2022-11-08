package parser

import (
	"e/ast"
	"e/lexer"
	"e/token"
	"fmt"
	"strconv"
)

//表达式前缀解析函数
type prefixParseFn func() ast.Expression

//表达式中缀解析函数
type midfixParseFn func(es ast.Expression) ast.Expression

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	midfixParseFns map[token.TokenType]midfixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerMidfix(tokenType token.TokenType, fn midfixParseFn) {
	p.midfixParseFns[tokenType] = fn
}

// New 入口函数
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.VAR, p.ParseVar)
	p.registerPrefix(token.INT, p.ParseInt)

	p.ParseProgram()
	return p
}

// ParseProgram 开始分析程序,得到语法分析树
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
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
func (p *Parser) addExpectError(t token.TokenType) {
	msg := fmt.Sprintf("语法出现错误\n期望得到的类型: %s\n而程序的类型:%s\n位置:%d\n值:%s", t, p.peekToken.Type, p.l.GetPosition(), p.peekToken.Value)
	p.errors = append(p.errors, msg)
}

//记录当前token
//并找到下一个token
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//每分析完一个语句,就开始分析下一个语句
//分析语句
func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	case token.VAR:
		return p.ParseExpressionStatement()
	case token.INT:
		return p.ParseIntStatement()
	default:
		return nil
		//return p.ParseExpressionStatement()
	}
}

//期望遇到某个token
func (p *Parser) expectCurrTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

//下一个token,期望遇到某个token
func (p *Parser) expectNextTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectNextToken(t token.TokenType) bool {
	if p.expectNextTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.addExpectError(t)
		return false
	}
}

func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}
	if !p.expectNextToken(token.VAR) {
		return nil
	}
	stmt.Name = &ast.VarStatement{Token: p.currToken, Value: p.currToken.Value}
	if !p.expectNextToken(token.ASSIGN) {
		return nil
	}
	if !p.expectCurrTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	//直到return语句遇到了;号
	if !p.expectCurrTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) ParseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) ParseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	//拿到表达式的值
	stmt.Value = p.ParseExpression(LOWEST)
	//如果下一个是分号,就继续往后移一下，方便下一个程序调用
	if p.expectCurrTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
func (p *Parser) ParseIntStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}
	//拿到表达式的值
	stmt.Value = p.ParseExpression(LOWEST)
	//如果下一个是分号,就继续往后移一下，方便下一个程序调用
	if p.expectCurrTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) ParseVar() ast.Expression {
	return &ast.VarStatement{Token: p.currToken, Value: p.currToken.Value}
}

func (p *Parser) ParseInt() ast.Expression {
	stmt := &ast.IntegerStatement{Token: p.currToken}
	value, err := strconv.ParseInt(p.currToken.Value, 0, 64)
	if err != nil {
		p.addExpectError(p.currToken.Type)
		return nil
	}
	stmt.Value = value
	return stmt
}
