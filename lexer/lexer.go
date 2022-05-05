/* A Lexer will take source code as input and ouput tokens */

package lexer

import (
	// "testing/"
	"iktefish/_s/tree/main/go/cottonbug-i/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.EOF, "EOF"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFTPAREN, "("},
		{token.RIGHTPAREN, ")"},
		{token.LEFTBRACE, "{"},
		{token.RIGHTBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatal("tests[%d] - Token type wrong! Expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatal("tests[%d] - Literal wrong! Expected=%q, got= %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
