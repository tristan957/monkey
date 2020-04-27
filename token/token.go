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
	// MINUS represents the minus operator
	MINUS = "-"
	// BANG represents the bang operator
	BANG = "!"
	// ASTERISK represents the asterisk operator
	ASTERISK = "*"
	// FORWARD_SLASH represents the forward slash operator
	FORWARD_SLASH = "/"
	// LESS_THAN represents the less than operator
	LESS_THAN = "<"
	// GREATER_THAN represents the greater than operator
	GREATER_THAN = ">"
	// COMMA represents the ',' delimiter
	COMMA = ","
	// SEMICOLON represents the ';' delimiter
	SEMICOLON = ";"
	// LEFT_PARENTHESES represents an opening parentheses
	LEFT_PARENTHESES = "("
	// RIGHT_PARENTHESES represents a closing parentheses
	RIGHT_PARENTHESES = ")"
	// LEFT_BRACE represents an opening brace
	LEFT_BRACE = "{"
	// RIGHT_BRACE represents a closing brace
	RIGHT_BRACE = "{"
	// FUNCTION represents the 'fn' keyword
	FUNCTION = "FUNCTION"
	// LET represents the 'let' keywork
	LET = "LET"
	// TRUE represents the 'true' keyword
	TRUE = "true"
	// FALSE represents the 'false' keyword
	FALSE = "false"
	// IF represents the 'if' keyword
	IF = "if"
	// ELSE represents the 'else' keyword
	ELSE = "else"
	// RETURN represents the 'return' keyword
	RETURN = "return"
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
