package interpreter

type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

// Interpret returns true if both expressions evaluate to true.
func (ae *AndExpression) Interpret(context string) bool {
	return ae.Expr1.Interpret(context) && ae.Expr2.Interpret(context)
}
