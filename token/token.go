/*
Package tokens describes tokens recognized by lexer,
and implements utility functions.
*/

package token

type TokenType string

// TO DO: Attach filename and line number for better error reporting
type Token struct {
	Type	TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"
	
	// Identifiers and literals
	IDENT 	= "IDENT"	// foobar, x, y, ...
	INT		= "INT"		// 123456...
	
	// Operators
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG		= "!"
	ASTERISK	= "*"
	SLASH		= "/"
	
	LT	= "<"
	GT = ">"
	
	EQ		= "=="
	NOTEQ	= "!="
	
	//	Delimiters
	COMMA		= ","
	SEMICOLON	= ";"
	
	LPAREN	= "("
	RPAREN	= ")"
	LBRACE	= "{"
	RBRACE	= "}"
	LBRACK	= "["
	RBRACK	= "]"
	
	// Keywords
	FUNCTION	= "FN"
	LET			= "LET"
	TRUE		= "TRUE"
	FALSE		= "FALSE"
	IF			= "IF"
	ELSE		= "ELSE"
	RETURN		= "RETURN"
	
	// Graph Keywords
	GRAPH		= "G"
	VERTEX		= "VERTEX"
	ALLVERTS	= "V"
	EDGE		= "EDGE"
	ALLEDGES	= "E"
)

// TO DO: Implement graph utility by recognizing keywords
var keywords = map[string]TokenType{
	"fn":		FUNCTION,
	"let":		LET,
	"true":		TRUE,
	"false":	FALSE,
	"if":		IF,
	"else":		ELSE,
	"return":	RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // return keyword's TokenType cosntant
	}
	return IDENT // return TokenType for user-defined identifiers
}