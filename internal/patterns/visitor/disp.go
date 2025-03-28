package visitor

import "fmt"

type DisplayVisitor struct{}

func (dv *DisplayVisitor) VisitCircle(c *Circle) {
	fmt.Printf("DisplayVisitor: Circle (Radius=%.2f)\n", c.Radius)
}

func (dv *DisplayVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("DisplayVisitor: Rectangle (Width=%.2f, Height=%.2f)\n", r.Width, r.Height)
}
