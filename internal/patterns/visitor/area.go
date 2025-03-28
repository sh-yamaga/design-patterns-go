package visitor

import (
	"fmt"
	"math"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/visitor/element"
)

type AreaVisitor struct{}

func (av *AreaVisitor) VisitCircle(c *element.Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("AreaVisitor: Area of Circle = %.2f\n", area)
}

func (av *AreaVisitor) VisitRectangle(r *element.Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("AreaVisitor: Area of Rectangle = %.2f\n", area)
}
