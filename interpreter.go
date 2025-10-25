package miniexpr

import (
	"fmt"
	"math"
)

// Interpreter handles expression interpretation.
type Interpreter struct {
}

// NewInterpreter creates a new interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// VisitLiteralExpr implements the visitor pattern fpr the LiteralExpr.
func (v Interpreter) VisitLiteralExpr(expr LiteralExpr) (float64, error) {
	return expr.Value, nil
}

// VisitBinaryExpr implements the visitor pattern fpr the BinaryExpr.
// nolint: nolintlint, cyclop
func (v Interpreter) VisitBinaryExpr(expr BinaryExpr) (float64, error) {
	left, err := expr.Left.Accept(v)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	right, err := expr.Right.Accept(v)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	// nolint: nolintlint, exhaustive
	switch expr.Operator.Type {
	case TokenMinus:
		return left - right, nil
	case TokenPlus:
		return left + right, nil
	case TokenStar:
		return left * right, nil
	case TokenSlash:
		if right == 0 {
			return 0, ErrDivByZero
		}

		return left / right, nil
	case TokenLeftShift:
		return float64(int64(left) << int64(right)), nil
	case TokenRightShift:
		return float64(int64(left) >> int64(right)), nil
	case TokenPow, TokenStarStar:
		return math.Pow(left, right), nil
	default:
		return 0, ErrUnknownOperator
	}
}

// VisitUnaryExpr implements the visitor pattern fpr the UnaryExpr.
func (v Interpreter) VisitUnaryExpr(expr UnaryExpr) (float64, error) {
	right, err := expr.Right.Accept(v)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	// nolint: nolintlint, exhaustive
	switch expr.Operator.Type {
	case TokenMinus:
		return -right, nil
	case TokenPlus:
		return +right, nil
	default:
		return 0, ErrUnknownOperator
	}
}

// VisitGroupingExpr implements the visitor pattern for the GroupingExpr.
func (v Interpreter) VisitGroupingExpr(expr GroupingExpr) (float64, error) {
	val, err := expr.Expr.Accept(v)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}

// Interpret will interpret an expression.
func (v Interpreter) Interpret(expr Expr) (float64, error) {
	val, err := expr.Accept(v)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}
