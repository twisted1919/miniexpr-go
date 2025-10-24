package miniexpr

// Visitor is meant to implement the Visitor pattern.
type Visitor interface {
	VisitLiteralExpr(e LiteralExpr) (float64, error)
	VisitBinaryExpr(e BinaryExpr) (float64, error)
	VisitUnaryExpr(e UnaryExpr) (float64, error)
	VisitGroupingExpr(e GroupingExpr) (float64, error)
}
