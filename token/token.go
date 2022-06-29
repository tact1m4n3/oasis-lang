package token

type Token int

const (
	_ Token = iota

	ILLEGAL
	UNEXPECTED
	EOF

	IDENT
	INT

	ASSIGN
	ADD
	SUB
	MUL
	DIV
	MOD

	AND
	OR
	XOR
	LSHIFT
	RSHIFT
	TILDE

	ADD_ASSIGN
	SUB_ASSIGN
	MUL_ASSIGN
	DIV_ASSIGN
	MOD_ASSIGN

	AND_ASSIGN
	OR_ASSIGN
	XOR_ASSIGN
	LSHIFT_ASSIGN
	RSHIFT_ASSIGN

	LAND
	LOR
	NOT

	EQ
	NEQ
	LT
	LTE
	GT
	GTE

	COMMA
	SEMI

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	LET
	IF
	ELSE
	WHILE
	CONTINUE
	BREAK
	FUNC
	RETURN
)

var TokenName = map[Token]string{
	ILLEGAL:    "ILLEGAL",
	UNEXPECTED: "UNEXPECTED",
	EOF:        "EOF",

	IDENT: "IDENT",
	INT:   "INT",

	ASSIGN: "=",
	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	DIV:    "/",
	MOD:    "%",

	AND:    "&",
	OR:     "|",
	XOR:    "^",
	LSHIFT: "<<",
	RSHIFT: ">>",
	TILDE:  "~",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	DIV_ASSIGN: "/=",
	MOD_ASSIGN: "%=",

	AND_ASSIGN:    "&=",
	OR_ASSIGN:     "|=",
	XOR_ASSIGN:    "^=",
	LSHIFT_ASSIGN: "<<=",
	RSHIFT_ASSIGN: ">>=",

	LAND: "&&",
	LOR:  "||",
	NOT:  "!",

	EQ:  "==",
	NEQ: "!=",
	LT:  "<",
	LTE: "<=",
	GT:  ">",
	GTE: ">=",

	COMMA: ",",
	SEMI:  ";",

	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	LET:      "let",
	IF:       "if",
	ELSE:     "else",
	WHILE:    "while",
	CONTINUE: "continue",
	BREAK:    "break",
	FUNC:     "func",
	RETURN:   "return",
}

func (tok Token) String() string {
	return TokenName[tok]
}

var keywords = map[string]Token{
	"let":      LET,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"continue": CONTINUE,
	"break":    BREAK,
	"func":     FUNC,
	"return":   RETURN,
}

func LookupIdent(ident string) Token {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
