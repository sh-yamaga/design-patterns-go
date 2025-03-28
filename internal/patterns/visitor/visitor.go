package visitor

type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}
