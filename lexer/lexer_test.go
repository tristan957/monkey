package lexer

import (
	"testing"

	"git.sr.ht/~tristan957/monkey/token"
)

func TestNextToken_OneLineString(t *testing.T) {
	input := "=+-*/(){},;<>!"

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
		expectedSpan    token.Span
	}{
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 1, Column: 1}, End: &token.Position{Line: 1, Column: 1}}},
		{token.PLUS, "+", token.Span{Start: &token.Position{Line: 1, Column: 2}, End: &token.Position{Line: 1, Column: 2}}},
		{token.MINUS, "-", token.Span{Start: &token.Position{Line: 1, Column: 3}, End: &token.Position{Line: 1, Column: 3}}},
		{token.ASTERISK, "*", token.Span{Start: &token.Position{Line: 1, Column: 4}, End: &token.Position{Line: 1, Column: 4}}},
		{token.FORWARD_SLASH, "/", token.Span{Start: &token.Position{Line: 1, Column: 5}, End: &token.Position{Line: 1, Column: 5}}},
		{token.LEFT_PARENTHESES, "(", token.Span{Start: &token.Position{Line: 1, Column: 6}, End: &token.Position{Line: 1, Column: 6}}},
		{token.RIGHT_PARENTHESES, ")", token.Span{Start: &token.Position{Line: 1, Column: 7}, End: &token.Position{Line: 1, Column: 7}}},
		{token.LEFT_BRACE, "{", token.Span{Start: &token.Position{Line: 1, Column: 8}, End: &token.Position{Line: 1, Column: 8}}},
		{token.RIGHT_BRACE, "}", token.Span{Start: &token.Position{Line: 1, Column: 9}, End: &token.Position{Line: 1, Column: 9}}},
		{token.COMMA, ",", token.Span{Start: &token.Position{Line: 1, Column: 10}, End: &token.Position{Line: 1, Column: 10}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 1, Column: 11}, End: &token.Position{Line: 1, Column: 11}}},
		{token.LESS_THAN, "<", token.Span{Start: &token.Position{Line: 1, Column: 12}, End: &token.Position{Line: 1, Column: 12}}},
		{token.GREATER_THAN, ">", token.Span{Start: &token.Position{Line: 1, Column: 13}, End: &token.Position{Line: 1, Column: 13}}},
		{token.BANG, "!", token.Span{Start: &token.Position{Line: 1, Column: 14}, End: &token.Position{Line: 1, Column: 14}}},
		{token.EOF, "", token.Span{Start: &token.Position{Line: 1, Column: 15}, End: &token.Position{Line: 1, Column: 15}}},
	}

	l := NewFromString(input)
	if err := l.Initialize(); err != nil {
		t.Fatal(err)
	}

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatal(err)
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, actual=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, actual=%q", i, tt.expectedLiteral, tok.Literal)
		}

		if !tok.Span.Equals(&tt.expectedSpan) {
			t.Fatalf("tests[%d] - wrong span, expected=%s, actual=%s", i, tt.expectedSpan, tok.Span)
		}
	}
}

func TestNextToken_ProgramString(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);

if (5 < 10) {
	return true;
} else {
	return false;
}

let b = 0b010101;
let o = 0o123456;
let h = 0xabF127;
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
		expectedSpan    token.Span
	}{
		{token.LET, "let", token.Span{Start: &token.Position{Line: 1, Column: 1}, End: &token.Position{Line: 1, Column: 3}}},
		{token.IDENTIFIER, "five", token.Span{Start: &token.Position{Line: 1, Column: 5}, End: &token.Position{Line: 1, Column: 8}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 1, Column: 10}, End: &token.Position{Line: 1, Column: 10}}},
		{token.INTEGER, "5", token.Span{Start: &token.Position{Line: 1, Column: 12}, End: &token.Position{Line: 1, Column: 12}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 1, Column: 13}, End: &token.Position{Line: 1, Column: 13}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 2, Column: 1}, End: &token.Position{Line: 2, Column: 3}}},
		{token.IDENTIFIER, "ten", token.Span{Start: &token.Position{Line: 2, Column: 5}, End: &token.Position{Line: 2, Column: 7}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 2, Column: 9}, End: &token.Position{Line: 2, Column: 9}}},
		{token.INTEGER, "10", token.Span{Start: &token.Position{Line: 2, Column: 11}, End: &token.Position{Line: 2, Column: 12}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 2, Column: 13}, End: &token.Position{Line: 2, Column: 13}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 4, Column: 1}, End: &token.Position{Line: 4, Column: 3}}},
		{token.IDENTIFIER, "add", token.Span{Start: &token.Position{Line: 4, Column: 5}, End: &token.Position{Line: 4, Column: 7}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 4, Column: 9}, End: &token.Position{Line: 4, Column: 9}}},
		{token.FUNCTION, "fn", token.Span{Start: &token.Position{Line: 4, Column: 11}, End: &token.Position{Line: 4, Column: 12}}},
		{token.LEFT_PARENTHESES, "(", token.Span{Start: &token.Position{Line: 4, Column: 13}, End: &token.Position{Line: 4, Column: 13}}},
		{token.IDENTIFIER, "x", token.Span{Start: &token.Position{Line: 4, Column: 14}, End: &token.Position{Line: 4, Column: 14}}},
		{token.COMMA, ",", token.Span{Start: &token.Position{Line: 4, Column: 15}, End: &token.Position{Line: 4, Column: 15}}},
		{token.IDENTIFIER, "y", token.Span{Start: &token.Position{Line: 4, Column: 17}, End: &token.Position{Line: 4, Column: 17}}},
		{token.RIGHT_PARENTHESES, ")", token.Span{Start: &token.Position{Line: 4, Column: 18}, End: &token.Position{Line: 4, Column: 18}}},
		{token.LEFT_BRACE, "{", token.Span{Start: &token.Position{Line: 4, Column: 20}, End: &token.Position{Line: 4, Column: 20}}},
		{token.IDENTIFIER, "x", token.Span{Start: &token.Position{Line: 5, Column: 2}, End: &token.Position{Line: 5, Column: 2}}},
		{token.PLUS, "+", token.Span{Start: &token.Position{Line: 5, Column: 4}, End: &token.Position{Line: 5, Column: 4}}},
		{token.IDENTIFIER, "y", token.Span{Start: &token.Position{Line: 5, Column: 6}, End: &token.Position{Line: 5, Column: 6}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 5, Column: 7}, End: &token.Position{Line: 5, Column: 7}}},
		{token.RIGHT_BRACE, "}", token.Span{Start: &token.Position{Line: 6, Column: 1}, End: &token.Position{Line: 6, Column: 1}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 6, Column: 2}, End: &token.Position{Line: 6, Column: 2}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 8, Column: 1}, End: &token.Position{Line: 8, Column: 3}}},
		{token.IDENTIFIER, "result", token.Span{Start: &token.Position{Line: 8, Column: 5}, End: &token.Position{Line: 8, Column: 10}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 8, Column: 12}, End: &token.Position{Line: 8, Column: 12}}},
		{token.IDENTIFIER, "add", token.Span{Start: &token.Position{Line: 8, Column: 14}, End: &token.Position{Line: 8, Column: 16}}},
		{token.LEFT_PARENTHESES, "(", token.Span{Start: &token.Position{Line: 8, Column: 17}, End: &token.Position{Line: 8, Column: 17}}},
		{token.IDENTIFIER, "five", token.Span{Start: &token.Position{Line: 8, Column: 18}, End: &token.Position{Line: 8, Column: 21}}},
		{token.COMMA, ",", token.Span{Start: &token.Position{Line: 8, Column: 22}, End: &token.Position{Line: 8, Column: 22}}},
		{token.IDENTIFIER, "ten", token.Span{Start: &token.Position{Line: 8, Column: 24}, End: &token.Position{Line: 8, Column: 26}}},
		{token.RIGHT_PARENTHESES, ")", token.Span{Start: &token.Position{Line: 8, Column: 27}, End: &token.Position{Line: 8, Column: 27}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 8, Column: 28}, End: &token.Position{Line: 8, Column: 28}}},
		{token.IF, "if", token.Span{Start: &token.Position{Line: 10, Column: 1}, End: &token.Position{Line: 10, Column: 2}}},
		{token.LEFT_PARENTHESES, "(", token.Span{Start: &token.Position{Line: 10, Column: 4}, End: &token.Position{Line: 10, Column: 4}}},
		{token.INTEGER, "5", token.Span{Start: &token.Position{Line: 10, Column: 5}, End: &token.Position{Line: 10, Column: 5}}},
		{token.LESS_THAN, "<", token.Span{Start: &token.Position{Line: 10, Column: 7}, End: &token.Position{Line: 10, Column: 7}}},
		{token.INTEGER, "10", token.Span{Start: &token.Position{Line: 10, Column: 9}, End: &token.Position{Line: 10, Column: 10}}},
		{token.RIGHT_PARENTHESES, ")", token.Span{Start: &token.Position{Line: 10, Column: 11}, End: &token.Position{Line: 10, Column: 11}}},
		{token.LEFT_BRACE, "{", token.Span{Start: &token.Position{Line: 10, Column: 13}, End: &token.Position{Line: 10, Column: 13}}},
		{token.RETURN, "return", token.Span{Start: &token.Position{Line: 11, Column: 2}, End: &token.Position{Line: 11, Column: 7}}},
		{token.TRUE, "true", token.Span{Start: &token.Position{Line: 11, Column: 9}, End: &token.Position{Line: 11, Column: 12}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 11, Column: 13}, End: &token.Position{Line: 11, Column: 13}}},
		{token.RIGHT_BRACE, "}", token.Span{Start: &token.Position{Line: 12, Column: 1}, End: &token.Position{Line: 12, Column: 1}}},
		{token.ELSE, "else", token.Span{Start: &token.Position{Line: 12, Column: 3}, End: &token.Position{Line: 12, Column: 6}}},
		{token.LEFT_BRACE, "{", token.Span{Start: &token.Position{Line: 12, Column: 8}, End: &token.Position{Line: 12, Column: 8}}},
		{token.RETURN, "return", token.Span{Start: &token.Position{Line: 13, Column: 2}, End: &token.Position{Line: 13, Column: 7}}},
		{token.FALSE, "false", token.Span{Start: &token.Position{Line: 13, Column: 9}, End: &token.Position{Line: 13, Column: 13}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 13, Column: 14}, End: &token.Position{Line: 13, Column: 14}}},
		{token.RIGHT_BRACE, "}", token.Span{Start: &token.Position{Line: 14, Column: 1}, End: &token.Position{Line: 14, Column: 1}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 16, Column: 1}, End: &token.Position{Line: 16, Column: 3}}},
		{token.IDENTIFIER, "b", token.Span{Start: &token.Position{Line: 16, Column: 5}, End: &token.Position{Line: 16, Column: 5}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 16, Column: 7}, End: &token.Position{Line: 16, Column: 7}}},
		{token.INTEGER, "0b010101", token.Span{Start: &token.Position{Line: 16, Column: 9}, End: &token.Position{Line: 16, Column: 16}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 16, Column: 17}, End: &token.Position{Line: 16, Column: 17}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 17, Column: 1}, End: &token.Position{Line: 17, Column: 3}}},
		{token.IDENTIFIER, "o", token.Span{Start: &token.Position{Line: 17, Column: 5}, End: &token.Position{Line: 17, Column: 5}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 17, Column: 7}, End: &token.Position{Line: 17, Column: 7}}},
		{token.INTEGER, "0o123456", token.Span{Start: &token.Position{Line: 17, Column: 9}, End: &token.Position{Line: 17, Column: 16}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 17, Column: 17}, End: &token.Position{Line: 17, Column: 17}}},
		{token.LET, "let", token.Span{Start: &token.Position{Line: 18, Column: 1}, End: &token.Position{Line: 18, Column: 3}}},
		{token.IDENTIFIER, "h", token.Span{Start: &token.Position{Line: 18, Column: 5}, End: &token.Position{Line: 18, Column: 5}}},
		{token.ASSIGN, "=", token.Span{Start: &token.Position{Line: 18, Column: 7}, End: &token.Position{Line: 18, Column: 7}}},
		{token.INTEGER, "0xabF127", token.Span{Start: &token.Position{Line: 18, Column: 9}, End: &token.Position{Line: 18, Column: 16}}},
		{token.SEMICOLON, ";", token.Span{Start: &token.Position{Line: 18, Column: 17}, End: &token.Position{Line: 18, Column: 17}}},
		{token.EOF, "", token.Span{Start: &token.Position{Line: 19, Column: 1}, End: &token.Position{Line: 19, Column: 1}}},
	}

	l := NewFromString(input)
	if err := l.Initialize(); err != nil {
		t.Fatal(err)
	}

	for i, tt := range tests {
		tok, err := l.NextToken()
		if err != nil {
			t.Fatal(err)
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong, expected=%q, actual=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, actual=%q", i, tt.expectedLiteral, tok.Literal)
		}

		if !tok.Span.Equals(&tt.expectedSpan) {
			t.Fatalf("tests[%d] - wrong span, expected=%s, actual=%s", i, tt.expectedSpan, tok.Span)
		}
	}
}
