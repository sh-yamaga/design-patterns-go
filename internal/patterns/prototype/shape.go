package prototype

type Shape interface {
	Draw()
	Clone() Shape
}
