package parser

//表达式运算优先级
//等级越高越靠后

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < or >
	SUM         // +
	PRODUCT     // *
	PREFEIX     // -X or !X
	CALL        //fn()
)
