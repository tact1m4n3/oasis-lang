package token

type TokenType int

type Token struct {
	Type TokenType
	Lit  string
}

const (
	_ TokenType = iota

	ILLEGAL
	EOF

	IDENT
	INT

	ADD
	SUB
	MUL
	DIV

	ASSIGN
	LT
	GT
	NOT
	EQ
	NEQ

	COMMA
	SEMI

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	LET
	IF
	ELSE
	FOR
	FUNC
	RETURN
)

var TokenName = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT: "IDENT",
	INT:   "INT",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
	LT:     "<",
	GT:     ">",
	NOT:    "!",
	EQ:     "==",
	NEQ:    "!=",

	COMMA: ",",
	SEMI:  ";",

	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	LET:    "let",
	IF:     "if",
	ELSE:   "else",
	FOR:    "for",
	FUNC:   "fn",
	RETURN: "return",
}

func (tt TokenType) String() string {
	return TokenName[tt]
}

var keywords = map[string]TokenType{
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"fn":     FUNC,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
