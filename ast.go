package miniexpr

import "fmt"

// Expr defines our expressions.
type Expr interface {
	Accept(v Visitor) (float64, error)
}

// LiteralExpr holds state for the literal expressions.
type LiteralExpr struct {
	Value float64
}

// Accept implements the Expr interface.
func (e LiteralExpr) Accept(v Visitor) (float64, error) {
	val, err := v.VisitLiteralExpr(e)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}

// BinaryExpr holds state for the binary expressions.
type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

// Accept implements the Expr interface.
func (e BinaryExpr) Accept(v Visitor) (float64, error) {
	val, err := v.VisitBinaryExpr(e)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}

// UnaryExpr holds state for the unary expressions.
type UnaryExpr struct {
	Operator Token
	Right    Expr
}

// Accept implements the Expr interface.
func (e UnaryExpr) Accept(v Visitor) (float64, error) {
	val, err := v.VisitUnaryExpr(e)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}

// GroupingExpr holds state for the grouping expressions.
type GroupingExpr struct {
	Expr Expr
}

// Accept implements the Expr interface.
func (e GroupingExpr) Accept(v Visitor) (float64, error) {
	val, err := v.VisitGroupingExpr(e)
	if err != nil {
		return 0, fmt.Errorf("evaluation failure: %w", err)
	}

	return val, nil
}
