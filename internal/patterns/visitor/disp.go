package visitor

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/visitor/element"
)

type DisplayVisitor struct{}

func (dv *DisplayVisitor) VisitCircle(c *element.Circle) {
	fmt.Printf("DisplayVisitor: Circle (Radius=%.2f)\n", c.Radius)
}

func (dv *DisplayVisitor) VisitRectangle(r *element.Rectangle) {
	fmt.Printf("DisplayVisitor: Rectangle (Width=%.2f, Height=%.2f)\n", r.Width, r.Height)
}
