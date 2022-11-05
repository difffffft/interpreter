package token

type Token struct {
	Type  TokenType
	Value string
}

var KeyWords = map[TokenType][]string{
	NULL:     {"null"},
	IMPORT:   {"import"},
	CONST:    {"const"},
	LET:      {"let"},
	ENUM:     {"enum"},
	FUNCTION: {"function"},
	RETURN:   {"return"},
	TRUE:     {"true"},
	FALSE:    {"false"},
	IF:       {"if"},
	ELSE:     {"else"},
	SWITCH:   {"switch"},
	CASE:     {"case"},
	DEFAULT:  {"default"},
	TRY:      {"try"},
	CATCH:    {"catch"},
	THROW:    {"throw"},
	FOR:      {"for"},
	IN:       {"in"},
	WHILE:    {"while"},
	DO:       {"do"},
	BREAK:    {"break"},
	CONTINUE: {"continue"},
	CLASS:    {"class"},
	NEW:      {"new"},
	STATIC:   {"static"},
	THIS:     {"this"},
	SUPER:    {"super"},
}

func IsKeyWords(s string) TokenType {
	//判断是否是关键字
	for k, v := range KeyWords {
		for _, j := range v {
			if j == s {
				return k
			}
		}
	}
	//不是关键字就是变量
	return VAR
}
