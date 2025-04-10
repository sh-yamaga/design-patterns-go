package interpreter

type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

// Interpret returns true if either of the expressions evaluates to true.
func (oe *OrExpression) Interpret(context string) bool {
	return oe.Expr1.Interpret(context) || oe.Expr2.Interpret(context)
}
