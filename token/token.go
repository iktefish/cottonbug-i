// Lex the following theoretical subset of Cottonbug
/*
let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
*/

package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	INDENT = "INDENT"
	INT    = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA      = ","
	SEMICOLON  = ";"
	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
