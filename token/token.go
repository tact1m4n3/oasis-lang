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
	PLUS
	MINUS
	ASTERISK
	SLASH
	MOD

	AND
	OR
	XOR
	LSHIFT
	RSHIFT
	NOT

	LAND
	LOR
	BANG

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
	RETURN
	IF
	ELSE
	FUNC
	WHILE
)

var TokenName = map[Token]string{
	ILLEGAL:    "ILLEGAL",
	UNEXPECTED: "UNEXPECTED",
	EOF:        "EOF",

	IDENT: "IDENT",
	INT:   "INT",

	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	ASTERISK: "*",
	SLASH:    "/",
	MOD:      "%",

	AND:    "&",
	OR:     "|",
	XOR:    "^",
	LSHIFT: "<<",
	RSHIFT: ">>",
	NOT:    "~",

	LAND: "&&",
	LOR:  "||",
	BANG: "!",

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

	LET:    "let",
	RETURN: "return",
	IF:     "if",
	ELSE:   "else",
	FUNC:   "func",
	WHILE:  "while",
}

func (tok Token) String() string {
	return TokenName[tok]
}

var keywords = map[string]Token{
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"func":   FUNC,
	"while":  WHILE,
}

func LookupIdent(ident string) Token {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
