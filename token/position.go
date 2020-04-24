package token

import "fmt"

// Position represents the point in a file or string where a token is
type Position struct {
	Line   int
	Column int
}

// Equals returns whether two positions are the same
func (p *Position) Equals(other *Position) bool {
	if p == other {
		return true
	}

	return p.Line == other.Line && p.Column == other.Column
}

// String returns a string representation of a Position
func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Line, p.Column)
}

// Copy returns a new Position with the same values.
func (p Position) Copy() *Position {
	return &Position{
		Line:   p.Line,
		Column: p.Column,
	}
}
