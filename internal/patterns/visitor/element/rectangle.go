package element

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}
