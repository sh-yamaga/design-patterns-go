package element

type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}
