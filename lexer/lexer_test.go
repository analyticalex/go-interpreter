package lexer

import (
	"testing"
	"github.com/analyticalex/go-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+,;(){}[]" // test string
	
	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral	string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
	}
	
	l := New(input) // Creates new Lexer
	
	for i, tt := range tests {
		tok := l.NextToken()
		
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]: TokenType wrong. expected = %q, actual = %q", i, tt.expectedType, tok.Type)
		}
		
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]: Literal wrong. expected = %q, actual = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}