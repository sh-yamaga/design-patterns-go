package interpreter

import "strings"

type TerminalExpression struct {
	Word string
}

// Interpret checks if a specific word is contained in the context
func (te *TerminalExpression) Interpret(context string) bool {
	return strings.Contains(strings.ToLower(context), strings.ToLower(te.Word))
}
