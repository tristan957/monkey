package token

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier turns keyords into a Type
func LookupIdentifier(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}
