package interpreter

type Expression interface {
	Interpret(context string) bool
}
