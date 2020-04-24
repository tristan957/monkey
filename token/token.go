package token

const (
	// IDENTIFIER represents variable/function names
	IDENTIFIER = "IDENTIFIER"
	// INTEGER represents integer constants
	INTEGER = "INTEGER"
	// ASSIGN represents the assignment operator
	ASSIGN = "="
	// PLUS represents the addition operator
	PLUS = "+"
	// COMMA represents the ',' delimiter
	COMMA = ","
	// SEMICOLON represents the ';' delimiter
	SEMICOLON = ";"
	// LPAREN represents an opening parentheses
	LPAREN = "("
	// RPAREN represents a closing parentheses
	RPAREN = ")"
	// LBRACE represents an opening brace
	LBRACE = "{"
	// RBRACE represents a closing brace
	RBRACE = "{"
	// FUNCTION represents the 'fn' keyword
	FUNCTION = "FUNCTION"
	// LET represents the 'let' keywork
	LET = "LET"
	// UNKNOWN represents an unknown token
	UNKNOWN = "UNKNOWN"
	// EOF represents the end of file
	EOF = "EOF"
)

// Type is the type of the token
type Type string

// Token is the type of the token and its literal representation
type Token struct {
	Type    Type
	Literal string
	Span    Span
}
