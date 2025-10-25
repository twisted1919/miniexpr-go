package miniexpr

import (
	"errors"
	"fmt"
)

var (
	// ErrUnknownExpression is returned for unknown expressions.
	ErrUnknownExpression = errors.New("unknown expression")
	// ErrUnknownOperator is returned for unknown operators.
	ErrUnknownOperator = errors.New("unknown operator")
	// ErrDivByZero is returned for division by zero.
	ErrDivByZero = errors.New("division by zero")
)

// UnexpectedTokenError holds error information when we encounter an unexpected token.
type UnexpectedTokenError struct {
	token byte
	pos   int
}

// NewUnexpectedTokenError creates a new UnexpectedTokenError.
func NewUnexpectedTokenError(token byte, pos int) UnexpectedTokenError {
	return UnexpectedTokenError{token: token, pos: pos}
}

// Error implements the error interface.
func (e UnexpectedTokenError) Error() string {
	return fmt.Sprintf("syntax error, unexpected token %s at position %d", string(e.token), e.pos)
}

// Is checks if a given error is actually a UnexpectedTokenError.
func (e UnexpectedTokenError) Is(err error) bool {
	var errPtr *UnexpectedTokenError

	return errors.As(err, &errPtr)
}
