package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifer + literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 12345

	// operator
	ASSIGN = "="
	PLUS = "+"

	// delimitor
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET = "LET"
)