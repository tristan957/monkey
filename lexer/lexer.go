package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"git.sr.ht/~tristan957/monkey/token"
)

// Lexer keeps track of reading an input
type Lexer struct {
	input        *bufio.Reader
	currPosition *token.Position
	nextPosition *token.Position
	ch           byte
}

// NewFromString creates a new Lexer when given a string input.
func NewFromString(input string) *Lexer {
	reader := bufio.NewReader(strings.NewReader(input))
	l := &Lexer{
		input: reader,
		currPosition: &token.Position{
			Line:   1,
			Column: 1,
		},
		nextPosition: &token.Position{
			Line:   1,
			Column: 1,
		},
	}

	return l
}

// Initialize puts the Lexer into a fully working state and will return an
// error if it cannot be initialized properly.
func (l *Lexer) Initialize() error {
	if err := l.readChar(); err != nil {
		return fmt.Errorf("Unable to initialize the lexer: %w", err)
	}

	return nil
}

// NextToken returns the next token of the sequence
func (l *Lexer) NextToken() (token.Token, error) {
	var tok token.Token

	if err := l.consumeWhitespace(); err != nil {
		return tok, nil
	}

	switch l.ch {
	case '=':
		tok = l.newToken(token.ASSIGN)
	case '+':
		tok = l.newToken(token.PLUS)
	case '-':
		tok = l.newToken(token.MINUS)
	case '*':
		tok = l.newToken(token.ASTERISK)
	case '/':
		tok = l.newToken(token.FORWARD_SLASH)
	case '!':
		tok = l.newToken(token.BANG)
	case '<':
		tok = l.newToken(token.LESS_THAN)
	case '>':
		tok = l.newToken(token.GREATER_THAN)
	case ';':
		tok = l.newToken(token.SEMICOLON)
	case '(':
		tok = l.newToken(token.LEFT_PARENTHESES)
	case ')':
		tok = l.newToken(token.RIGHT_PARENTHESES)
	case ',':
		tok = l.newToken(token.COMMA)
	case '{':
		tok = l.newToken(token.LEFT_BRACE)
	case '}':
		tok = l.newToken(token.RIGHT_BRACE)
	case 0:
		currPositionCopy := l.currPosition.Copy()
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Span.Start = currPositionCopy
		tok.Span.End = currPositionCopy
	default:
		if isLetter(l.ch) {
			tok.Span.Start = l.currPosition.Copy()
			identifier, err := l.readIdentifier()
			if err != nil {
				return tok, err
			}

			tok.Literal = identifier
			tok.Type = token.LookupIdentifier(identifier)
			tok.Span.End = l.currPosition.Copy()
			// Since we read until we find a non-letter, the indetifier actually ends one
			// column before.
			tok.Span.End.Column--

			return tok, nil
		} else if isDigit(l.ch) {
			tok.Span.Start = l.currPosition.Copy()
			number, err := l.readInteger()
			if err != nil {
				return tok, err
			}

			tok.Literal = number
			tok.Type = token.INTEGER
			tok.Span.End = l.currPosition.Copy()
			// Since we read until we find a non-digit, the indetifier actually ends one
			// column before.
			tok.Span.End.Column--

			return tok, nil
		} else {
			tok = l.newToken(token.UNKNOWN)
		}
	}

	if err := l.readChar(); err != nil {
		return tok, err
	}

	return tok, nil
}

// newToken creates a new token based on the state of the lexer
func (l *Lexer) newToken(ttype token.Type) token.Token {
	currPositionCopy := l.currPosition.Copy()
	return token.Token{Type: ttype, Literal: string(l.ch), Span: token.Span{Start: currPositionCopy, End: currPositionCopy}}
}

// readChar reads a single character of the input
func (l *Lexer) readChar() error {
	_, err := l.input.Peek(1)
	if err != nil {
		if err == io.EOF {
			l.ch = 0
			l.currPosition.Line = l.nextPosition.Line
			l.currPosition.Column = l.nextPosition.Column
		} else {
			l.ch = 0
			return fmt.Errorf("Failed to peek character at line %d, column %d: %w", l.nextPosition.Line, l.nextPosition.Column, err)
		}
	} else {
		ch, err := l.input.ReadByte()
		if err != nil {
			return fmt.Errorf("Failed to read character at line %d, column %d: %w", l.nextPosition.Line, l.nextPosition.Column, err)
		}

		l.ch = ch
		l.currPosition.Line = l.nextPosition.Line
		l.currPosition.Column = l.nextPosition.Column
		if l.ch == byte('\n') {
			l.nextPosition.Line++
			l.nextPosition.Column = 1
		} else {
			l.nextPosition.Column++
		}
	}

	return nil
}

// isLetter checks if the input is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if the input is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readInteger reads an identifier name
func (l *Lexer) readIdentifier() (string, error) {
	var builder strings.Builder
	for isLetter(l.ch) {
		builder.WriteByte(l.ch)
		if err := l.readChar(); err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

// readInteger reads an integer
func (l *Lexer) readInteger() (string, error) {
	var builder strings.Builder
	for isDigit(l.ch) {
		builder.WriteByte(l.ch)
		if err := l.readChar(); err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

// consumeWhitespace eats all whitespace characters between tokens
func (l *Lexer) consumeWhitespace() error {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if err := l.readChar(); err != nil {
			return err
		}
	}

	return nil
}
