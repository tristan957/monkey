package token

import "fmt"

// Span represents the area a token takes up. Start and End are pointers
// because we can make a slight optimization where Start and End point to the
// same underlying object in the case of 1 charater tokens.
type Span struct {
	Start *Position
	End   *Position
}

// NewSpan creates a Span
func NewSpan(startLine, startColumn, endLine, endColumn int) *Span {
	startPosition := &Position{
		Line:   startLine,
		Column: startColumn,
	}

	if startLine == endLine && startColumn == endColumn {
		return &Span{
			Start: startPosition,
			End:   startPosition,
		}
	}

	return &Span{
		Start: &Position{
			Line:   startLine,
			Column: startColumn,
		},
		End: &Position{
			Line:   endLine,
			Column: endColumn,
		},
	}
}

// NewSpanFromSinglularToken creates a Span representing a token with a width
// of 1. This function is a shortcut if you know the token type beforehand.
func NewSpanFromSinglularToken(line, column int) *Span {
	position := &Position{
		Line:   line,
		Column: column,
	}

	return &Span{
		Start: position,
		End:   position,
	}
}

// Equals returns whether two Spans span the same area.
func (s *Span) Equals(other *Span) bool {
	if s == other {
		return true
	}

	return s.Start.Equals(other.Start) && s.End.Equals(other.End)
}

// String returns a string representation of a Span
func (s *Span) String() string {
	return fmt.Sprintf("{%s -> %s}", s.Start, s.End)
}
